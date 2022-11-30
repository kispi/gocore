package gocore

import (
	"fmt"
	"net/http"
	"time"
)

func log(c *Context) {
	fmt.Printf("\n[%v] %v %v %v %v",
		time.Now().Format("2006-01-02 15:04:05"),
		c.Status,
		c.Req.Method,
		c.Req.URL.Path,
		time.Since(*c.End),
	)
}

func routeGenerator(httpMethod string, path string, handler func(c *Context)) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		c := createContext(w, r)
		if r.Method != httpMethod {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
		} else {
			handler(c)
		}
		t := time.Now()
		c.End = &t
		log(c)
	})
}

func (s *Server) RouteGet(path string, handler func(c *Context)) {
	routeGenerator(http.MethodGet, path, handler)
}

func (s *Server) RoutePost(path string, handler func(c *Context)) {
	routeGenerator(http.MethodPost, path, handler)
}

func (s *Server) RoutePut(path string, handler func(c *Context)) {
	routeGenerator(http.MethodPut, path, handler)
}

func (s *Server) RouteDelete(path string, handler func(c *Context)) {
	routeGenerator(http.MethodDelete, path, handler)
}
