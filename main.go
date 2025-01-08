package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gonuts/commander"
)

const (
	toDoList = ".todoList"
)

func main() {

	filename := ""
	currToDo := false
	currDir, err := os.Getwd()

	if err == nil {
		filename = filepath.Join(currDir, toDoList)
		_, err = os.Stat(filename)
		currToDo = true
	}

	if !currToDo {
		home := os.Getenv("HOME")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		filename = filepath.Join(home, toDoList)
	}

	command := &commander.Command{
		UsageLine: os.Args[0],
		Short:     "To-Do List",
	}

	//Possible options for To-Do List Manager
	command.Subcommands = []*commander.Command{
		addTask(filename),
		listTasks(filename),
		markTask(filename),
		deleteTask(filename),
	}

	err = command.Dispatch(context.Background(), os.Args[1:])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
