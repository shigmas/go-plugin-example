package main

import (
	"fmt"
	"path/filepath"
	// I don't remember why this works to find the local plug package
	"pluggie/plug"
	"plugin"
)

const (
	pluginPath = "plugins"
	testString = "HelloWorld"
)

func getPlug(plugRel string) plug.MyPlug {
	pPath := filepath.Join(pluginPath, plugRel)
	var p *plugin.Plugin
	var s plugin.Symbol
	var e error
	if p, e = plugin.Open(pPath); e != nil {
		panic(e)
	}
	if s, e = p.Lookup("CreatePlug"); e != nil {
		panic(e)
	}

	return s.(func() plug.MyPlug)()
}

func main() {
	aPlug := getPlug("a/a.plugin")
	bPlug := getPlug("b/b.plugin")

	aPlug.Init()
	bPlug.Init()

	fmt.Println(aPlug.DoPlug(testString))
	fmt.Println(bPlug.DoPlug(testString))
}
