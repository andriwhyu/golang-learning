package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Define target directory")
		return
	}

	writerSlice := make([]byte, 0)
	for _, selectedDir := range args {
		dirEntry, err := os.ReadDir(selectedDir)

		if err != nil {
			fmt.Printf("Error when read directory: %s", err)
			return
		}

		selectedDir = selectedDir + "\n"
		writerSlice = append(writerSlice, selectedDir...)

		for _, dir := range dirEntry {
			if dir.IsDir() {
				dirName := "\t" + dir.Name() + "/\n"
				writerSlice = append(writerSlice, dirName...)
			}
		}

		writerSlice = append(writerSlice, "\n"...)
	}

	err := os.WriteFile("out.txt", writerSlice, 0644)
	if err != nil {
		fmt.Printf("Error on writing file: %s\n", err)
		return
	}

}
