package handle

import "github.com/gin-gonic/gin"

/*
	desc: 普通账号注册
	type: post
	params:
		username string
		password string
		phone string
		code string
*/

func GeneralRegister(c *gin.Context) {

}

/*
	desc: 普通账号账密登录
	type: post
	params:
 		username string
		password string
*/

func GeneralPwdLogin(c *gin.Context) {

}

/*
	desc: 普通账号手机登录
	type: post
	params:
		phone string
		code string
*/

func GeneralPhoneLogin(c *gin.Context) {

}

/*
	desc: 普通账号修改密码
	type: post
	params:
		id string
		password string
		new_password string
		phone string
		code string
*/

func GeneralPwdUpdate(c *gin.Context) {

}

/*
	desc: 普通账号退出登录
	type: post
	params:
 		id string
*/

func GeneralLogout(c *gin.Context) {

}
