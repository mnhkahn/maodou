package maodou

import (
	. "github.com/mnhkahn/maodou/logs"
	"github.com/mnhkahn/maodou/models"
	// "github.com/rakyll/ticktock"
	// "github.com/rakyll/ticktock/t"
	"log"
	// "net"
	"time"
)

type Handler interface {
	Init()
	Start()
	Index(resp *Response)
	Detail(resp *Response)
	Result(result *models.Result)
	Config() *HandlerConfig
}

type HandlerConfig struct {
	cawl_every time.Duration
	interval   time.Duration
}

type MaoDou struct {
	req      *Request
	resp     *Response
	settings *HandlerConfig
	Debug    bool
}

func (this *MaoDou) SetRate(times ...time.Duration) {
	this.settings = new(HandlerConfig)
	if len(times) == 1 {
		this.settings.cawl_every = times[0]
		this.req = NewRequest(0)
	} else if len(times) == 2 {
		this.settings.cawl_every = times[0]
		this.req = NewRequest(times[1])
	}
}

func (this *MaoDou) Init() {
	// this.req = NewRequest(5)
}

func (this *MaoDou) Start() {
	log.Println("Start Method is not override.")
}

func (this *MaoDou) Cawl(url string) *Response {
	log.Printf("Start Parse %s.", url)
	return this.req.Cawl(url)
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

func (this *MaoDou) Config() *HandlerConfig {
	return this.settings
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

func Register(handler Handler) {
	app := NewController(handler)

	duration := time.Duration(0)
	if handler.Config() != nil && handler.Config().cawl_every > 0 {
		duration = handler.Config().cawl_every
	}
	if duration > 0 {
		timer := time.NewTicker(duration)
		for {
			select {
			case <-timer.C:
				go func() {
					app.Run()
				}()
			}
		}
	} else {
		app.Run()
	}

	// go func() {
	// err := ticktock.Schedule(
	// 	"maodou",
	// 	app,
	// 	&t.When{Every: t.Every(100).Seconds()})
	// log.Println(err)
	// app.Run()
	// defer ticktock.Cancel("maodou")
	// }()

	// ServerAddr, _ := net.ResolveUDPAddr("udp", ":10001")

	// /* Now listen at selected port */
	// ServerConn, _ := net.ListenUDP("udp", ServerAddr)
	// defer ServerConn.Close()

	// buf := make([]byte, 1024)
	// ServerConn.ReadFromUDP(buf)
}
