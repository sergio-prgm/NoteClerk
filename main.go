package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Printf("Hello Working\n")
	// content := []byte("Contenido\nMuy bueno")
	date := time.Now()
	filedate := date.Format("2006_01_02")
	titlename := date.Format("2006-Jan-02")

	filename := fmt.Sprintf("%s.md", filedate)
	f, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	w := bufio.NewWriter(f)

	_, err = w.WriteString(titlename)

	if err != nil {
		panic(err)
	}
	w.Flush()
}
