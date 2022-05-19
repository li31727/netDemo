package gee

import (
	"fmt"
	"net/http"
)

type  Engine struct {

	router map[string]HandlerFunc
}

type HandlerFunc func (w http.ResponseWriter,req *http.Request)




func New()*Engine{

	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine)addRoute(method string,pattern string ,handler HandlerFunc){
	key:=method+"-"+pattern
	engine.router[key]=handler
}



func (engine *Engine)GET(pattern string, handler func(http.ResponseWriter, *http.Request)){
	//这里将路由与handler绑定
	engine.addRoute("GET",pattern,handler)

}

func (engine *Engine)POST(pattern string, handler func(http.ResponseWriter, *http.Request)){
	//这里将路由与handler绑定
	engine.addRoute("POST",pattern,handler)

}


func (engine *Engine)Run(port string)(err error){
	return http.ListenAndServe(port,engine)  //这里执行路由对应的Handler
}

func(engine *Engine)ServeHTTP(w http.ResponseWriter,req *http.Request){
	//
	key:=req.Method+"-"+req.URL.Path
	if handler,ok:=engine.router[key];ok{
		handler(w,req)
	}else{
		fmt.Fprintf(w,"404 NOT FOUND: %s\n",req.URL)
	}


}