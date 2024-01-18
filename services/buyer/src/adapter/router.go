package adapter

import "github.com/gin-gonic/gin"

type Router struct {
	engine *gin.Engine
}

func NewRouter() *Router {
	router := new(Router)
	router.engine = gin.New()
	router.engine.Use(gin.Logger())
	return router
}

func (r *Router) Run() {
	err := r.engine.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func (r *Router) GroupRegister(path string) *gin.RouterGroup {
	return r.engine.Group(path)
}
