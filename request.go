package maodou

import (
	"fmt"
	"github.com/franela/goreq"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Request struct {
	goreq.Request
	root     string
	interval time.Duration
}

func NewRequest(interval time.Duration) *Request {
	goreq.SetConnectTimeout(time.Duration(60) * time.Second)
	req := new(Request)
	req.Method = "GET"
	req.UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.89 Safari/537.36"
	req.Timeout = time.Duration(60) * time.Second
	req.AddHeader("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6,zh-TW;q=0.4")
	req.AddHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.interval = interval
	// req.ShowDebug = true
	return req
}

func (this *Request) Cawl(url string) (*Response, error) {
	this.Uri = url

	log.Println("Start to Parse:", this.Uri)
	// Add referer
	if this.root == "" {
		this.root = url
	} else {
		this.AddHeader("Referer", this.root)
	}

	http_resp, err := this.Do()
	if err != nil {
		log.Printf("Cawl Error: %s\n", err.Error())
		return nil, err
	}

	var resp *Response
	if http_resp.StatusCode == http.StatusOK {
		resp, err = NewResponse(http_resp.Body, url)
		if err != nil {
			log.Printf("Cawl Error: %s.\n", err.Error())
			return resp, err
		} else {
			log.Println("Cawl Success.")
		}
	} else if http_resp.StatusCode == http.StatusMovedPermanently || http_resp.StatusCode == http.StatusFound {
		return this.Cawl(http_resp.Header.Get("Location"))
	} else {
		log.Printf("Cawl Got Status Code %d.\n", http_resp.StatusCode)
		return resp, fmt.Errorf("Cawl Got Status Code %d.", http_resp.StatusCode)
	}

	if this.interval > 0 {
		time.Sleep(this.interval)
	}
	return resp, nil
}

func (this *Request) DumpRequest() string {
	return this.Uri + "?" + url.Values(this.QueryString.(url.Values)).Encode()
}
