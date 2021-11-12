package main

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"

	"github.com/shigmas/go-plugin-example/pkg/plug"
)

func getPlug(plugPath string) (plug.MyPlug, error) {
	pluginLib, e := plugin.Open(plugPath)
	if e != nil {
		return nil, e
	}
	s, e := pluginLib.Lookup("CreatePlug")
	if e != nil {
		fmt.Println("Plugin does not implement CreatePlug")
		return nil, e
	}

	return s.(func() plug.MyPlug)(), nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("%s <plugin directory containing a .plugin file> <test string>\n", os.Args[0])
		os.Exit(1)
	}

	subs, err := os.ReadDir(os.Args[1])
	if err != nil {
		fmt.Println("Unable to read ", os.Args[1])
		os.Exit(1)
	}

	plugins := make([]plug.MyPlug, 0)
	for _, sub := range subs {
		plugName := filepath.Join(os.Args[1], sub.Name())
		if filepath.Ext(plugName) == ".plugin" {
			plug, err := getPlug(plugName)
			if err != nil {
				fmt.Printf("Unable to load %s: %v\n", plugName, err)
				continue
			}
			plugins = append(plugins, plug)
		}
	}

	for _, p := range plugins {
		p.Init()
	}

	for _, p := range plugins {
		fmt.Println(p.DoPlug(os.Args[2]))
	}
}
