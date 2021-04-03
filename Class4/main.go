package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	var users[100] User

	//注册
	// signup?username=qwedfv&password=**********
	router.GET("/signup", func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password") // shortcut for c.Request.URL.Query().Get("lastname")
		i:=0;
		for ; i < 100; i++ {//判断用户名是否已存在
			if(users[i].username == username){
				c.String(http.StatusOK, "%s has already exsited", username)//存在则输出已存在
				break
			}
		}
		if(i==100){//不存在则新建一个user
			j:=0;
			for ; users[j].username!=""&&j<100; j++ {
			}
			users[j].username = username
			users[j].password = password
			c.String(http.StatusOK, "sign up successfully %s", username)
		}

	})

	//登录
	router.GET("/signin", func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")
		i:=0;
		for ; users[i].username != username&&i<100; i++ {//寻找该用户名
		}
		if(i<100){
			if(users[i].password == password){//判断用户名与密码是否匹配
				c.String(http.StatusOK, "%s signin successfully", username)//匹配输出成功
			}else{
				c.String(http.StatusOK, "username or password is wrong")//不匹配输出错误
			}
		}else{
			c.String(http.StatusOK, "user is not found")//找不到用户名输出not found
		}

	})

	//更改密码
	router.GET("/changepassword", func(c *gin.Context) {
		username := c.Query("username")
		fpassword := c.Query("fpassword")
		ppassword := c.Query("ppassword")
		i:=0;
		for ; users[i].username != username&&i<100; i++ {//寻找该用户名
		}
		if(i<100){
			if(users[i].password == fpassword){//判断用户名与密码是否匹配
				users[i].password = ppassword
				c.String(http.StatusOK, "psaaword change successfully")//匹配更改密码并输出成功
			}else{
				c.String(http.StatusOK, "username or password is wrong")//不匹配输出错误
			}
		}else{
			c.String(http.StatusOK, "user is not found")//找不到用户名输出not found
		}

	})

	//注销
	router.GET("/logout", func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")
		i:=0;
		for ; users[i].username != username&&i<100; i++ {//寻找该用户名
		}
		if(i<100){
			if(users[i].password == password){//判断用户名与密码是否匹配
				users[i].username=""//匹配删除数据
				users[i].password=""
				c.String(http.StatusOK, "%s lougout successfully", username)
			}else{
				c.String(http.StatusOK, "username or password is wrong")//不匹配输出错误
			}
		}else{
			c.String(http.StatusOK, "user is not found")//找不到用户名输出not found
		}

	})

	router.Run(":8080")
}

type User struct {
	username string
	password string
}