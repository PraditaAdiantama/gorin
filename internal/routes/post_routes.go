package routes

import (
	"giron/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(rg *gin.RouterGroup) {
	posts := rg.Group("/posts")
	{
		posts.GET("", handler.GetPosts)
        posts.GET("/:id", handler.GetPost)
        posts.POST("", handler.CreatePost)
        posts.PATCH("/:id", handler.UpdatePost)
        posts.DELETE("/:id", handler.DeletePost)
	}
}
