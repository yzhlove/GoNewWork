package main

import (
	"fmt"
	"go.uber.org/dig"
	"net/http"
)

//dig 使用

type Config struct {
	Host string
	Port string
	Used bool
}

type Req struct {
	dig.Out

	ConfList []Config `group:"server,flatten"`
}

func NewConfig(host, port string) func() Req {
	return func() Req {
		req := Req{}
		req.ConfList = append(req.ConfList, Config{Host: host, Port: port})
		return req
	}
}

type Group struct {
	dig.In

	ConfList []Config `group:"server"`
}

func (g Group) Next() (Config, bool) {
	for i, sev := range g.ConfList {
		if !sev.Used {
			g.ConfList[i].Used = true
			return sev, true
		}
	}
	return Config{}, false
}

type Server struct {
	conf Config
}

func NewServer(conf Config) *Server {
	return &Server{conf: conf}
}

func (s *Server) Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("服务的地址为=> %s:%s", s.conf.Host, s.conf.Port)))
	})
	fmt.Println("listener on address:", s.conf)
	go http.ListenAndServe(s.conf.Host+":"+s.conf.Port, mux)
	return nil
}

func main() {

	container := dig.New()
	container.Provide(NewConfig("127.0.0.1", "50051"))
	container.Provide(NewConfig("127.0.0.1", "50052"))
	container.Provide(NewConfig("127.0.0.1", "50053"))

	err := container.Invoke(func(g Group) error {
		for range g.ConfList {
			c, ok := g.Next()
			if ok {
				if err := NewServer(c).Run(); err != nil {
					return err
				}
			}
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	ch := make(chan struct{})
	<-ch
}
