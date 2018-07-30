package main

import (
	"testing"
	"fmt"
)


type F interface {
	hello() string
}

type A struct {
	Name string
}

func (a *A) hello() string {
	return a.Name
}

type B struct {
	Name string
}

func (a *B) hello() string {
	return a.Name
}

func Test_Inter(t *testing.T){
	ha:=A{
		Name:"aaa",
	}
	hb :=B{
		Name:"bbbb",
	}
	fmt.Printf("a:%s\nb:%s\n",wa(ha),wa(hb))
}

func wa(n interface{}) string{
	switch n.(type) {
	case A:
		return n.(A).Name
	case B:
		return "this is b"
	default:
		return "wocao"
	}
}