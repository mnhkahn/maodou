package supervisor

import (
	"fmt"
	"net/http"
	"time"

	"github.com/tylerb/graceful"

	. "github.com/mnhkahn/maodou/logs"
)

type SupervisorController struct {
	engine *http.ServeMux
}

func NewSupervisorController() *SupervisorController {
	controller := new(SupervisorController)
	controller.engine = http.NewServeMux()
	return controller
}

func (this *SupervisorController) Route(http_method, route string, function func(w http.ResponseWriter, req *http.Request)) {
	this.engine.HandleFunc(route, function)
}

func (this *SupervisorController) Run(ip string, port int, graceful_timeout int) {
	ColorLog("[INFO] Listening and serving HTTP on %s...\n", fmt.Sprintf("%s:%d", ip, port))
	graceful.Run(fmt.Sprintf("%s:%d", ip, port), time.Duration(graceful_timeout)*time.Second, this.engine)
}
