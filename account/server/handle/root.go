package handle

import (
	"github.com/gin-gonic/gin"
	"unname/model/data"
	"net/http"
	"unname/model/response"
	"unname/utils"
	"time"
	"unname/config"
	"unname/redis"
)

/*
	desc: 超级账号账密登录
	type: post
	params:
		username string
		password string
 */

func RootRegister(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	now := time.Now().Format(config.TimeLayout)
	var root = &data.AccountRoot{
		Id:         utils.GetOnlyId(),
		Username:   username,
		Password:   password,
		CreateTime: now,
		UpdateTime: now,
	}

	if root.Insert() == false {
		c.JSON(http.StatusOK, response.Response{
			Message: "插入失败",
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		Data: root.Result(),
	})
}

func RootPwdLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var root = &data.AccountRoot{
		Username: username,
		Password: password,
	}
	if root.Exist() == false {
		c.JSON(http.StatusOK, response.Response{
			Message: "数据不存在",
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		Data: root.Result(),
	})
}

/*
	desc: 超级账号密码修改
	type: post
	params:
		id  string
		password string
		new_password string
		phone string
		code string
 */

func RootPwdUpdate(c *gin.Context) {

}

/*
	desc: 超级账号手机绑定
	type: post
	params:
		id string
		phone string
		code string
 */

func RootPhoneBind(c *gin.Context) {
	id := c.PostForm("id")
	phone := c.PostForm("phone")

	var root = &data.AccountRoot{
		Id: id,
	}
	if root.Exist() == false {
		c.JSON(http.StatusOK, response.Response{
			Message: "账号id不存在",
		})

		return
	}

	code := c.PostForm("code")

	rdsCode := redis.HashGet(config.PhoneCodeMember, phone)

	if rdsCode != code {
		c.JSON(http.StatusOK, response.Response{
			Message: "手机验证码不正确",
		})
		return
	}

	root.Phone = phone
	if root.Update() == false {
		c.JSON(http.StatusOK, response.Response{
			Message: "手机信息更新失败",
		})

		return
	}

	c.JSON(http.StatusOK, response.Response{
		Data: root.Result(),
	})

}

/*
	desc: 超级账号手机登录
	type: post
	params:
		phone string
		code string
 */

func RootPhoneLogin(c *gin.Context) {

}

/*
	desc: 超级账号退出登录
	type: get
	params:
		id string
 */

func RootLogout(c *gin.Context) {

}

/*
	desc: 超级账号查看已创建的管理账号列表
	type: get
	params:
		id string
 */

func RootAdminList(c *gin.Context) {

}

/*
	desc: 超级账号查看管理账号下的子账号列表
	type: get
	params:
		id string
		admin_id string
 */
func RootAdminSubList(c *gin.Context) {

}

/*
	desc: 超级账号查看已创建的普通账号列表
	type: get
	params:
		id string
 */

func RootGeneralList(c *gin.Context) {

}
