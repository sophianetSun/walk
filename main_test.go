package main

import (
	"strings"
	"testing"
)

func TestSuffixCheckUppercase(t *testing.T) {
	s1 := "abc.bmp"
	s2 := "abc.BMP"
	lower_s2 := strings.ToLower(s2)

	if strings.HasSuffix(s1, "bmp") != true {
		t.Error("말도 안됨")
	}
	if strings.HasSuffix(s2, "bmp") != false {
		t.Errorf("HasSuffix Not check uppercase, %s", s2)
	}
	if strings.HasSuffix(lower_s2, "bmp") != true {
		t.Errorf("lower bmp == bmp")
	}
}

func TestRegExp(t *testing.T) {
	testList := []string{
		"abc.bmp", "abc.BMP", "abc.jpg", "abc.JPG", "abc.a.png", "abc.a.PNG",
	}
	for _, fileName := range testList {
		if !CheckImageFile(fileName) {
			t.Errorf("Not Parse filename %s", fileName)
		}
	}
}
