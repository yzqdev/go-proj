package page

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/yuedun/ginode-mongo/db"
	"github.com/yuedun/ginode-mongo/pkg/component"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

type any = interface{}

// List 列表
func List(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil {
		offset = 0
	}
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		limit = 10
	}
	websiteID := c.Query("websiteID")
	title := c.Query("title")
	url := c.Query("url")
	id, err := primitive.ObjectIDFromHex(websiteID)
	if err != nil {
		panic(err)
	}
	search := Page{
		WebsiteID: id,
		Title:     title,
		URL:       url,
	}
	wbService := NewService(db.NewDB("website"))
	list, total, err := wbService.GetPageList(offset, limit, search)
	if err != nil {
		panic(err)
	}
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

// GetPage 获取单个网站
func GetPage(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	webService := NewService(db.NewDB("website"))
	name := c.Query("name")
	fmt.Println("url:", name)
	page, err := webService.GetPage(name)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": page,
	})

}

//Create 创建页面
func Create(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	pageService := NewService(db.NewDB("website"))
	newObj := Page{}
	c.ShouldBind(&newObj)
	//newObj.ID = primitive.NewObjectID()
	websiteID := c.Param("websiteId")
	webID, err := primitive.ObjectIDFromHex(websiteID)
	if err != nil {
		panic(err)
	}
	newObj.WebsiteID = webID
	newObj.CreatedAt = time.Now()
	newObj.UpdatedAt = time.Now()
	newObj.Status = 1
	fmt.Println(newObj)
	err = pageService.CreatePage(&newObj)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    newObj,
		"message": "ok",
	})
}

//Update 修改页面
func Update(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	pageService := NewService(db.NewDB("website"))
	page := Page{}
	c.ShouldBind(&page)
	err := pageService.UpdatePage(&page)
	if err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":    page,
			"message": "ok",
		})
	}
}

//Delete 删除页面
func Delete(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	pageID := c.Param("id")
	id, err := primitive.ObjectIDFromHex(pageID)
	if err != nil {
		panic(err)
	}
	pageService := NewService(db.NewDB("website"))
	err = pageService.DeletePage(id)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

//GetPageComponents 获取页面组件
func GetPageComponents(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	pageID := c.Param("id")
	id, err := primitive.ObjectIDFromHex(pageID)
	if err != nil {
		panic(err)
	}
	pageService := NewService(db.NewDB("website"))
	components, err := pageService.GetPageComponents(id)
	if err != nil {
		panic(err)
	}
	fmt.Println("components", components)
	c.JSON(http.StatusOK, gin.H{
		"data":    components,
		"message": "ok",
	})

}

// UpdatePageComponents 单独修改网站组件
func UpdatePageComponents(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	pageID := c.Param("id")
	id, err := primitive.ObjectIDFromHex(pageID)
	var pageComponents []component.Component
	c.ShouldBind(&pageComponents)
	fmt.Println(pageComponents)
	pageService := NewService(db.NewDB("website"))
	err = pageService.UpdatePageComponents(id, pageComponents)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// CopyPage 复制网站（页面），多个页面组成一个网站
func CopyPage(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	pageID := c.Param("id")
	id, err := primitive.ObjectIDFromHex(pageID)
	if err != nil {
		panic(err)
	}
	url := c.Param("url")
	pageService := NewService(db.NewDB("website"))
	err = pageService.CopyPage(id, url)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
