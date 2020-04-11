package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptrace"
	"net/http/httputil"
	"net/url"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	c := &http.Client{}
	b, err := getRawHN(c, "01-rawhn.html")
	if err != nil {
		log.Fatal(err)
	}
	urls, err := getURLs(b, "02-urls.csv")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("raw urls: ", len(urls))
	urs, err := filterURLs(urls, "03-urls.csv")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("filtered urls: ", len(urs))
	tr, err := testURLs(c, urs, "04-test.json")
	if err != nil {
		log.Fatal(err)
	}
	stats(tr)
}

func stats(tr map[string]Site) {
	headers := map[string]int{}
	hostparts := map[string]int{}
	status := map[string]int{}
	avgdns := a{}
	avgfirst := a{}
	avgsize := a{}
	hServer := map[string]int{}

	for _, v := range tr {
		for _, h := range v.Host {
			hostparts[h]++
		}
		for _, PX := range []RoundTrip{v.P80, v.P443} {
			if len(PX.Response) > 0 {
				PR, _ := http.ReadResponse(bufio.NewReader(bytes.NewReader(PX.Response)), nil)
				defer PR.Body.Close()
				b, _ := ioutil.ReadAll(PR.Body)

				avgdns.s += int64(PX.DNSTime)
				avgdns.n++
				avgfirst.s += int64(PX.FirstByteTime)
				avgfirst.n++
				avgsize.s += int64(len(b))
				avgsize.n++
				status[PR.Status]++
				for k, v := range PR.Header {
					headers[k] += len(v)
				}
				hServer[PR.Header.Get("server")]++
			}

		}
	}

	fmt.Printf("DNS Time avg: %v\n", time.Duration(avgdns.s/avgdns.n))
	fmt.Printf("FirstByte Time avg: %v\n", time.Duration(avgfirst.s/avgfirst.n))
	fmt.Printf("Size avg: %v\n", avgsize.s/avgsize.n)
	fmt.Printf("\n\nHost parts\nrank    count\n")
	for i, h := range topn(hostparts, 10) {
		fmt.Printf("%-8d%-8d%s\n", i+1, h.n, h.s)
	}
	fmt.Printf("\n\nHost parts\nrank    count\n")
	for i, h := range topn(status, 10) {
		fmt.Printf("%-8d%-8d%s\n", i+1, h.n, h.s)
	}
	fmt.Printf("\n\nServer header\nrank    count\n")
	for i, h := range topn(hServer, -1) {
		fmt.Printf("%-8d%-8d%s\n", i+1, h.n, h.s)
	}
	fmt.Printf("\n\nHeaders\nrank    count\n")
	for i, h := range topn(headers, -1) {
		fmt.Printf("%-8d%-8d%s\n", i+1, h.n, h.s)
	}
}

func topn(m map[string]int, n int) []x {
	delete(m, "")
	s := make([]x, 0, len(m))
	for k, v := range m {
		s = append(s, x{k, v})
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i].n > s[j].n
	})
	if n == -1 || n > len(s) {
		return s
	}
	return s[:n]
}

type x struct {
	s string
	n int
}

type a struct {
	s int64
	n int64
}

type Site struct {
	Name string
	Host []string
	P80  RoundTrip
	P443 RoundTrip
}

type RoundTrip struct {
	DNSTime       time.Duration
	FirstByteTime time.Duration
	Response      []byte
}

func testURL(c *http.Client, us string) Site {
	u, _ := url.Parse(us)
	var dnstime, conntime time.Time
	var dnsdur, firstbyte time.Duration
	ct := &httptrace.ClientTrace{
		DNSStart: func(httptrace.DNSStartInfo) {
			dnstime = time.Now()
		},
		DNSDone: func(httptrace.DNSDoneInfo) {
			dnsdur = time.Since(dnstime)
		},
		GotConn: func(httptrace.GotConnInfo) {
			conntime = time.Now()
		},
		GotFirstResponseByte: func() {
			firstbyte = time.Since(conntime)
		},
	}

	var uxs []string
	if strings.HasPrefix(us, "https") {
		uxs = []string{strings.Replace(us, "https", "http", 1), us}
	} else {
		uxs = []string{us, strings.Replace(us, "http", "https", 1)}
	}

	var p80, p443 RoundTrip
	for i, ux := range uxs {
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		req, err := http.NewRequest(http.MethodGet, ux, nil)
		req = req.WithContext(httptrace.WithClientTrace(ctx, ct))
		res, err := c.Do(req)
		if err != nil {
			log.Println(ux, err)
			continue
		}
		b, err := httputil.DumpResponse(res, true)
		if err != nil {
			log.Println(ux, err)
			continue
		}
		if i == 0 {
			p80 = RoundTrip{
				DNSTime:       dnsdur,
				FirstByteTime: firstbyte,
				Response:      b,
			}
		} else {
			p443 = RoundTrip{
				DNSTime:       dnsdur,
				FirstByteTime: firstbyte,
				Response:      b,
			}
		}
	}

	return Site{
		Name: us,
		Host: strings.Split(u.Host, "."),
		P80:  p80,
		P443: p443,
	}
}

