package user

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuedun/ginode-mongo/db"
)

// GetUserInfo 根据用户名获取用户信息
func GetUserInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	username := c.Param("username")
	userService := NewService(db.NewDB("website"))
	user, err := userService.GetUserInfoByName(username)
	if err != nil {
		fmt.Println("err:", err)
	}
	user.Avatar = "default.jpg"
	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"message": "ok",
	})
}

// GetUserList 获取用户列表
func GetUserList(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	//username := c.Param("username")
	searchObj := User{}
	userService := NewService(db.NewDB("website"))
	user, count, err := userService.GetUserList(0, 20, searchObj)
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"count":   count,
		"message": "ok",
	})
}

//CreateUser
func CreateUser(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	userService := NewService(db.NewDB("website"))
	user := User{}
	if err := c.ShouldBind(&user); err != nil {
		panic(err)
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	err := userService.CreateUser(&user)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"message": "ok",
	})
}

//UpdateUser post json
func UpdateUser(c *gin.Context) {
	userService := NewService(db.NewDB("website"))
	var user User
	userID, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(userID)
	//user.Addr = c.PostForm("addr")
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data":    nil,
			"message": "err",
		})
	} else {
		err := userService.UpdateUser(&user)
		if err != nil {
			fmt.Println("err:", err)
		}
		c.JSON(http.StatusOK, gin.H{
			"data":    user,
			"message": "ok",
		})
	}
}

//DeleteUser
func DeleteUser(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))
	userService := NewService(db.NewDB("website"))
	err := userService.DeleteUser(userID)
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
