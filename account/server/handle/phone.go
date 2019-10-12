package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"unname/common/model/response"
	"unname/config"
	"unname/utils"
	"unname/utils/redis"
)

/*
	desc: 账号注册码
	type: post
	params:
 		phone
*/

func PhoneRegisterCode(c *gin.Context) {
	phone := c.Query("phone")
	code := utils.GetSixCode()
	redis.HashAdd(config.PhoneCodeMember, phone, code)
	c.JSON(http.StatusOK, response.Response{
		Data: response.PhoneCodeResponse{
			Message:   "OK",
			RequestID: utils.GetUUID(),
			Code:      code,
		},
	})
}

/*
	desc: 账号创建码
	type: post
	params:
		phone string
*/

func PhoneCreateCode(c *gin.Context) {

}

/*
	desc: 账号注销码
	type: post
	params:
		phone string
*/

func PhoneCancelCode(c *gin.Context) {

}

/*
	desc: 账号绑定码
	type: post
	params:
		phone string
*/

func PhoneBindCode(c *gin.Context) {
	phone := c.Query("phone")
	code := utils.GetSixCode()
	redis.HashAdd(config.PhoneCodeMember, phone, code)
	c.JSON(http.StatusOK, response.Response{
		Data: response.PhoneCodeResponse{
			Message:   "OK",
			RequestID: utils.GetUUID(),
			Code:      code,
		},
	})
}
