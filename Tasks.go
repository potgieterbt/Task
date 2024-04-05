package main

import (
	"bufio"
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

func addTask(path string, task string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(task + "\n"); err != nil {
		panic(err)
	}
	return err
}

func removeTask(path string, Index int) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	temp, err := os.Create("temp.txt")
	if err != nil {
		return err
	}

	defer temp.Close()

	scanner := bufio.NewScanner(f)
	i := 1
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if i != Index {
			fmt.Fprintln(temp, scanner.Text())
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	f.Close()

	f, err = os.Create("test.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	err = os.Rename("temp.txt", "./test.txt")
	if err != nil {
		return err
	}

	return err
}

func listTasks(path string) error {

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	i := 1
	for scanner.Scan() {
		fmt.Println(i, scanner.Text())
		i++
	}

	if err := scanner.Err(); err != nil {
		return err
	}
    if i == 1 {
        fmt.Println("Task List is empty")
    }
	return err
}

func fileInit(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		f.Close()
	}
	return nil
}

func main() {
	var fpath string
	var task string
	var taskidx int

	flag.StringVarP(&task, "add", "a", "", "Task input")
	flag.IntVarP(&taskidx, "remove", "r", -1, "Index of task to be removed")
	listPtr := flag.BoolP("list", "l", false, "a bool")
	flag.StringVarP(&fpath, "file", "f", "./test.txt", "File to use for Tasks")

	flag.Parse()

	err := fileInit(fpath)
	if err != nil {
		panic(err.Error())
	}

	if *listPtr {
		err = listTasks(fpath)
		if err != nil {
			panic(err.Error())
		}
	}

	if taskidx != -1 {
		err := removeTask(fpath, taskidx)
		if err != nil {
			panic(err.Error())
		}
		err = listTasks(fpath)
		if err != nil {
			panic(err.Error())
		}
	}

	if len(task) > 0 {
		addTask(fpath, task)
		err = listTasks(fpath)
		if err != nil {
			panic(err.Error())
		}
	}

    if len(os.Args) < 2 {
		err = listTasks(fpath)
		if err != nil {
			panic(err.Error())
		}
    }

}
