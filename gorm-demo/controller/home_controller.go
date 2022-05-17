package controller

import (
	"github.com/gin-gonic/gin"
	"gorm-demo/model"
	"gorm-demo/utils"
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

func SetSysUser(c *gin.Context) {
	msg := model.SetCom()
	utils.JSON(c, http.StatusOK, "success", msg)
}
func ClearUseAndCompany(c *gin.Context) {
	flag := model.QueryClearUserAndCompany()
	utils.JSON(c, http.StatusOK, "success", flag)
}
