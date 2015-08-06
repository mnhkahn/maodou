package maodou

import (
	"fmt"
	"log"
	"net/http"
	urlpkg "net/url"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/franela/goreq"
)

type Request struct {
	goreq.Request
	proxies     []*ProxyConfig
	proxy_index int
	root        string
	Interval    time.Duration
}

func NewRequest(interval time.Duration) *Request {
	goreq.SetConnectTimeout(time.Duration(60) * time.Second)
	req := new(Request)
	req.Method = "GET"
	req.UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.89 Safari/537.36"
	req.Timeout = time.Duration(60) * time.Second
	req.AddHeader("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6,zh-TW;q=0.4")
	req.AddHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Interval = interval
	// req.ShowDebug = true

	if len(req.proxies) == 0 {
		resp, _ := req.Cawl("http://www.xici.net.co/")
		if resp != nil {
			resp.Doc(`#ip_list > tbody > tr`).Each(func(i int, s *goquery.Selection) {
				if i > 1 && len(req.proxies) < 10 && s.Children().Length() == 7 {
					p := new(ProxyConfig)
					p.ip = s.Children().Get(1).FirstChild.Data
					p.port, _ = strconv.Atoi(s.Children().Get(2).FirstChild.Data)
					if s.Children().Get(3).FirstChild != nil {
						p.location = s.Children().Get(3).FirstChild.Data
					}
					p.safe = s.Children().Get(4).FirstChild.Data == "高匿"
					p.schema = s.Children().Get(5).FirstChild.Data
					p.verifytime = s.Children().Get(6).FirstChild.Data
					req.proxies = append(req.proxies, p)
				}
			})
		}
	}

	return req
}

const (
	CAWL_NOPROXY = 0
	CAWL_PROXY   = 1
)

func (this *Request) Cawl(paras ...interface{}) (*Response, error) {
	this.Uri = paras[0].(string)

	// Add referer
	if this.root == "" {
		this.root = this.Uri
	} else {
		this.AddHeader("Referer", this.root)
	}

	if len(this.proxies) > 0 && (len(paras) == 1 || (len(paras) == 2 && paras[1].(int) == CAWL_PROXY)) {
		u := new(urlpkg.URL)
		p := this.proxies[this.proxy_index]
		u.Scheme = "http"
		u.Host = fmt.Sprintf("%s:%d", p.ip, p.port)
		this.Proxy = u.String()
	}

	log.Println("Start to Parse: ", this.Uri, "proxy:", this.Proxy)

	http_resp, err := this.Do()
	// 修复代理错乱的问题，需要重置代理
	this.Proxy = ""
	if err != nil {
		this.proxy_index = (this.proxy_index + 1) % len(this.proxies)
		log.Printf("Cawl Error: %s\n", err.Error())
		return nil, err
	}

	var resp *Response
	if http_resp.StatusCode == http.StatusOK {
		resp, err = NewResponse(http_resp.Body, this.Uri)
		if err != nil {
			log.Printf("Cawl Error: %s.\n", err.Error())
			return resp, err
		} else {
			log.Println("Cawl Success.")
		}
	} else {
		this.proxy_index = (this.proxy_index + 1) % len(this.proxies)
		if http_resp.StatusCode == http.StatusMovedPermanently || http_resp.StatusCode == http.StatusFound {
			log.Println(this.Uri, http_resp.StatusCode)
			return this.Cawl(http_resp.Header.Get("Location"))
		} else {
			log.Printf("Cawl Got Status Code %d.\n", http_resp.StatusCode)
			return resp, fmt.Errorf("Cawl Got Status Code %d.", http_resp.StatusCode)
		}
	}

	if this.Interval > 0 {
		time.Sleep(this.Interval)
	}
	return resp, nil
}

func (this *Request) DumpRequest() string {
	return this.Uri + "?" + urlpkg.Values(this.QueryString.(urlpkg.Values)).Encode()
}

type ProxyConfig struct {
	ip         string
	port       int
	location   string
	safe       bool
	schema     string
	verifytime string
}
