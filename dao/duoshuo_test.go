package dao

import (
	"github.com/mnhkahn/maodou/models"
	"testing"
)

func TestDuoshuo(t *testing.T) {
	result := new(models.Result)
	Dao, err := NewDao("duoshuo", `{"short_name":"cyeam","secret":"df66f048bd56cba5bf219b51766dec0d"}`)
	if err != nil {
		panic(err)
	}
	Dao.AddResult(result)
}
