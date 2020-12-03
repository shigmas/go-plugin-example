package main

import (
	"fmt"
	"pluggie/plug"
	"strings"
)

type (
	PlugLower struct {
	}
)

var _ plug.MyPlug = (*PlugLower)(nil)

// Factory that we expect every plugin to have
func CreatePlug() plug.MyPlug {
	return &PlugLower{}
}

func (a *PlugLower) Init() {
	fmt.Println("Initialize PlugLower")
}

func (b *PlugLower) DoPlug(arg string) string {
	return strings.ToLower(arg)
}
