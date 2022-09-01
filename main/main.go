package main

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"sync"

	readfile "github.com/epaitoo/mgrep/readFile"
)

func main() {
	var files []string
	var input, subStr, dir string
	var wg sync.WaitGroup
	var m sync.Mutex

	fmt.Println("Enter the command Below: \"mgrep [search_string] [search_dir]\"")	

	// check user input
	_, err := fmt.Scanln(&input, &subStr, &dir)
	if input != "mgrep" {
		fmt.Println("First input must be \"mgrep\"")
	} else if subStr == "" {
		fmt.Println("search term cannot be empty")
	} else if dir == "" {
		fmt.Println("search directory cannot be empty")
	}

	if err != nil{
		log.Panic(err)
	}

	//Recurse into directory and append path to files
	e := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		//check if the path is not a directory append to files list
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if e != nil {
		log.Println(e)
	}

	for _, file := range files {
		// run go routine here
		wg.Add(1)
		go readfile.FileReader(file, subStr, &wg, &m)
	}

	wg.Wait()

}