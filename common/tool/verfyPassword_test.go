package tool

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {

	fmt.Println(VerifyPassword(8, 18, "12345678bB"))

}
