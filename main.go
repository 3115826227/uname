package main

import (
	"github.com/gin-gonic/gin"
	"unname/account"
	_ "unname/pg/db"
	_ "unname/pg/new"
)

func main() {
	c := gin.Default()

	account.Route(c)

	c.Run(":7050")
}
