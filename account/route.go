package account

import (
	"github.com/gin-gonic/gin"
	"unname/account/server/handle"
)

func Route(route *gin.Engine) {

	root := route.Group("/account/root")
	{
		root.POST("/register", handle.RootRegister)
		root.POST("/pwd_login", handle.RootPwdLogin)
		root.POST("/pwd_update", handle.RootPwdUpdate)
		root.POST("/phone_bind", handle.RootPhoneBind)
		root.POST("/phone_login", handle.RootPhoneLogin)
		root.GET("/logout", handle.RootLogout)
		root.GET("/admin_list", handle.RootAdminList)
		root.GET("/admin_sub_list", handle.RootAdminSubList)
		root.GET("/general_list", handle.RootGeneralList)

		root.POST("/admin_create", handle.AdminCreate)
		root.POST("/admin_cancel", handle.AdminCancel)
	}

	admin := route.Group("account/admin")
	{
		admin.POST("/init", handle.AdminInit)
		admin.POST("/pwd_login", handle.AdminPwdLogin)
		admin.POST("/phone_login", handle.AdminPhoneLogin)
		admin.POST("/pwd_update", handle.AdminPwdUpdate)
		admin.GET("/logout", handle.AdminLogout)
		admin.GET("/sub_list", handle.AdminSubList)

		admin.POST("/sub_create", handle.SubCreate)
		admin.POST("/sub_cancel", handle.SubCancel)
	}

	sub := route.Group("/account/sub")
	{
		sub.POST("/init", handle.SubInit)
		sub.POST("/pwd_login", handle.SubPwdLogin)
		sub.POST("/phone_login", handle.SubPhoneLogin)
		sub.POST("/pwd_update", handle.SubPwdUpdate)
		sub.GET("/logout", handle.SubLogout)
	}

	general := route.Group("/account/general")
	{
		general.POST("/register", handle.GeneralRegister)
		general.POST("/pwd_login", handle.GeneralPwdLogin)
		general.POST("/phone_login", handle.GeneralPhoneLogin)
		general.POST("/pwd_update", handle.GeneralPwdUpdate)
		general.GET("/logout", handle.GeneralLogout)
	}

	phone := route.Group("/account")
	{
		phone.GET("/register_code", handle.PhoneRegisterCode)
		phone.GET("/create_code", handle.PhoneCreateCode)
		phone.GET("/cancel_code", handle.PhoneCancelCode)
		phone.GET("/bind_code", handle.PhoneBindCode)
	}

}
