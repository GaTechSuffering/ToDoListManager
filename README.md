# Project Overview

A basic command-line To-Do List Manager

## Usage

Functionality Included:
    ☐  Add Task        Adds an item to the To-Do List
    ☐  List Task       Lists all items on the To-Do List
    ☐  Mark Task       Marks an item to the To-Do List as done
    ☐  Delete Task     Deletes an item within the To-Do List

## List Task

    .\ToDoListManager.exe list [options]
          
    ☐ Task 1: feed cat
    ☐ Task 2: feed dog

## Add Task

    .\ToDoListManager.exe add [message]

    .\ToDoListManager.exe add get groceries
    .\ToDoListManager.exe list
    ☐ Task 1: feed cat
    ☐ Task 2: feed dog
    ☐ Task 3: get groceries

## Mark Task

    .\ToDoListManager.exe mark [ID]

    .\ToDoListManager.exe mark 3
    Task "get groceries" was marked done.
    .\ToDoListManager.exe list
    ☐ Task 1: feed cat
    ☐ Task 2: feed dog
    ☑ Task 3: (Completed) get groceries

## Delete Task

    .\ToDoListManager.exe delete [ID]

    .\ToDoListManager.exe delete 1
    Task "(Completed) feed cat" was deleted.
    .\ToDoListManager.exe list    
    ☐ Task 1: feed dog
    ☑ Task 2: (Completed) get groceries

### Build and Run Application

go get github.com/tachristine/todo