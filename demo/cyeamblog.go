// Package demo
package main

import (
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/mnhkahn/maodou"
	"github.com/mnhkahn/maodou/models"
)

type CyeamBlogCrawler struct {
	maodou.MaoDou
}

func (this *CyeamBlogCrawler) Index(resp *maodou.Response, jobs chan string) error {
	resp.Doc(`#content > div > div > div > h2 > a`).Each(func(i int, s *goquery.Selection) {
		href, has := s.Attr("href")
		if has {
			this.AddJob("http://blog.cyeam.com/" + href)
		}
	})
	return nil
}

func (this *CyeamBlogCrawler) Detail(resp *maodou.Response) (*models.Result, error) {
	var err error
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
	date := resp.Doc("#content > div.row-fluid.post-full > div > div.date > span:nth-child(2)").Text()
	res.CreateTime, err = time.Parse("02 January 2006", date)
	if err != nil {
		return res, err
	}
	return res, nil
}

func main() {
	maodou.Register(NewCyeamBlogCrawler())
}

func NewCyeamBlogCrawler() *CyeamBlogCrawler {
	c := new(CyeamBlogCrawler)
	c.MaoDou.Init("http://blog.cyeam.com/#all.html")
	return c
}
