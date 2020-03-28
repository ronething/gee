// author: ashing
// time: 2020/3/28 11:45 上午
// mail: axingfly@gmail.com
// Less is more.

package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (e *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	// TODO: 重复路由判断？
	e.router[key] = handler // 赋值
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 Not Found: %s\n", r.URL)
	}
}

func (e *Engine) Run(addr string) (err error) {
	for k, v := range e.router {
		fmt.Println(k, v)
	}
	return http.ListenAndServe(addr, e)
}
