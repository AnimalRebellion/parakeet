package proxy

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var localns *Server

func TestApi(r *gin.RouterGroup, ns *Server) error {
	localns = ns
	r.POST("/person", func(c *gin.Context) {
		var person Person
		var err error
		if err = c.ShouldBindJSON(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = localns.Send("test", &person)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		// queryParams := c.Request.URL.Query()
		c.JSON(http.StatusOK, gin.H{"status": "Message sent!"})
	})

	return localns.Receive("test")

	// r.GET("/events/:id", func(c *gin.Context) {
	// 	id, _ := strconv.Atoi(c.Param("id"))
	// 	c.JSON(http.StatusOK, id)
	// })
}
