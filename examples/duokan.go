package main

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/mnhkahn/maodou"
	"github.com/mnhkahn/maodou/cygo"
	"github.com/mnhkahn/maodou/models"
)

type Duokan struct {
	maodou.MaoDou
}

func NewDuokan() *Duokan {
	duo := new(Duokan)
	duo.Init()
	duo.SetDao("duoshuo", `{"short_name":"cyeam","secret":"df66f048bd56cba5bf219b51766dec0d","thread_key":"39392"}`)
	duo.SetRate(time.Duration(24)*time.Hour, time.Duration(5)*time.Second)
	return duo
}

func (this *Duokan) Start() {
	this.Index(nil)
}

func (this *Duokan) Index(resp *maodou.Response) {
	resp, err := this.Cawl("http://www.duokan.com/book/39392")
	if err == nil {
		this.Detail(resp)
	}
}

func (this *Duokan) Detail(resp *maodou.Response) {
	res := new(models.Result)

	res.Id = cygo.Now().Format(cygo.DATE_DAY)
	res.Title = resp.Doc("title").Text()
	res.Link = resp.Url
	res.ParseDate = cygo.Now()
	price := new(Price)
	price.Price, _ = strconv.ParseFloat(strings.TrimLeft(resp.Doc(".price > em").Text(), "¥ "), 64)
	resp.Doc(".price > i").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			price.ListPrice, _ = strconv.ParseFloat(strings.TrimLeft(s.Find("del").Text(), "¥ "), 64)
		} else {
			price.PaperPrice, _ = strconv.ParseFloat(strings.TrimLeft(s.Find("del").Text(), "¥ "), 64)
		}
	})
	log.Printf("%s 现价%v 原价%v", res.Title, price.Price, price.ListPrice)
	// this.Result(res)
}

func (this *Duokan) Result(result *models.Result) {
	this.Dao.AddResult(result)
}

type Price struct {
	Price      float64 `json:"price"`
	ListPrice  float64 `json:"list_price"`
	PaperPrice float64 `json:"paper_price"`
}

func main() {
	duokan := NewDuokan()
	maodou.Register(duokan)
}
