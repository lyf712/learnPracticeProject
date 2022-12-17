package test

import (
	"fmt"
	"testing"
)

func TestNewcoder(t *testing.T) {
	s := "nowcoder"
	a := 0
	switch s {
	case "nowcoder":
		a++
		fallthrough
	case "haha":
		a++
		fallthrough
	default:
		a++
	}
	fmt.Println(a)
}
