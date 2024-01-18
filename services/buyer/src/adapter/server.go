package adapter

import (
	"buyer/src/dto"
	"buyer/src/query"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *Router
}

func NewServer(r *Router) *Server {
	server := new(Server)
	server.Router = r
	return server
}

func (s *Server) Routing() {
	v1 := s.Router.GroupRegister("/api/v1")

	// Add your routes here
	v1.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	v1.GET("/users", func(c *gin.Context) {
		users := query.GetUser()
		c.JSON(200, gin.H{
			"users": users,
		})
	})

	v1.POST("/user", func(c *gin.Context) {
		body := &dto.RequestUser{}
		err := c.ShouldBind(body)
		if err != nil {
			panic(err)
		}
		user := query.CreateUser(body.UserName, body.UserAge)
		c.JSON(200, gin.H{
			"age":  user.Age,
			"name": user.Name,
		})
	})

	v1.POST("/bind", func(c *gin.Context) {
		body := &dto.RequestCreateHotel{}
		err := c.ShouldBind(body)
		if err != nil {
			panic(err)
		}
		c.JSON(200, gin.H{
			"address": body.HotelAddress,
			"name":    body.HotelName,
			"zzz":     body.Dan,
		})
	})

}
