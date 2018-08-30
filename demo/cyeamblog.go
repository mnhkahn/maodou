// Package demo
package main

import (
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/mnhkahn/gogogo/logger"

	"github.com/PuerkitoBio/goquery"
	"github.com/mnhkahn/maodou"
	"github.com/mnhkahn/maodou/models"
)

type CyeamBlogCrawler struct {
	maodou.MaoDou
}

func (this *CyeamBlogCrawler) Start() {
	resp, err := this.Cawl("http://blog.cyeam.com/")
	if err != nil {
		logger.Warn(err)
	} else {
		this.Index(resp)
	}
}

func (this *CyeamBlogCrawler) Index(resp *maodou.Response) {
	resp.Doc(`#content > div > div > div > h2 > a`).Each(func(i int, s *goquery.Selection) {
		href, has := s.Attr("href")
		//logger.Infof("AA %+v", href)
		if has {
			resp, err := this.Cawl("http://blog.cyeam.com/" + href)
			if err != nil {
				logger.Warn(err)
			} else {
				this.Detail(resp)
			}
		}
	})
}

func (this *CyeamBlogCrawler) Detail(resp *maodou.Response) {
	res := new(models.Result)
	u, _ := url.Parse(resp.Url)
	res.Id = u.Path
	res.Title = resp.Doc("#content > div.page-header.c6 > h1").Text()
	res.Author = "Bryce"
	res.Figure, _ = resp.Doc("#content > div.row-fluid.post-full > div > p.figure.center > img").Attr("src")
	res.Link = resp.Url
	res.ParseDate = time.Now()
	res.Description = resp.Doc("#content > div.row-fluid.post-full > div > blockquote").Text()
	res.Detail = resp.Doc("#content > div.row-fluid.post-full > div > div.content").Text()
	_tag := resp.Doc("#content > div.row-fluid.post-full > div > ul:nth-child(7)").Text()
	tags := []string{}
	for _, _t := range strings.Split(_tag, " ") {
		_tt := strings.TrimSpace(_t)
		if len(_tt) > 1 {
			_, err := strconv.Atoi(_tt)
			if err != nil {
				tags = append(tags, strings.Split(strings.TrimSpace(_t), " ")[0])
			}
		}
	}
	res.Tags = strings.Join(tags, " ")
	_category := resp.Doc("#content > div.row-fluid.post-full > div > ul:nth-child(6) > li:nth-child(2) > a").Text()
	res.Category = strings.Split(strings.TrimSpace(_category), " ")[0]
	this.Result(res)
}

func (this *CyeamBlogCrawler) Result(result *models.Result) {
	logger.Info(result.Title, result.Link)
	//if result.Figure != "" {
	//	Dao, err := dao.NewDao("duoshuo", `{"short_name":"cyeam","secret":"df66f048bd56cba5bf219b51766dec0d"}`)
	//	if err != nil {
	//		panic(err)
	//	}
	//	Dao.AddResult(result)
	//}
}

func main() {
	maodou.Register(NewCyeamBlogCrawler())
}

func NewCyeamBlogCrawler() *CyeamBlogCrawler {
	c := new(CyeamBlogCrawler)
	c.MaoDou.Init()
	return c
}
