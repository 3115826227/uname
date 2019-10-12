package new

import (
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"unname/model/data"
)

var (
	pgURL = "host=localhost port=5433 user=postgres password=ps4 dbname=unname sslmode=disable"
)

func init() {
	db, err := xorm.NewEngine("postgres", pgURL)
	if err != nil {
		panic(err)
	}
	db.ShowSQL(false)
	err = db.Sync2(
		new(data.AccountRoot),
	)
	db.ShowExecTime(false)
}
