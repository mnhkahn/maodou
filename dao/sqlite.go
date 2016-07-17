package dao

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/mnhkahn/maodou/models"
)

type SqliteDao struct {
}

func (this *SqliteDao) NewDaoImpl(dsn string) (DaoContainer, error) {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", dsn)

	orm.RegisterModel(new(Result))

	// create table
	orm.RunSyncdb("default", false, true)

	d := new(SqliteDaoContainer)
	d.o = orm.NewOrm()
	return d, nil
}

type SqliteDaoContainer struct {
	is_debug bool
	o        orm.Ormer
}

func (this *SqliteDaoContainer) Debug(is_debug bool) {

}

func (this *SqliteDaoContainer) AddResult(p *Result) error {
	id, err := this.o.Insert(p)
	p.Id = uint64(id)
	return err
}

func (this *SqliteDaoContainer) AddResults(p []Result) {

}

func (this *SqliteDaoContainer) DelResult(id interface{}) {

}

func (this *SqliteDaoContainer) DelResults(source string) {

}

func (this *SqliteDaoContainer) UpdateResult(p *Result) {

}

func (this *SqliteDaoContainer) AddOrUpdate(p *Result) {
	this.AddResult(p)
}

func (this *SqliteDaoContainer) GetResults() ([]*Result, error) {
	var ress []*Result
	_, err := this.o.QueryTable("result").All(&ress)
	return ress, err
}

func (this *SqliteDaoContainer) GetResultById(id uint64) (*Result, error) {
	p := new(Result)
	p.Id = id
	err := this.o.Read(p)
	return p, err
}

func (this *SqliteDaoContainer) GetResultByLink(url string) *Result {
	p := new(Result)
	return p
}

func (this *SqliteDaoContainer) GetResult(author, sort string, limit, start int) []Result {
	return nil
}

func (this *SqliteDaoContainer) IsResultUpdate(p *Result) bool {
	is_update := false
	return is_update
}

func (this *SqliteDaoContainer) Search(q string, limit, start int) (int, float64, []Result) {
	return 0, 0, nil
}

func init() {
	Register("sqlite", &SqliteDao{})
}
