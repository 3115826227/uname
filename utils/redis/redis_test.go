package redis

import (
	"fmt"
	"testing"
	"unname/config"
)

func TestHashExist(t *testing.T) {
	ok := HashExist("ids", "43001570")
	fmt.Println(ok)
}

func TestHashAdd(t *testing.T) {
	HashAdd(config.PhoneCodeMember, "456", "15688220367")
}

func TestHashGet(t *testing.T) {
	str := HashGet(config.PhoneCodeMember, "456")
	fmt.Println(str)
}
