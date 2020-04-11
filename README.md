# hnblogs20200407

rudimentary analysis of https://news.ycombinator.com/item?id=22800136

inspired by https://www.dannysalzman.com/2020/04/08/analyzing-hn-readers-personal-blogs

[![License](https://img.shields.io/github/license/seankhliao/hnblogs20200407.svg?style=flat-square)](LICENSE)
![Version](https://img.shields.io/github/v/tag/seankhliao/hnblogs20200407?sort=semver&style=flat-square)
[![pkg.go.dev](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://pkg.go.dev/go.seankhliao.com/hnblogs20200407)

## Results

```txt
DNS Time avg: 983.412505ms
FirstByte Time avg: 390.35683ms
Size avg: 72375


Host parts
rank    count
1       235     com
2       71      www
3       32      io
4       25      blog
5       20      net
6       17      dev
7       14      medium
8       13      blogspot
9       13      github
10      10      co


Host parts
rank    count
1       749     200 OK
2       7       404 Not Found
3       2       200
4       1       200 OK
5       1       521


Server header
rank    count
1       132     cloudflare
2       104     Netlify
3       97      GitHub.com
4       85      nginx
5       64      Apache
6       30      GSE
7       20      nginx/1.14.0 (Ubuntu)
8       17      AmazonS3
9       16      Cowboy
10      12      nginx/1.10.3 (Ubuntu)
11      12      nginx/1.17.6
12      10      LiteSpeed
13      10      nginx/1.14.2
14      10      Caddy
15      8       tsa_o
16      8       now
17      6       Apache/2.4.29 (Ubuntu)
18      6       nginx/1.17.9
19      6       nginx/1.10.3
20      6       openresty
21      5       nginx/1.16.1
22      4       nginx/1.14.1
23      4       lighttpd/1.4.53
24      2       Flywheel/5.1.0
25      2       openresty/1.15.8.3
26      2       Apache/2.4.29
27      2       nginx/1.17.5
28      2       YouTube Frontend Proxy
29      2       neocities
30      2       Apache/2.4.25 (Debian)
31      2       Apache/2
32      2       nginx/1.10.1
33      2       Squarespace
34      2       aruba-proxy
35      2       Apache/2.4.34 (Red Hat) OpenSSL/1.0.2k-fips PHP/7.3.16
36      2       gunicorn/19.9.0
37      2       nginx/1.14.0 + Phusion Passenger 6.0.4
38      2       Werkzeug/0.16.0 Python/3.8.1
39      2       Apache/2.4.18 (Ubuntu)
40      2       Apache/2.4.41 (Unix)
41      2       cat factory 1.0
42      2       mw1319.eqiad.wmnet
43      2       Apache/2.4.6
44      2       Apache/2.4.41 (Ubuntu)
45      2       nginx/1.15.5 (Ubuntu)
46      2       nginx/1.14.0
47      2       nginx/1.13.0
48      1       WSGIServer/0.2 CPython/3.8.1
49      1       UploadServer
50      1       WEBrick/1.4.2 (Ruby/2.5.1/2018-03-29)
51      1       Apache/2.4.7 (Ubuntu)


Headers
rank    count
1       760     Content-Type
2       759     Date
3       728     Server
4       628     Vary
5       522     Cache-Control
6       397     Set-Cookie
7       394     Etag
8       331     Last-Modified
9       312     Strict-Transport-Security
10      271     Expires
11      254     Age
12      207     Link
13      183     Accept-Ranges
14      183     X-Content-Type-Options
15      157     Via
16      151     X-Frame-Options
17      151     X-Cache
18      138     Access-Control-Allow-Origin
19      136     X-Proxy-Cache
20      133     X-Xss-Protection
21      132     Cf-Ray
22      132     Cf-Cache-Status
23      130     Expect-Ct
24      111     X-Cache-Hits
25      111     X-Served-By
26      111     X-Timer
27      110     X-Nf-Request-Id
28      110     X-Github-Request-Id
29      103     X-Fastly-Request-Id
30      97      X-Powered-By
31      79      Connection
32      76      Alt-Svc
33      65      Referrer-Policy
34      48      Content-Security-Policy
35      45      Content-Length
36      44      X-Ac
37      38      Pragma
38      38      X-Hacker
39      36      X-Ua-Compatible
40      32      Host-Header
41      31      Sepia-Upstream
42      28      X-Amz-Cf-Pop
43      28      X-Amz-Cf-Id
44      17      Medium-Fulfilled-By
45      17      X-Envoy-Upstream-Service-Time
46      16      X-Request-Id
47      14      X-Obvious-Info
48      14      X-Server-Cache
49      14      X-Permitted-Cross-Domain-Policies
50      14      X-Obvious-Tid
51      14      X-Download-Options
52      14      Upgrade
53      13      X-Amz-Id-2
54      13      X-Amz-Request-Id
55      12      Status
56      12      X-Cache-Enabled
57      11      P3p
58      10      X-Runtime
59      10      Content-Disposition
60      10      X-Endurance-Cache-Level
61      10      X-Nananana
62      8       X-Now-Cache
63      8       X-Cache-Handler
64      8       Access-Control-Allow-Methods
65      8       X-Now-Id
66      8       X-Transaction
67      8       X-Cache-Status
68      8       Content-Language
69      8       X-Now-Trace
70      8       X-Response-Time
71      8       X-Connection-Hash
72      8       X-Twitter-Response-Tags
73      6       X-Turbo-Charged-By
74      6       Feature-Policy
75      6       X-Cacheable
76      6       Access-Control-Allow-Headers
77      5       X-Pingback
78      4       X-Tumblr-Pixel-0
79      4       X-Robots-Tag
80      4       X-Amz-Version-Id
81      4       X-Fw-Static
82      4       X-Tumblr-User
83      4       X-Goog-Hash
84      4       X-Fw-Serve
85      4       X-Fw-Type
86      4       X-Tumblr-Pixel-1
87      4       X-Tumblr-Pixel
88      4       X-Fw-Hash
89      4       X-Fw-Server
90      4       X-Rid
91      3       X-Goog-Stored-Content-Encoding
92      3       X-Goog-Generation
93      3       X-Goog-Stored-Content-Length
94      3       X-Goog-Metageneration
95      3       X-Iplb-Instance
96      3       X-Guploader-Uploadid
97      3       X-Goog-Storage-Class
98      2       X-Varnish
99      2       X-Hosted-By
100     2       X-Httpd
101     2       X-Seen-By
102     2       X-Aspnet-Version
103     2       X-Fw-Version
104     2       Superexpress
105     2       Cc-Fetch-Error
106     2       X-Client-Ip
107     2       X-Mod-Pagespeed
108     2       X-Wix-Request-Id
109     2       X-Uri
110     2       X-Tumblr-Pixel-2
111     2       X-Backend
112     2       X-Hits
113     2       X-Litespeed-Cache
114     2       Cc-Stable-Domain
115     2       X-Neocities-Cdn
116     2       X-Contextid
117     2       X-Amz-Meta-Kopi-Version
118     2       Access-Control-Allow-Credentials
119     2       Cc-Cache-Status
120     2       X-Servername
121     2       X-Amz-Meta-Md5-Hash
122     2       Server-Timing
123     2       X-Fw-Dynamic
124     2       X-Ghost-Cache-Status
125     2       X-Cache-Type
126     2       X-Common-Inject
127     2       X-Amz-Meta-Jets3t-Original-File-Date-Iso8601
128     2       Public-Key-Pins
129     2       X-Aspnetmvc-Version
130     2       Upgrade-Insecure-Requests
131     2       Fastly-Restarts
132     2       X-Cache-Device-Type
133     2       X-Clacks-Overhead
134     2       Access-Control-Expose-Headers
135     2       Proxied-To
136     2       X-Cache-Hit
137     2       X-Cached
138     1       Ms-Author-Via
```
