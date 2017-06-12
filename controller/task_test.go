package controller

import (
	"testing"
	"github.com/jinzhu/now"
	"fmt"
)

func TestTime(t *testing.T) {
	fmt.Println(now.BeginningOfDay().Unix())
}
