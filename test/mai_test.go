package test

import (
	"fmt"
	"testing"
)

func Print1to20() int {
	return 0
}

func TestPrint(t *testing.T) {
	res := Print1to20()
	fmt.Println("hey")
	if res != 210 {
		t.Errorf("Wrong result of Print1to20")
	}
}
