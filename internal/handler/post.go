package handler

import (
	"giron/internal/model"
	"giron/internal/repository"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
    posts, err := repository.GetAllPosts()
    if err != nil{
        c.JSON(500, "Something went wrong")
    }

	c.JSON(200, posts)
}

func GetPost(c *gin.Context){
    id := c.Param("id")
    
    post,err := repository.GetOnePost(id)
    if err != nil{
        c.JSON(500, "Something went wrong")
    }

    c.JSON(200, post)
}

func CreatePost(c *gin.Context){
    var post model.Post
    if err := c.BindJSON(&post); err != nil {
        c.JSON(500, err)
    }

    result, err := repository.CreatePost(&post)
    if err != nil {
        c.JSON(500, err)
    }

    c.JSON(201, result)
}
