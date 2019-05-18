package utils

import (
	"strings"
	"testing"
)

func TestNewUUID(t *testing.T) {
	s, e := NewUUID()
	if e!=nil {
		t.Error(e.Error())
	}
	t.Log(s)
}

func TestNewUUIDSimplicity(t *testing.T) {
	s, e := NewUUIDSimplicity()
	if e!=nil {
		t.Error(e.Error())
	}
	t.Log(s)
}

func TestMd5String(t *testing.T) {
	s := Md5String("123")
	s = strings.ToUpper(s)
	if s != "202CB962AC59075B964B07152D234B70" {
		t.Error("非正常",s)
	}
}

func TestGetCurrentTime(t *testing.T) {
	time := GetCurrentTime()
	t.Log(time)
}

func TestGetTimeToString(t *testing.T) {
	s, e := GetTimeToString("1558057058")
	if e!=nil {
		t.Error(e.Error())
	}
	t.Log(s)
}

func TestGetTimeStringToTime(t *testing.T) {
	s, e := GetTimeStringToTime("2019-05-17")
	if e!=nil {
		t.Error(e.Error())
	}
	t.Log(s)
}