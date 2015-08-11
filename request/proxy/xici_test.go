package proxy

import (
	"fmt"
	"testing"
	"time"
)

func TestXici(t *testing.T) {
	p, _ := NewProxy("xici", `{"max_cawl_cnt":3,"cnt":4,"min_cnt":1,"root":"http://cyeam.com"}`)
	p.Init()
	one := p.One()
	one.Delayed += time.Second
	p.Update(one)
	fmt.Println(p.One())
	fmt.Println(p.One())
	fmt.Println(p.One())
	fmt.Println(p.One())
	fmt.Println(p.One())
	fmt.Println(p.One())
	fmt.Println(p.One())
	// time.Sleep(5 * time.Second)
}
