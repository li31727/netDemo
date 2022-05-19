package gee

import (
	"net/http"
)

type  Engine struct {

	router *router
}

type HandlerFunc func ( *Context)




func New()*Engine{

	return &Engine{router: newRouter()}
}

func (engine *Engine)addRoute(method string,pattern string ,handler HandlerFunc){
	engine.router.addRoute(method,pattern,handler)
}



func (engine *Engine)GET(pattern string, handler HandlerFunc ){
	//这里将路由与handler绑定
	engine.addRoute("GET",pattern,handler)

}

func (engine *Engine)POST(pattern string, handler HandlerFunc){
	//这里将路由与handler绑定
	engine.addRoute("POST",pattern,handler)

}


func (engine *Engine)Run(port string)(err error){
	return http.ListenAndServe(port,engine)  //这里执行路由对应的Handler
}

func(engine *Engine)ServeHTTP(w http.ResponseWriter,req *http.Request){
	//
	c:=newContext(w,req)
	engine.router.handle(c)


}