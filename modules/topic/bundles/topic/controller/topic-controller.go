package controller

import (
	"github.com/reversTeam/nlike/kernel"
	"net/http"
)

type TopicController struct {
	kernel.Controller
}

func NewController() *TopicController {
	topic := &TopicController{
		kernel.Controller{},
	}
	topic.AddRoute("^/topics$", topic.List)
	topic.AddRoute("^/topics/create$", topic.Create)
	topic.AddRoute("^/topics/[0-9a-f]{8}-([0-9a-f]{4}-){3}[0-9a-f]{12}$", topic.View)
	topic.AddRoute("^/topics/[0-9a-f]{8}-([0-9a-f]{4}-){3}[0-9a-f]{12}/update$", topic.Update)
	topic.AddRoute("^/topics/[0-9a-f]{8}-([0-9a-f]{4}-){3}[0-9a-f]{12}/delete$", topic.Delete)

	return topic
}

func (o *TopicController) List(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Method of list all topics HTTP\n"))
}

func (o *TopicController) Create(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Method of view one topic HTTP\n"))
}

func (o *TopicController) View(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Method of view one topic HTTP\n"))
}

func (o *TopicController) Update(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Method of update one topic HTTP\n"))
}

func (o *TopicController) Delete(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Method of delete one topic HTTP\n"))
}
