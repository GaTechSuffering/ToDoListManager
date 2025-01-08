package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gonuts/commander"
)

func addTask(filename string) *commander.Command {

	addTaskHelper := func(cmd *commander.Command, args []string) error {
		//Calls "Help/Usage" Menu when there is no argument
		if len(args) == 0 {
			cmd.Usage()
			return nil
		}

		//Opens To-Do list file to read/write from
		fileO, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			return err
		}
		defer fileO.Close()

		//Takes in the task and displays to the user
		task := strings.Join(args, " ")
		_, err = fmt.Fprintln(fileO, task)
		fmt.Printf("Task \"%s\" was added.", task)
		return err
	}

	//Usage of Add Task
	return &commander.Command{
		Run:       addTaskHelper,
		UsageLine: "add [message]",
		Short:     "Adds an item to the To-Do List",
	}
}
