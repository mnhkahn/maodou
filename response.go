package maodou

import (
	"github.com/PuerkitoBio/goquery"
	"io"
)

type Response struct {
	Document *goquery.Document
}

// func NewResponse() *Response {
// 	resp := new(Response)
// 	return resp
// }
//
func NewResponse(r io.Reader) (*Response, error) {
	resp := new(Response)
	var err error
	resp.Document, err = goquery.NewDocumentFromReader(r)
	return resp, err
}

func (this *Response) Doc(css string) *goquery.Selection {
	return this.Document.Find(css)
}
