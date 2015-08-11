package main

import (
	"flag"
	"log"
	urlpkg "net/url"
	"time"

	"github.com/franela/goreq"
)

var host = flag.String("h", "host", "")

func main() {
	goreq.SetConnectTimeout(5 * time.Second)

	u := new(urlpkg.URL)
	u.Scheme = "http"
	u.Host = *host
	res, err := goreq.Request{Uri: "http://www.douban.com/group/haixiuzu/discussion", Proxy: u.String(), Timeout: time.Duration(5) * time.Second, ShowDebug: true}.Do()
	log.Println(res, err)
}
