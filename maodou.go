package maodou

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/mnhkahn/gogogo/logger"
	"github.com/mnhkahn/maodou/dao"
	"github.com/mnhkahn/maodou/models"
)

type Handler interface {
	Init(link string)
	Cawl(paras ...interface{}) (*Response, error)
	Start(link string) (*Response, error)
	Index(resp *Response, jobs chan string) error
	Detail(resp *Response) (*models.Result, error)
	Result(result *models.Result) error
	Config() *HandlerConfig
}

type HandlerConfig struct {
	d          bool
	link       string // start crawl link
	dao_name   string
	dsn        string
	cawl_every time.Duration
	interval   time.Duration
	ip         string
	jobs       chan string
	wg         sync.WaitGroup
}

type MaoDou struct {
	Dao      dao.DaoContainer
	req      *Request
	resp     *Response
	settings *HandlerConfig
	Debug    bool
}

func (this *MaoDou) SetRate(times ...time.Duration) {
	if len(times) == 1 {
		this.settings.cawl_every = times[0]
	} else if len(times) == 2 {
		this.settings.cawl_every = times[0]
		this.req.Interval = times[1]
	}
}

func (this *MaoDou) SetProxy(proxy_name, proxy_dsn string) {
	this.req.InitProxy(proxy_name, proxy_dsn)
}

func (this *MaoDou) SetD(d bool) {
	this.settings.d = d
}

func (this *MaoDou) SetDao(dao_name, dsn string) {
	if this.Dao != nil {
		this.Dao.Close()
	}

	var err error
	this.settings.dao_name, this.settings.dsn = dao_name, dsn
	this.Dao, err = dao.NewDao(this.settings.dao_name, this.settings.dsn)
	if err != nil {
		panic(err)
	}
}

func (this *MaoDou) SetJobLen(l int) {
	if l <= 0 {
		l = 1
	}
	this.settings.jobs = make(chan string, l)
}

func (this *MaoDou) Init(link string) {
	var err error
	this.settings = new(HandlerConfig)

	this.settings.link = link
	this.settings.cawl_every = 0
	this.settings.interval = 0
	this.req = NewRequest(0)
	this.SetJobLen(1)
	this.SetDao("peanut", "./maodou.db")
	this.SetD(true)
	if err != nil {
		panic(err)
	}
}

func (this *MaoDou) Start(link string) (*Response, error) {
	resp, err := this.Cawl(link)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (this *MaoDou) Cawl(paras ...interface{}) (*Response, error) {
	return this.req.Cawl(paras...)
}

func (this *MaoDou) Index(resp *Response, jobs chan string) error {
	return fmt.Errorf("the Index Method is not override")
}

func (this *MaoDou) AddJob(job string) {
	this.settings.jobs <- job
}

func (this *MaoDou) Detail(resp *Response) (*models.Result, error) {
	return nil, fmt.Errorf("the Detail is not override")
}

func (this *MaoDou) Result(result *models.Result) error {
	if this.Dao != nil {
		return this.Dao.AddResult(result)
	}
	return fmt.Errorf("dao is not init")
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
	duration := time.Duration(0)
	if this.handler.Config() != nil && this.handler.Config().cawl_every > 0 {
		duration = this.handler.Config().cawl_every
	}
	APP.run()
	if duration > 0 {
		timer := time.NewTicker(duration)
		for {
			select {
			case <-timer.C:
				go func() {
					APP.run()
				}()
			}
		}
	}

	return nil
}

func (this *App) run() {
	var done = make(chan struct{})
	logger.Infof("[INFO] Start parse at %s...\n", time.Now().Format(time.RFC3339))
	if resp, err := this.handler.Start(this.handler.Config().link); err == nil {
		// produce
		go this.handler.Index(resp, this.handler.Config().jobs)
		// consume
		go func() {
			for job := range this.handler.Config().jobs {
				logger.Debug("do job:", job)
				if resp, err := this.handler.Cawl(job); err == nil {
					if res, err := this.handler.Detail(resp); err == nil {
						if err := this.handler.Result(res); err != nil {
							logger.Warn(err)
						}
					} else {
						logger.Warn(err)
					}
				} else {
					logger.Warn(err)
				}
			}
			if this.handler.Config().d {
				done <- struct{}{}
			}
		}()
		if this.handler.Config().d {
			<-done
		}
	} else {
		logger.Warn(err)
	}

	logger.Infof("[SUCC] Parse end.\n")
}

func Register(handler Handler) {
	APP = NewController(handler)
	APP.Run()
}

func init() {
	rand.Seed(time.Now().UnixNano())
	println(CYEAM_LOG)
}

var CYEAM_LOG = `
 ______                  _                    
|  ___ \                | |                   
| | _ | | ____  ___   _ | | ___  _   _              ____ _____   _   _   _ 
| || || |/ _  |/ _ \ / || |/ _ \| | | |    ___     / _\ V / __| / \ | \_/ |  
| || || ( ( | | |_| ( (_| | |_| | |_| |   (___)   ( (_ \ /| _| | o || \_/ |
|_||_||_|\_||_|\___/ \____|\___/ \____|            \__||_||___||_n_||_| |_| 
 `
