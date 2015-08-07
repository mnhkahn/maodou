package proxy

import (
	"fmt"
	"testing"
	"time"
)

func TestXici(t *testing.T) {
	p, _ := NewProxy("xici", `{"max_cawl_cnt":1,"cnt":2,"min_cnt":1}`)
	p.Init()
	fmt.Println(p.One())
	fmt.Println(p.One())
	fmt.Println(p.One())
	time.Sleep(5 * time.Second)
}
