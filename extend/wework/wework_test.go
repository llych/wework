package wework

import (
	"fmt"
	"testing"
	"wework/extend/conf"
)

func TestToken(t *testing.T) {
	conf.Setup()
	fmt.Println(Token())
}
