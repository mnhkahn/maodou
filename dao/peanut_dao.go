// Package dao
package dao

import (
	"strings"
	"time"

	"github.com/mnhkahn/gogogo/logger"

	"github.com/mnhkahn/peanut/index"

	. "github.com/mnhkahn/maodou/models"
)

type PeanutDao struct {
}

func (this *PeanutDao) NewDaoImpl(dsn string) (DaoContainer, error) {
	var err error
	d := new(PeanutContainer)
	d.ind, err = index.NewIndex(dsn)
	if err != nil {
		return d, err
	}
	return d, nil
}

type PeanutContainer struct {
	ind *index.Index
}

func (this *PeanutContainer) Debug(is_debug bool) {

}

func (this *PeanutContainer) AddResult(p *Result) error {
	document := this.toDucument(p)
	err := this.ind.AddDocument(document)
	return err
}

func (this *PeanutContainer) toDucument(p *Result) *index.Document {
	document := new(index.Document)
	document.PK = p.Id
	document.Title = p.Title
	document.Brief = p.Description
	document.FullText = p.Detail
	document.PubDate = p.ParseDate.UnixNano()
	document.Category = p.Category
	document.Tags = strings.Split(p.Tags, " ")
	document.Link = p.Link
	document.Figure = p.Figure
	return document
}

func (this *PeanutContainer) toResults(ps []*index.Document) []Result {
	res := make([]Result, 0, len(ps))
	for _, p := range ps {
		rr := Result{
			Id:          p.PK,
			Title:       p.Title,
			Description: p.Brief,
			Detail:      p.FullText,
			ParseDate:   time.Unix(0, p.PubDate),
			Category:    p.Category,
			Tags:        strings.Join(p.Tags, " "),
			Link:        p.Link,
			Figure:      p.Figure,
		}
		res = append(res, rr)
	}
	return res
}

func (this *PeanutContainer) AddResults(p []Result) {

}

func (this *PeanutContainer) DelResult(id interface{}) {

}

func (this *PeanutContainer) DelResults(source string) {

}

func (this *PeanutContainer) UpdateResult(p *Result) {

}

func (this *PeanutContainer) AddOrUpdate(p *Result) {
	this.AddResult(p)
}

func (this *PeanutContainer) GetResults() ([]*Result, error) {
	return nil, nil
}

func (this *PeanutContainer) GetResultById(id uint64) (*Result, error) {
	return nil, nil
}

func (this *PeanutContainer) GetResultByLink(url string) *Result {
	p := new(Result)
	return p
}

func (this *PeanutContainer) GetResult(author, sort string, limit, start int) []Result {
	return nil
}

func (this *PeanutContainer) IsResultUpdate(p *Result) bool {
	return false
}

func (this *PeanutContainer) Search(q string, limit, start int) (int, float64, []Result) {
	cnt, res, err := this.ind.Search(&index.Param{
		Query:  q,
		Offset: start,
		Size:   limit,
	})
	if err != nil {
		logger.Warn(err)
	}

	return cnt, 0, this.toResults(res)
}

func init() {
	Register("peanut", &PeanutDao{})
}
