package maodou

import (
	"github.com/franela/goreq"
	"log"
	"net/http"
	"strings"
	"time"
)

type Request struct {
	goreq.Request
	root string
}

func NewRequest() *Request {
	goreq.SetConnectTimeout(time.Duration(60) * time.Second)
	req := new(Request)
	req.Method = "GET"
	req.UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.89 Safari/537.36"
	req.Timeout = time.Duration(60) * time.Second
	req.AddHeader("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6,zh-TW;q=0.4")
	req.AddHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	// req.ShowDebug = true
	return req
}

func (this *Request) Cawl(url string) *Response {
	this.Uri = url

	// Add referer
	if this.root == "" {
		this.root = url
	} else {
		this.AddHeader("Referer", this.root)
	}

	http_resp, err := this.Do()
	if err != nil {
		log.Println("Cawl Error: %s", err.Error())
	}

	println(http_resp.StatusCode, "****", http_resp.Header.Get("Location"))

	res_str := ""
	if http_resp.StatusCode == http.StatusOK {
		res_str, err = http_resp.Body.ToString()
		if err != nil {
			log.Println("Cawl Error: %s", err.Error())
		}
	}

	resp, err := NewResponse(strings.NewReader(res_str), url)
	if err != nil {
		log.Println("Cawl Error: %s", err.Error())
	}
	time.Sleep(5 * time.Second)
	return resp
}
