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
	document.PubDate = p.CreateTime.UnixNano()
	document.Category = p.Category
	document.Tags = strings.Split(p.Tags, " ")
	document.Link = p.Link
	document.Figure = p.Figure
	document.PV = p.PV
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
			PV:          p.PV,
		}
		res = append(res, rr)
	}
	return res
}

func (this *PeanutContainer) DelResult(id interface{}) {

}

func (this *PeanutContainer) GetResults() ([]*Result, error) {
	return nil, nil
}

func (this *PeanutContainer) GetResultById(id uint64) (*Result, error) {
	return nil, nil
}

func (this *PeanutContainer) Search(q string, limit, start int, sort string, asc bool) (int, float64, []Result) {
	cnt, res, err := this.ind.Search(&index.Param{
		Query:  q,
		Offset: start,
		Size:   limit,
		Sort:   index.Sorter{sort, asc},
	})

	if err != nil {
		logger.Warn(err)
	}

	return cnt, 0, this.toResults(res)
}

func (this *PeanutContainer) Close() error {
	return this.ind.Close()
}

func init() {
	Register("peanut", &PeanutDao{})
}
