package iface

import "testing"

func TestShowEface(t *testing.T){
	ShowEface(42)
	ShowEface("hello")
	ShowEface([]int{1,2,3})
}