package new

import (
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"unname/config"
	"unname/utils/log"
)

var (
	DB *xorm.Engine
)

func init() {
	var err error
	DB, err = xorm.NewEngine("postgres", config.PgURL)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Logger.Warn(err.Error())
	}

}
