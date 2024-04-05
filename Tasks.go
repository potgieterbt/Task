package main

import (
	"bufio"
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

// "io"

func addTask(path string) error {
	err := os.WriteFile(path, []byte("this is a test"), 0644)
	return err
}

func removeTask(path string, Index int) error {
	// Working idea:
	//      Read line by line into a temp file until reach Index value
	//      Skip line == Index value and continue to read into temp file
	//      Rename temp file to original file
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
			fmt.Fprintln(temp, scanner.Text()) // print values to f, one per line
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
	i := 0
	for scanner.Scan() {
		fmt.Println(i, scanner.Text())
		i++
	}

	if err := scanner.Err(); err != nil {
		return err
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
	}

	fmt.Println("Add: ", task)
	fmt.Println("Remove: ", taskidx)
	fmt.Println("List: ", *listPtr)
	fmt.Println("File: ", fpath)
}
