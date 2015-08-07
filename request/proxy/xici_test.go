package proxy

import (
	"fmt"
	"testing"
	"time"
)

func TestXici(t *testing.T) {
	p, _ := NewProxy("xici", `{"max_cawl_cnt":1,"cnt":2,"min_cnt":1,"root":"http://cyeam.com"}`)
	p.Init()
	a := p.One()
	fmt.Println(a)
	p.DeleteProxy(a.Id)
	fmt.Println(p.One())
	time.Sleep(5 * time.Second)
}
