package utils

import (
	"fmt"
	"testing"
)

func TestGetOnlyId(t *testing.T) {
	id := GetOnlyId()
	fmt.Println(id)
}

func TestGetUUID(t *testing.T) {
	uuid := GetUUID()
	fmt.Println(uuid)
}

func TestGetSixCode(t *testing.T) {
	code := GetSixCode()
	fmt.Println(code)
}

func TestGetFourCode(t *testing.T) {
	code := GetFourCode()
	fmt.Println(code)
}
