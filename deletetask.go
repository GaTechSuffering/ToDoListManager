package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/gonuts/commander"
)

func deleteTask(filename string) *commander.Command {

	deleteTaskHelper := func(cmd *commander.Command, args []string) error {
		//Calls "Help/Usage" Menu when there is no argument or greater than the expected
		if len(args) == 0 || len(args) > 1 {
			cmd.Usage()
			return nil
		}

		//Converts user input to an int
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			return err
		}

		fileO, err := os.Create(filename + "_")
		if err != nil {
			return err
		}
		defer fileO.Close()

		//Attempts to open file
		fr, errr := os.Open(filename)
		if errr != nil {
			return err
		}

		//Counts how many tasks are in the list
		sr := bufio.NewScanner(fr)
		lineCount := 0
		for sr.Scan() {
			lineCount++
		}

		//Returns to the user an error if input is greater than number of tasks
		if lineCount < id {
			fmt.Println("Task does not exist")
			return err
		}
		fr.Close()

		//Attempts to open file
		f, err := os.Open(filename)
		if err != nil {
			return err
		}

		//Reads from the file and attempts to find a match
		br := bufio.NewReader(f)
		for n := 1; ; n++ {
			b, _, err := br.ReadLine()
			if err != nil {
				if err != io.EOF {
					return err
				}
				break
			}
			match := false
			if id == n {
				match = true
			}

			//If a match is found, delete. Else, do nothing.
			if match {
				fmt.Printf("Task \"%s\" was deleted.\n", string(b))
			} else {
				_, err = fmt.Fprintf(fileO, "%s\n", string(b))
				if err != nil {
					return err
				}
			}
		}

		f.Close()
		fileO.Close()
		err = os.Remove(filename)
		if err != nil {
			return err
		}
		return os.Rename(filename+"_", filename)
	}

	//Usage of Delete Task
	return &commander.Command{
		Run:       deleteTaskHelper,
		UsageLine: "delete [ID]",
		Short:     "Deletes an item within the To-Do List",
	}

}
