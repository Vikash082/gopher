package main

import (
	"fmt"
	"os"
	"path/filepath"
	"task/cmd"
	"task/db"

	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	terror(db.Init(dbPath))
	terror(cmd.Execute())
}

func terror(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