func testURLs(c *http.Client, urs []string, cacheFile string) (map[string]Site, error) {
	res := map[string]Site{}
	f, err := os.Open(cacheFile)
	if err == nil {
		defer f.Close()
		err = json.NewDecoder(f).Decode(&res)
		if err == nil {
			return res, nil
		}
	}

	var wg sync.WaitGroup
	sites := make(chan Site)
	wg.Add(len(urs))
	go func() {
		wg.Wait()
		close(sites)
	}()
	for _, u := range urs {
		go func(u string) {
			defer wg.Done()
			sites <- testURL(c, u)
		}(u)
	}
	for site := range sites {
		res[site.Name] = site
	}

	f, err = os.Create(cacheFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	e := json.NewEncoder(f)
	e.SetIndent("", "  ")
	err = e.Encode(res)
	return res, err
}

func filterURLs(urls []*url.URL, cacheFile string) ([]string, error) {
	f, err := os.Open(cacheFile)
	if err == nil {
		defer f.Close()
		r := csv.NewReader(f)
		rr, err := r.ReadAll()
		if err == nil {
			var urs []string
			for _, row := range rr {
				urs = append(urs, row[0])
			}
			return urs, nil
		}
	}

	hostpath := map[string]string{}
	for _, u := range urls {
		switch strings.ToLower(u.Host) {

		// blacklist
		case "", "news.ycombinator.com", "xkcd.com", "github.com":
			continue

		// whitelist subpaths
		case "medium.com", "dev.to", "techmeme.im", "twitter.com", "youtube.com":
			hp := strings.ToLower(u.Host) + "/" + strings.Split(u.Path, "/")[1]
			hostpath[hp] = u.Scheme

		default:
			hp := strings.ToLower(u.Host)
			hostpath[hp] = u.Scheme
		}
	}

	var urs []string
	for k := range hostpath {
		urs = append(urs, k)
	}
	sort.Strings(urs)
	var rr [][]string
	for i, u := range urs {
		urs[i] = hostpath[u] + "://" + u
		rr = append(rr, []string{hostpath[u] + "://" + u})
	}
	f, err = os.Create(cacheFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	err = csv.NewWriter(f).WriteAll(rr)
	return urs, err
}

func getURLs(b []byte, cacheFile string) ([]*url.URL, error) {
	var urls []*url.URL
	f, err := os.Open(cacheFile)
	if err == nil {
		defer f.Close()
		r := csv.NewReader(f)
		rr, err := r.ReadAll()
		if err == nil {
			for _, row := range rr {
				u, err := url.Parse(row[0])
				if err != nil {
					log.Printf("parse url %v: %s\n", err, row[0])
					continue
				}
				urls = append(urls, u)
			}
			return urls, nil
		}
	}

	d, err := goquery.NewDocumentFromReader(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	var rr [][]string
	d.Find("a[rel=nofollow]").Each(func(i int, s *goquery.Selection) {
		us, ok := s.Attr("href")
		if !ok {
			return
		}
		u, err := url.Parse(us)
		if err != nil {
			log.Printf("parse url %v: %s\n", err, us)
			return
		}
		urls = append(urls, u)
		rr = append(rr, []string{us})
	})

	f, err = os.Create(cacheFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	err = csv.NewWriter(f).WriteAll(rr)
	return urls, err
}

func getRawHN(c *http.Client, cacheFile string) ([]byte, error) {
	b, err := ioutil.ReadFile(cacheFile)
	if err == nil {
		return b, err
	}

	var buf bytes.Buffer
	baseURL := "https://news.ycombinator.com/item?id=22800136"
	for _, suffix := range []string{"", "&p=2", "p=3"} {
		r, err := c.Get(baseURL + suffix)
		if err != nil {
			return nil, err
		}
		defer r.Body.Close()
		if r.StatusCode != 200 {
			return nil, errors.New("get " + baseURL + suffix + " status " + r.Status)
		}
		io.Copy(&buf, r.Body)
	}

	err = ioutil.WriteFile(cacheFile, buf.Bytes(), 0644)

	return buf.Bytes(), err
}
