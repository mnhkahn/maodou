package maodou

import (
	. "github.com/mnhkahn/maodou/logs"
	"github.com/mnhkahn/maodou/models"
	"github.com/rakyll/ticktock"
	"github.com/rakyll/ticktock/t"
	"log"
	"net"
	"time"
)

type Handler interface {
	Init()
	Start()
	Index(resp *Response)
	Detail(resp *Response)
	Result(result *models.Result)
}

type MaoDou struct {
	req  *Request
	resp *Response
}

func (this *MaoDou) Init() {
	this.req = NewRequest()
}

func (this *MaoDou) Start() {
	log.Println("Start Method is not override.")
}

func (this *MaoDou) Cawl(paras ...string) *Response {
	log.Printf("Start Parse %s.", paras[0])
	return this.req.Cawl(paras...)
}

func (this *MaoDou) Index(resp *Response) {
	log.Println("Start Method is not override.")
	this.Detail(nil)
}

func (this *MaoDou) Detail(resp *Response) {
	log.Println("Start Method is not override.")
	this.Result(nil)
}

func (this *MaoDou) Result(result *models.Result) {
	log.Println("Start Method is not override.")
}

var APP *App

type App struct {
	handler Handler
}

func NewController(handler Handler) *App {
	app := new(App)
	app.handler = handler
	return app
}

func (this *App) Run() error {
	ColorLog("[INFO] Start parse at %s...\n", time.Now().Format(time.RFC3339))
	this.handler.Init()
	this.handler.Start()
	ColorLog("[SUCC] Parse end.\n")
	return nil
}

func Register(handler Handler, duration int) {
	app := NewController(handler)
	go func() {
		ticktock.Schedule(
			"maodou",
			app,
			&t.When{Every: t.Every(duration).Seconds()})
	}()

	app.Run()

	ServerAddr, _ := net.ResolveUDPAddr("udp", ":10001")

	/* Now listen at selected port */
	ServerConn, _ := net.ListenUDP("udp", ServerAddr)
	defer ServerConn.Close()

	buf := make([]byte, 1024)
	ServerConn.ReadFromUDP(buf)
}

// func init() {
// 	APP = NewController()
// }
