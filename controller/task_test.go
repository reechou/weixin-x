package controller

import (
	"fmt"
	"github.com/jinzhu/now"
	"testing"
)

func TestTime(t *testing.T) {
	fmt.Println(now.BeginningOfDay().Unix())
}
