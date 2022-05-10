package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	yaml "gopkg.in/yaml.v3"
)

//profile variables
type Conf struct {
	MongoURL string `yaml:mongoURL`
}

func GetConf(filename string) (*Conf, error) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var c = new(Conf)
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return c, nil
}

/**
 * md5加密
 */
func GetMD5(password string) string {
	Md5Inst := md5.New()
	Md5Inst.Write([]byte(password))
	Result := Md5Inst.Sum(nil)
	// 以下两种输出结果一样
	fmt.Println("格式化>>>>>>>%x", Result)
	fmt.Println("hex解码>>>>>>>", hex.EncodeToString(Result))
	return fmt.Sprintf("%x", Result)
}

/**
 * 生成密码
 */
func GeneratePassword(mobile string) string {
	b := []byte(mobile)
	p := b[7:]
	password := "hello" + string(p)
	return GetMD5(password)
}

/**
 * 将结构体转为bson.M，去除空字段，避免零值
 * search必须是查询的表的结构体，否则映射出的字段可能不是表中存在的字段
 */
func Query(search interface{}) (condition bson.M, err error) {
	//将结构体转为字节数组，userInfo中的字段根据需要设置值，需要保证没有值时不会有默认值出现
	userbyte, err := bson.Marshal(search)
	if err != nil {
		return nil, err
	}
	//将字节码转为bson.M类型
	bson.Unmarshal(userbyte, &condition)
	log.Println("结构转bson：", condition)
	return condition, nil
}
