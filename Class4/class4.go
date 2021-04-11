package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	var busers[100] BlogUser

	//注册
	// signup?username=qwedfv&password=**********
	router.GET("/signup", func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password") // shortcut for c.Request.URL.Query().Get("lastname")
		i:=0;
		for ; i < 100; i++ {//判断用户名是否已存在
			if(busers[i].username == username){
				c.String(http.StatusOK, "%s has already exsited", username)//存在则输出已存在
				break
			}
		}
		if(i==100){//不存在则新建一个user
			j:=0;
			for ; busers[j].username!=""&&j<100; j++ {
			}
			busers[j].username = username
			busers[j].password = password
			c.String(http.StatusOK, "sign up successfully %s", username)
		}

	})

	//登录
	router.GET("/signin", func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")
		i:=0;
		for ; i<100&&busers[i].username != username; i++ {//寻找该用户名
		}
		fmt.Println(i)
		if(i<100){
			if(busers[i].password == password){//判断用户名与密码是否匹配
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
		for ; busers[i].username != username&&i<100; i++ {//寻找该用户名
		}
		if(i<100){
			if(busers[i].password == fpassword){//判断用户名与密码是否匹配
				busers[i].password = ppassword
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
		for ; busers[i].username != username&&i<100; i++ {//寻找该用户名
		}
		if(i<100){
			if(busers[i].password == password){//判断用户名与密码是否匹配
				busers[i].username=""//匹配删除数据
				busers[i].password=""
				c.String(http.StatusOK, "%s lougout successfully", username)
			}else{
				c.String(http.StatusOK, "username or password is wrong")//不匹配输出错误
			}
		}else{
			c.String(http.StatusOK, "user is not found")//找不到用户名输出not found
		}

	})

	//发布文章
	//http://localhost:8080/id/publish?title="abc"&content="12345678"
	router.GET("/id/publish", func(c *gin.Context) {
		username := c.Param("id")
		title := c.Query("title")
		content := c.Query("content")
		i:=0;
		for ; busers[i].username != username&&i<100; i++ {//寻找该用户名
		}
		if(i<100){
			busers[i].blogs[busers[i].num_blog].content = content
			busers[i].blogs[busers[i].num_blog].title = title
			busers[i].num_blog++
			c.String(http.StatusOK, "%s publish %s successfully", username,title)
		}else{
			c.String(http.StatusOK, "user is not found")//找不到用户名输出not found
		}

	})

	//点赞
	router.GET("/id/upvote/title", func(c *gin.Context) {
		username := c.Param("id")
		title := c.Param("title")
		i:=0;
		for ; busers[i].username != username&&i<100; i++ {//寻找该用户名
		}
		if(i<100){
			//if()//根据文章名或文章id找到文章{
				//Blog.num_upvote++
			c.String(http.StatusOK, "%s upvote %s successfully", username,title)
			//}else{
				//找不到该文章
			//}
		}else{
			c.String(http.StatusOK, "user is not found")//找不到用户名输出not found
		}

	})


	router.Run(":8080")
}

type BlogUser struct {
	username string
	password string
	num_blog int
	blogs[100] Blog
}

type Blog struct {
	title string
	content string
	num_upvote int
}
