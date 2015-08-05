package main

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/mnhkahn/maodou"
	"github.com/mnhkahn/maodou/cygo"
	"github.com/mnhkahn/maodou/dao"
	"github.com/mnhkahn/maodou/models"
)

type Haixiu struct {
	maodou.MaoDou
}

func (this *Haixiu) Start() {
	resp, err := this.Cawl("http://www.douban.com/group/haixiuzu/discussion", maodou.CAWL_NOPROXY)
	if err == nil {
		this.Index(resp)
	}
}

func (this *Haixiu) Index(resp *maodou.Response) {
	resp.Doc(`#content > div > div.article > div:nth-child(2) > table > tbody > tr > td.title > a`).Each(func(i int, s *goquery.Selection) {
		href, has := s.Attr("href")
		if has {
			resp, err := this.Cawl(href)
			if err == nil {
				this.Detail(resp)
			}
		}
	})
}

func (this *Haixiu) Detail(resp *maodou.Response) {
	if len(strings.Split(resp.Url, "/")) < 6 {
		return
	}
	res := new(models.Result)
	res.Id = strings.Split(resp.Url, "/")[5]
	res.Title = resp.Doc("#content > h1").Text()
	res.Author = resp.Doc("#content > div > div.article > div.topic-content.clearfix > div.topic-doc > h3 > span.from > a").Text()
	figures := []string{}
	resp.Doc("#link-report > div.topic-content > div.topic-figure.cc").Each(func(i int, s *goquery.Selection) {
		f, exists := s.Find("img").Attr("src")
		if exists {
			figures = append(figures, f)
		}
	})
	res.Figure = strings.Join(figures, ",")
	res.Link = resp.Url
	res.Source = "www.douban.com/group/haixiuzu/discussion"
	res.ParseDate = cygo.Now()
	res.Description = "haixiuzu"
	this.Result(res)
}

func (this *Haixiu) Result(result *models.Result) {
	if result.Figure != "" {
		Dao, err := dao.NewDao("duoshuo", `{"short_name":"cyeam","secret":"df66f048bd56cba5bf219b51766dec0d","thread_key":"haixiuzucyeam"}`)
		if err != nil {
			panic(err)
		}
		Dao.AddResult(result)
	}
}

func main() {
	haixiu := new(Haixiu)
	haixiu.Init()
	haixiu.SetRate(time.Duration(30) * time.Minute)
	maodou.Register(haixiu)
}
