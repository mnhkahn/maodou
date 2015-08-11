package main

import (
	"flag"
	"log"
	urlpkg "net/url"
	"time"

	"github.com/franela/goreq"
)

var host = flag.String("h", "host", "")
var url = flag.String("u", "http://www.douban.com/group/topic/78216948/", "")

func main() {
	goreq.SetConnectTimeout(5 * time.Second)

	flag.Parse()

	u := new(urlpkg.URL)
	u.Scheme = "http"
	u.Host = *host
	res, err := goreq.Request{Uri: *url, UserAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.89 Safari/537.36", Proxy: u.String(), Timeout: time.Duration(5) * time.Second, ShowDebug: true}.Do()
	log.Println(res.Status, err)
}
