package maodou

import (
	. "github.com/mnhkahn/maodou/logs"
	"github.com/mnhkahn/maodou/models"
	"log"
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

var APP *App

type App struct {
	handler Handler
}

func NewController() *App {
	app := new(App)
	return app
}

func (this *App) Run() {
	ColorLog("[INFO] Start parse at %s...\n", time.Now().Format(time.RFC3339))
	this.handler.Init()
	this.handler.Start()
	ColorLog("[SUCC] Parse end.\n")
}

func (this *App) Register(handler Handler) {
	this.handler = handler
	this.Run()
}

func init() {
	APP = NewController()
}

// ColorLog("[TRAC] Start to parse [%s]...\n")
