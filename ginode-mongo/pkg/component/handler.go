package component

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/yuedun/ginode-mongo/db"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

type any = interface{}

//ComponentList列表
func ComponentList(c *gin.Context) {
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
	componentSearch := Component{
		Name:     name,
		Category: category,
		Status:   1,
	}
	cmService := NewService(db.NewDB("website"))
	list, total, err := cmService.GetComponentList(offset, limit, componentSearch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	data := map[string]any{
		"result": list,
		"count":  total,
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    data,
	})
}

//get component
func GetComponent(c *gin.Context) {
	componentId := c.Param("id")
	componentService := NewService(db.NewDB("website"))
	id, err := primitive.ObjectIDFromHex(componentId)
	component, err := componentService.GetComponent(id)
	fmt.Println(component)
	if err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":    component,
			"message": "ok",
		})
	}
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
	componentService := NewService(db.NewDB("website"))
	comObj := Component{}
	err := c.ShouldBind(&comObj)
	if err != nil {
		panic(err)
	}
	comObj.ID = primitive.NewObjectID()
	nowTime := time.Now()
	comObj.CreatedAt = &nowTime
	comObj.UpdatedAt = &nowTime
	comObj.Status = 1
	err = componentService.CreateComponent(&comObj)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    comObj,
		"message": "ok",
	})
}

//Update
func Update(c *gin.Context) {
	componentService := NewService(db.NewDB("website"))
	component := Component{}
	if err := c.ShouldBind(&component); err != nil {
		fmt.Println("err shouldbind:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err := componentService.UpdateComponent(component)
	if err != nil {
		fmt.Println("err update:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":    component,
			"message": "ok",
		})
	}
}

//Delete
func Delete(c *gin.Context) {
	componentId := c.Param("id")
	id, err := primitive.ObjectIDFromHex(componentId)
	if err != nil {
		panic(err)
	}
	componentService := NewService(db.NewDB("website"))
	err = componentService.DeleteComponent(id)
	if err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}
