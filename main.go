package main

import (
	"github.com/gin-gonic/gin"
	"unname/account"
	_ "unname/utils/pg/db"
	_ "unname/utils/pg/new"
)

func main() {
	c := gin.Default()

	account.Route(c)

	c.Run(":7050")
}
