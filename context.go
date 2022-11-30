package gocore

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	AsHTML func(int, string)
	AsJSON func(int, interface{})
	Write  func(string)
}

type Context struct {
	Res    *Response
	Req    *http.Request
	Start  *time.Time
	End    *time.Time
	Status int
}

type H map[string]interface{}

func createContext(w http.ResponseWriter, r *http.Request) *Context {
	start := time.Now()
	resp := &Response{}
	c := &Context{
		Req:   r,
		Res:   resp,
		Start: &start,
	}
	c.Res.Write = func(str string) {
		w.WriteHeader(c.Status)
		w.Write([]byte(str))
	}
	c.Res.AsHTML = func(code int, str string) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		c.Status = code
		resp.Write(str)
	}
	c.Res.AsJSON = func(code int, object interface{}) {
		str, err := json.Marshal(object)
		if err != nil {
			c.Status = http.StatusInternalServerError
			resp.Write("Internal Server Error")
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		c.Status = code
		resp.Write(string(str))
	}
	return c
}
