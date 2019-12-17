package main

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("Usage: %s pdf myapp C:\\myapp.exe\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	extension := os.Args[1]
	appName := os.Args[2]
	exePath := os.Args[3]

	sfaKey, err := registry.OpenKey(registry.CLASSES_ROOT, "SystemFileAssociations", registry.CREATE_SUB_KEY)
	if err != nil {
		panic(err)
	}
	defer sfaKey.Close()

	extKey, _, err := registry.CreateKey(sfaKey, "."+extension, registry.CREATE_SUB_KEY)
	if err != nil {
		panic(err)
	}
	defer extKey.Close()

	shellKey, _, err := registry.CreateKey(extKey, "shell", registry.CREATE_SUB_KEY)
	if err != nil {
		panic(err)
	}
	defer shellKey.Close()

	appKey, _, err := registry.CreateKey(shellKey, appName, registry.CREATE_SUB_KEY)
	if err != nil {
		panic(err)
	}
	defer appKey.Close()

	commandKey, _, err := registry.CreateKey(appKey, "command", registry.WRITE)
	if err != nil {
		panic(err)
	}
	defer commandKey.Close()

	commandKey.SetStringValue("", exePath+` "%1"`)
}
