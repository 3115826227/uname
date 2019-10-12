package data

import (
	"go.uber.org/zap"
	"unname/log"
	. "unname/pg/db"
)

type AccountRoot struct {
	Id         string `xorm:"pk"`
	Username   string
	Password   string
	Phone      string
	Status     int
	CreateTime string
	UpdateTime string
}

type AccountRootResponse struct {
	Id         string `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"-"`
	Phone      string `json:"-"`
	Status     int    `json:"-"`
	CreateTime string `json:"-"`
	UpdateTime string `json:"-"`
}

func (root *AccountRoot) Result() AccountRootResponse {
	if root.Id <= "0" {
		return AccountRootResponse{}
	}
	return AccountRootResponse(*root)
}

func (root *AccountRoot) Insert() bool {
	_, err := DB.Insert(root)
	if err != nil {
		log.Logger.Warn("root insert failed",
			zap.String("err", err.Error()))
		return false
	}
	return true
}

func (root *AccountRoot) Update() bool {
	_, err := DB.Id(root.Id).Update(root)
	if err != nil {
		log.Logger.Warn("root update failed",
			zap.String("err", err.Error()))
		return false
	}
	return true
}

func (root *AccountRoot) Exist() bool {
	ok, err := DB.Get(root)
	if err != nil {
		log.Logger.Warn("root exist failed",
			zap.String("err", err.Error()))
		return false
	}
	return ok
}
