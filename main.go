package main

import (
	"fmt"
	"os"
)

func main() {

	err := saveText("file.txt", "Hello, World")
	if err != nil {
		fmt.Println("Done!")
	}
}

func saveText(name, text string) error {

	f, err := os.Create(name)
	if err != nil {
		return err
	}

	_, err = f.WriteString(text)
	if err != nil {
		return err
	}

	return nil
}
