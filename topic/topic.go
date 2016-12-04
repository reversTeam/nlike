package nlike

import (
	"github.com/reversTeam/nlike/server"
)

type Topic struct {
	Controller
}

func NewTopic() *Topic {
	topic := &Topic{
		Controller{},
	}
	topic.AddRoute("^/topics", topic.List)
	topic.AddRoute("^/topics/create", topic.Create)
	topic.AddRoute("^/topics/[0-9a-f]{8}-([0-9a-f]{4}-){3}[0-9a-f]{12}$", topic.View)
	topic.AddRoute("^/topics/[0-9a-f]{8}-([0-9a-f]{4}-){3}[0-9a-f]{12}$/update", topic.Update)
	topic.AddRoute("^/topics/[0-9a-f]{8}-([0-9a-f]{4}-){3}[0-9a-f]{12}$/delete", topic.Delete)

	return topic
}

func (o *Topic) List(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Method of list all topics HTTP\n"))
}

func (o *Topic) Create(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Method of view one topic HTTP\n"))
}

func (o *Topic) View(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Method of view one topic HTTP\n"))
}

func (o *Topic) Update(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Method of update one topic HTTP\n"))
}

func (o *Topic) Delete(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Method of delete one topic HTTP\n"))
}
