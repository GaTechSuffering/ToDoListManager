package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gonuts/commander"
)

const (
	ballotBoxUnchecked = "\u2610"
	ballotBoxChecked   = "\u2611"
)

func listTasks(filename string) *commander.Command {

	listTasksHelper := func(cmd *commander.Command, args []string) error {
		//Opens the file and attempts to read from it
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()

		//Read from file
		br := bufio.NewReader(f)
		n := 1
		for {
			b, _, err := br.ReadLine()
			if err != nil {
				if err != io.EOF {
					return err
				}
				break
			}
			line := string(b)

			//If the line is preceded by "(Completed)", add unicode for a checked ballot box
			//Else, add unicode for an unchecked ballot box
			if strings.HasPrefix(line, "(Completed)") {
				fmt.Printf("%s Task %01d: %s\n", ballotBoxChecked, n, strings.TrimSpace(line))
			} else {
				fmt.Printf("%s Task %01d: %s\n", ballotBoxUnchecked, n, strings.TrimSpace(line))
			}
			n++
		}
		return nil
	}

	flg := *flag.NewFlagSet("list", flag.ExitOnError)
	flg.Bool("fin", false, "Task not done")

	//Usage of List Tasks
	return &commander.Command{
		Run:       listTasksHelper,
		UsageLine: "list [options]",
		Short:     "Lists all items on the To-Do List",
		Flag:      flg,
	}
}
