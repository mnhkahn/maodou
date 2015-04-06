package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/mnhkahn/maodou"
	"github.com/mnhkahn/maodou/cygo"
	"github.com/mnhkahn/maodou/dao"
	"github.com/mnhkahn/maodou/models"
	"strings"
	"time"
)

type NBA struct {
	maodou.MaoDou
}

func (this *NBA) Start() {
	now := time.Now()
	this.Index(this.Cawl("http://match.sports.sina.com.cn/tvguide/program/top/?date=" + now.Format("2006-01-02")))
	// resp, _ := maodou.NewResponse(strings.NewReader(a), "http://match.sports.sina.com.cn/tvguide/program/top/?date="+now.Format("2006-01-02"))
	// this.Index(resp)
}

func (this *NBA) Index(resp *maodou.Response) {
	resp.Doc(`#m01_c > blockquote > ul > li`).Each(func(i int, s *goquery.Selection) {
		t, _ := s.Attr("data-mtype")
		if t == "mtypec02" {
			html, _ := s.Html()
			resp, _ := maodou.NewResponse(strings.NewReader(html), resp.Url)
			this.Detail(resp)
		}
	})
}

func (this *NBA) Detail(resp *maodou.Response) {
	res := new(models.Result)
	res.Id = fmt.Sprintf("%d", time.Now().Unix())
	res.Title = resp.Doc("div.mth_title").Text()
	res.Author = "体育电视节目表_节目表_新浪竞技风暴_新浪网"
	// res.Figure, _ = resp.Doc("#link-report > div.topic-content > div.topic-figure.cc > img").Attr("src")
	res.Detail = resp.Doc("div.mth_channel.mth_channel_text > a").Text()
	res.Link, _ = resp.Doc("div.mth_channel.mth_channel_text > a").Attr("href")
	res.Source = resp.Url
	res.ParseDate = cygo.Now()
	this.Result(res)
}

func (this *NBA) Result(result *models.Result) {
	fmt.Println(result)
	Dao, err := dao.NewDao("duoshuo", `{"short_name":"cyeam","secret":"df66f048bd56cba5bf219b51766dec0d","thread_key":"nba"}`)
	if err != nil {
		panic(err)
	}
	Dao.AddResult(result)
}

func main() {
	maodou.Register(new(NBA), 24*60)
}
