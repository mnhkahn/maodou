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

	// "github.com/mnhkahn/maodou/proxy"
)

type Request struct {
	goreq.Request
	proxies  []*ProxyConfig
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

	resp, _ := req.Cawl("http://www.xici.net.co/")
	resp.Doc(`#ip_list > tbody > tr`).Each(func(i int, s *goquery.Selection) {
		if i > 1 && len(req.proxies) < 10 && s.Children().Length() == 7 {
			p := new(ProxyConfig)
			p.ip = s.Children().Get(1).FirstChild.Data
			p.port, _ = strconv.Atoi(s.Children().Get(2).FirstChild.Data)
			p.location = s.Children().Get(3).FirstChild.Data
			p.safe = s.Children().Get(4).FirstChild.Data == "高匿"
			p.schema = s.Children().Get(5).FirstChild.Data
			p.verifytime = s.Children().Get(6).FirstChild.Data
			req.proxies = append(req.proxies, p)
		}
	})
	return req
}

func (this *Request) Cawl(url string) (*Response, error) {
	this.Uri = url

	// Add referer
	if this.root == "" {
		this.root = url
	} else {
		this.AddHeader("Referer", this.root)
	}

	if len(this.proxies) > 0 {
		u := new(urlpkg.URL)
		p := this.proxies[0]
		u.Scheme = "http"
		u.Host = fmt.Sprintf("%s:%d", p.ip, p.port)
		this.Proxy = u.String()
	}

	log.Println("Start to Parse: ", this.Uri, "proxy:", this.Proxy)

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
		log.Println(url, http_resp.StatusCode)
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
