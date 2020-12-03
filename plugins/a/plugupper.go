package main

import (
	"fmt"
	"pluggie/plug"
	"strings"
)

type (
	PlugUpper struct {
	}
)

func init() {
	fmt.Println("Upper registered")
}

var _ plug.MyPlug = (*PlugUpper)(nil)

// Factory that we expect every plugin to have
func CreatePlug() plug.MyPlug {
	return &PlugUpper{}
}

func (a *PlugUpper) Init() {
	fmt.Println("Initialize PlugUpper")
}

func (a *PlugUpper) DoPlug(arg string) string {
	return strings.ToUpper(arg)
}
