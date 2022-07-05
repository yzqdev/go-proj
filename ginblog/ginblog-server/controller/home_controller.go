package controller

import (
	"ginblog/model"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetArticleList(c *gin.Context) {

	articles := model.QueryArticleList()
	utils.JSON(c, http.StatusOK, "成功", articles)
}

// Req 使用protobuf
func Req(c *gin.Context) {
	req := &model.SearchRequest{
		Query:         "123",
		PageNumber:    1,
		ResultPerPage: 22,
	}

	c.ProtoBuf(http.StatusOK, req)
}
