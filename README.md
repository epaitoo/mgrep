# mgrep

 A very simple grep clone that can do simple substring searching within files
 
 ## Project Features
 -  Search through the files for a substring match with [goroutines](https://go.dev/tour/concurrency/1)
 -  Display matches to the terminal as they are found
 -  It auto-recurse into subdirectories to also find matches
 -  Use of Go [WaitGroups](https://gobyexample.com/waitgroups) and [Mutexes](https://gobyexample.com/mutexes) (Synchronization methods) to make sure all files are searched
 
 ## Getting Started
 These instructions will get you a copy of the project up and running on your local machine.
 
 - Click on the 'Clone or download' button or select 'Download Zip.'
 - Move to the project directory using the terminal 
 - Navigate the main directory and run the application with the `go run main.go` command
 - To perform a search use the command: `mgrep [SUBSTR_TO_SEARCH] [SEARCH_DIRECTORY]`
