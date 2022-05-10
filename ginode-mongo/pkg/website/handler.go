package website

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/yuedun/ginode-mongo/db"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

type any = interface{}

//WebsiteList列表
func WebsiteList(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	claims := jwt.ExtractClaims(c)
	userID, err := primitive.ObjectIDFromHex(claims["user_id"].(string))
	if err != nil {
		panic(err)
	}
	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil {
		offset = 0
	}
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		limit = 10
	}
	name := c.Query("name")
	category := c.Query("category")
	search := Website{
		UserID:   userID,
		Name:     name,
		Category: category,
		Status:   1,
	}
	wbService := NewService(db.NewDB("website"))
	list, total, err := wbService.GetWebsiteList(offset, limit, search)
	data := map[string]any{
		"result": list,
		"count":  total,
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    data,
	})
}

// 获取单个网站，对外提供
func GetWebsite(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	webService := NewService(db.NewDB("website"))
	websiteName := c.Query("name")
	pageURL := c.Query("url")
	fmt.Println("name:", websiteName, pageURL)
	website, page, err := webService.GetWebsite(websiteName, pageURL)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"website": website,
		"page":    page,
	})

}

//Create
func Create(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	claims := jwt.ExtractClaims(c)
	userID, err := primitive.ObjectIDFromHex(claims["user_id"].(string))
	if err != nil {
		panic(err)
	}
	websiteService := NewService(db.NewDB("website"))
	wbObj := Website{}
	c.ShouldBind(&wbObj)
	wbObj.ID = primitive.NewObjectID()
	wbObj.UserID = userID
	wbObj.CreatedAt = time.Now()
	wbObj.UpdatedAt = time.Now()
	wbObj.Status = 1
	fmt.Println(wbObj)
	err = websiteService.CreateWebsite(&wbObj)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    wbObj,
		"message": "ok",
	})
}

//Update
func Update(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	claims := jwt.ExtractClaims(c)
	userID, err := primitive.ObjectIDFromHex(claims["user_id"].(string))
	if err != nil {
		panic(err)
	}
	websiteService := NewService(db.NewDB("website"))
	website := Website{}
	c.ShouldBind(&website)
	website.UserID = userID
	fmt.Println(website)
	err = websiteService.UpdateWebsite(&website)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":    website,
			"message": "ok",
		})
	}
}

//Delete
func Delete(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	claims := jwt.ExtractClaims(c)
	userID, err := primitive.ObjectIDFromHex(claims["user_id"].(string))
	if err != nil {
		panic(err)
	}
	websiteId := c.Param("id")
	id, err := primitive.ObjectIDFromHex(websiteId)
	if err != nil {
		panic(err)
	}
	websiteService := NewService(db.NewDB("website"))
	err = websiteService.DeleteWebsite(userID, id)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})

}
