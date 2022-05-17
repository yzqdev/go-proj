package model

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
)

type Company struct {
	BaseModel
	Name   string
	UserId string
}

func QueryClearUserAndCompany() map[string]interface{} {
	db := GetDb()
	db.Exec("delete from sys_users")
	db.Exec("delete from companies")
	var sqlCompany []Company
	var sqlUser []SysUser
	db.Find(&sqlUser)
	db.Find(&sqlCompany).Association("Company")

	return gin.H{
		"success": true, "sysUser": sqlUser,
		"company": sqlCompany,
	}

}
func SetCom() map[string]interface{} {
	coms := []Company{
		{
			Name: "111",
		},
		{
			Name: "222",
		},
	}
	color.Redln(coms)
	sysU := SysUser{
		Name: "第一个user",

		Coms: coms,
	}
	db := GetDb()
	db.Create(&sysU)
	db.Save(&sysU)
	var sqlCompany Company
	var sqlUser SysUser
	err := db.Model(&sqlUser).Association("Coms").Find(&sqlCompany)

	if err != nil {
		color.Redln("添加错误")
		color.Redln(err)
		return gin.H{
			"success": false,
			"sysUser": sqlUser,
		}
	}
	return gin.H{
		"success": true, "sysUser": sqlUser,
	}

}
