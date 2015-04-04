package dao

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/franela/goreq"
	. "github.com/mnhkahn/maodou/models"
	"log"
	"net/url"
)

type DuoShuoConfig struct {
	ShortName string `json:"short_name"`
	Secret    string `json:"secret"`
}

type DuoShuoDao struct {
}

func (this *DuoShuoDao) NewDaoImpl(dsn string) (DaoContainer, error) {
	d := new(DuoShuoDaoContainer)

	config := new(DuoShuoConfig)
	err := json.Unmarshal([]byte(dsn), config)
	d.config = config
	if err != nil {
		return d, fmt.Errorf("Config for duoshuo is error: %v", err)
	}

	return d, nil
}

type DuoShuoDaoContainer struct {
	config   *DuoShuoConfig
	is_debug bool
	req      goreq.Request
}

func (this *DuoShuoDaoContainer) Debug(is_debug bool) {
	this.req.ShowDebug = is_debug
}

func (this *DuoShuoDaoContainer) AddResult(p *Result) {
	this.req.Method = "POST"
	this.req.Uri = "http://api.duoshuo.com/posts/import.json"
	this.req.ContentType = "application/x-www-form-urlencoded"

	addDuoShuo := url.Values{}
	addDuoShuo.Add("short_name", this.config.ShortName)
	addDuoShuo.Add("secret", this.config.Secret)
	addDuoShuo.Add("posts[0][post_key]", p.Id)
	addDuoShuo.Add("posts[0][thread_key]", "haixiuzucyeam")

	duoshuo_byte, _ := json.Marshal(*p)
	addDuoShuo.Add("posts[0][message]", base64.StdEncoding.EncodeToString(duoshuo_byte))
	this.req.Body = fmt.Sprintf("short_name=%s&secret=%s&posts[0][post_key]=%s&posts[0][thread_key]=%s&posts[0][message]=%s", this.config.ShortName, this.config.Secret, p.Id, "haixiuzucyeam", base64.URLEncoding.EncodeToString(duoshuo_byte))
	this.req.ShowDebug = true
	resp, err := this.req.Do()
	if err != nil {
		log.Println(err.Error())
	}
	if resp.StatusCode != 200 {
		err_str, _ := resp.Body.ToString()
		log.Printf("Error: %d, %s\n", resp.StatusCode, err_str)
	} else {
		log.Println("Add to DuoShuo Success.")
	}
}

func (this *DuoShuoDaoContainer) AddResults(p []Result) {

}

func (this *DuoShuoDaoContainer) DelResult(id interface{}) {

}

func (this *DuoShuoDaoContainer) DelResults(source string) {

}

func (this *DuoShuoDaoContainer) UpdateResult(p *Result) {

}

func (this *DuoShuoDaoContainer) AddOrUpdate(p *Result) {
	this.AddResult(p)
}

func (this *DuoShuoDaoContainer) GetResultById(id int) *Result {
	p := new(Result)
	return p
}

func (this *DuoShuoDaoContainer) GetResultByLink(url string) *Result {
	p := new(Result)
	return p
}

func (this *DuoShuoDaoContainer) GetResult(author, sort string, limit, start int) []Result {
	return nil
}

func (this *DuoShuoDaoContainer) IsResultUpdate(p *Result) bool {
	is_update := false
	return is_update
}

func (this *DuoShuoDaoContainer) Search(q string, limit, start int) (int, float64, []Result) {
	println("*************8serach")
	this.req.Method = "GET"
	this.req.Uri = "http://api.duoshuo.com/threads/listResults.json"
	this.req.ContentType = "application/x-www-form-urlencoded"

	// addDuoShuo := url.Values{}
	// addDuoShuo.Add("short_name", this.config.ShortName)
	// addDuoShuo.Add("secret", this.config.Secret)
	// addDuoShuo.Add("Results[0][Result_key]", p.Id)
	// addDuoShuo.Add("Results[0][thread_key]", "haixiuzu-cyeam")

	// duoshuo_byte, _ := json.Marshal(addDuoShuo)
	// addDuoShuo.Add("Results[0][message]", base64.URLEncoding.EncodeToString(duoshuo_byte))
	// this.req.Body = addDuoShuo.Encode()

	// resp, err := this.req.Do()
	// if err != nil {
	// 	panic(err)
	// }
	// if resp.StatusCode != 200 {
	// 	err_str, _ := resp.Body.ToString()
	// 	panic(err_str)
	// }
	return 0, 0, nil
}

func init() {
	Register("duoshuo", &DuoShuoDao{})
}
