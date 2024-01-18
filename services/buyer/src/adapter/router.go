package adapter

import "github.com/gin-gonic/gin"

type Router struct {
	Engine *gin.Engine
}

func NewRouter() *Router {
	router := new(Router)
	router.Engine = gin.New()
	router.Engine.Use(gin.Logger())
	return router
}

func (r *Router) Run() {
	err := r.Engine.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func (r *Router) GroupRegister(path string) *gin.RouterGroup {
	return r.Engine.Group(path)
}
