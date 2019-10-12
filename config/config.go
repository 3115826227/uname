package config

import "unname/utils/conf"

const (
	PhoneCodeMember = "phone:code"
	DateLayout      = "2006-01-02"
	TimeLayout      = "2006-01-02 15:04:05"
)

var (
	PgURL string
)

func init() {
	PgURL = conf.Getenv("POSTGRES_URL",
		"host=localhost port=5433 user=postgres password=ps4 dbname=unname sslmode=disable")
}
