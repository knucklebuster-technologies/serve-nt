package main

import (
	"os"
	"os/exec"
)

func startMongod(dbpath string) (*exec.Cmd, error) {
	ensureDbPath(dbpath)
	mongod := exec.Command("mongod", "--dbpath", dbpath)
	err := mongod.Start()
	if err != nil {
		return nil, err
	}
	return mongod, nil
}

func ensureDbPath(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_ = os.Mkdir(path, 0777)
	}
}
