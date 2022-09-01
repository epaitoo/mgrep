package readfile

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)


func FileReader(fileName string, subStr string, wg *sync.WaitGroup, m *sync.Mutex) {
	time.Sleep(5 * time.Millisecond)
	m.Lock()
	defer m.Unlock()
	defer wg.Done()

	//open a file
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error Opening file:", err)
	}

	//Close file after everything
	defer f.Close()

	line := 1


	// scan through file line by line
	// check if text match word in a file
	s := bufio.NewScanner(f)
	for s.Scan() {
		// check if the subString match the text in each line
		if strings.Contains(s.Text(), subStr) {
			fmt.Println()
			fmt.Printf("Line number: %d \n", line)
			fmt.Printf("filepath: %v \n", fileName)
			fmt.Printf("Complete line: %s \n",s.Text())
		}
		line++
	}
	
}