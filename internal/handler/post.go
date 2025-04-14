package handler

import (
	"giron/internal/dto"
	"giron/internal/model"
	"giron/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	posts, err := repository.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, posts)
}

func GetPost(c *gin.Context) {
	id := c.Param("id")

	post, err := repository.GetOnePost(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

    if post == nil {
        c.JSON(http.StatusNotFound, "Not found")
        return
    }

	c.JSON(http.StatusOK, post)
}

func CreatePost(c *gin.Context) {
	var post dto.CreatePostDTO

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	result, err := repository.CreatePost(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, result)
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")

    post, err := repository.GetOnePost(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, err.Error())
        return
    }

    if post == nil {
        c.JSON(http.StatusNotFound, "Not found")
        return
    }

	var body dto.UpdatePostDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
        return
	}

    var rlt *model.Post

	result, err := repository.UpdatePost(id, &body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
        return
	}

	c.JSON(http.StatusOK, result.Decode(rlt))
}

func DeletePost(c *gin.Context){
    id := c.Param("id")

    _, err := repository.DeletePost(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(204, nil)
}
