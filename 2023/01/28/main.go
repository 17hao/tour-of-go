package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	workDir, _ := os.Getwd()

	f, err := os.Open(fmt.Sprintf("%s/go.mod", workDir))
	if err != nil {
		log.Fatal(err)
	}

	//reader := bufio.NewReader(f)
	//for {
	//	line, err := reader.ReadString('\n')
	//	fmt.Print(line)
	//	if err != nil {
	//		if errors.Is(err, io.EOF) {
	//			break
	//		}
	//		log.Fatal(err)
	//	}
	//}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
