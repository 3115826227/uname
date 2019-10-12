package utils

import (
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"math/rand"
	"strconv"
	"time"
	"unname/utils/log"
	"unname/utils/redis"
)

var (
	end          = 100000000
	start        = 10000000
	OnlyIdMember = "ids"
)

func GetOnlyId() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := strconv.Itoa(r.Intn(end-start) + start)
	if redis.HashExist(OnlyIdMember, num) {
		return GetOnlyId()
	}
	redis.HashAdd(OnlyIdMember, num, time.Now().String())
	return num
}

func GetUUID() string {
	UUID, err := uuid.NewV4()
	if err != nil {
		log.Logger.Warn("uuid generate failed", zap.String("err", err.Error()))
		return GetUUID()
	}
	return UUID.String()
}

func GetSixCode() string {
	startCode, endCode := 100000, 1000000
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := strconv.Itoa(r.Intn(endCode-startCode) + startCode)
	return num
}

func GetFourCode() string {
	startCode, endCode := 1000, 10000
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := strconv.Itoa(r.Intn(endCode-startCode) + startCode)
	return num
}
