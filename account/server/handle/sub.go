package handle

import "github.com/gin-gonic/gin"

/*
	desc: 管理账号创建子账号
	type: post
	params:
		id string
		username string
		phone string
		code string
*/

func SubCreate(c *gin.Context) {

}

/*
	desc: 管理账号注销子账号
	type: post
	params:
		id string
		cancel_id string
		phone string
		code string
*/

func SubCancel(c *gin.Context) {

}

/*
	desc: 子账号初始化
	type: post
	params:
 		id string
		phone string
		code string
*/

func SubInit(c *gin.Context) {

}

/*
	desc: 子账号账密登录
	type: post
	params:
		username string
		password string
*/

func SubPwdLogin(c *gin.Context) {

}

/*
	desc: 子账号手机登录
	type: post
	params:
		phone string
		code string
*/

func SubPhoneLogin(c *gin.Context) {

}

/*
	desc: 子账号修改密码
	type: post
	params:
		id string
		password string
		new_password string
		phone string
		code string
*/

func SubPwdUpdate(c *gin.Context) {

}

/*
	desc: 子账号退出登录
	type: get
	params:
		id string
*/

func SubLogout(c *gin.Context) {

}
