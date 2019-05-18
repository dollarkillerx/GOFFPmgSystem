package utils

import (
	"fmt"
	"os/exec"
)

func NewUUID() (string,error) {
	out, err := exec.Command("uuidgen").Output()

	oot := fmt.Sprintf("%s", out)
	return oot,err
}

func NewUUIDSimplicity() (string,error) {
	s, e := NewUUID()
	var u string
	for _,k :=range s {
		if k != '-' {
			u = fmt.Sprintf("%s%s",u,string(k))
		}
	}
	return u,e
}