package controller

import (
	"github.com/reversTeam/nlike/kernel"
	"net/http"
)

type EchoController struct {
	kernel.Controller
}

func NewController() *EchoController {
	echo := &EchoController{
		*kernel.NewController("EchoController"),
	}
	echo.AddRoute("^/echos$", echo.List)
	echo.AddRoute("^/echos/create$", echo.Create)
	echo.AddRoute("^/echos/[0-9a-f]{8}-([0-9a-f]{4}-){3}[0-9a-f]{12}$", echo.View)
	echo.AddRoute("^/echos/[0-9a-f]{8}-([0-9a-f]{4}-){3}[0-9a-f]{12}/update$", echo.Update)
	echo.AddRoute("^/echos/[0-9a-f]{8}-([0-9a-f]{4}-){3}[0-9a-f]{12}/delete$", echo.Delete)

	return echo
}

func (o *EchoController) List(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Method of list all echos HTTP\n"))
}

func (o *EchoController) Create(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Method of view one echo HTTP\n"))
}

func (o *EchoController) View(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Method of view one echo HTTP\n"))
}

func (o *EchoController) Update(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Method of update one echo HTTP\n"))
}

func (o *EchoController) Delete(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Method of delete one echo HTTP\n"))
}
