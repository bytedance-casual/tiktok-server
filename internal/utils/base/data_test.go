package base

import (
	"fmt"
	"testing"
)

func TestIsDefaultValue(t *testing.T) {
	x := "4234"
	fmt.Println(IsDefaultValue(x))
}
