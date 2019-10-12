package handle

import "github.com/gin-gonic/gin"

/*
	desc: 超级账号创建管理账号
	type: post
	params:
		id string
		username string
		phone string
		code string
*/

func AdminCreate(c *gin.Context) {

}

/*
	desc: 超级账号注销管理账号
	type: post
	params:
		id string
		cancel_id string
		code string
*/
func AdminCancel(c *gin.Context) {

}

/*
	desc: 管理账号初始化
	type: post
	params:
		id string
		phone string
		code string
*/

func AdminInit(c *gin.Context) {

}

/*
	desc: 管理账号账密登录
	type: post
	params:
		username string
		password string
*/

func AdminPwdLogin(c *gin.Context) {

}

/*
	desc: 管理账号手机登录
	type: post
	params:
		phone string
		code string
*/

func AdminPhoneLogin(c *gin.Context) {

}

/*
	desc: 管理账号修改密码
	type: post
	params:
		id string
		password string
		new_password string
		phone string
		code string
*/

func AdminPwdUpdate(c *gin.Context) {

}

/*
	desc: 管理账号退出登录
	type: get
	params:
		id string
*/

func AdminLogout(c *gin.Context) {

}

/*
	desc: 管理账号查看已创建的子账号列表
	type: get
	params:
		id string
*/

func AdminSubList(c *gin.Context) {

}
