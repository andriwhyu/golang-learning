package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("required a dir name, exit...")
		return
	}

	directory := args[0]
	dirEntry, err := os.ReadDir(directory)

	if err != nil {
		fmt.Printf("Error when read directory: %s", err)
		return
	}

	// note: we can utilize the code, by define the capacity of the slices. It will save more resources (CPU & Memory) when allocated new slices.
	// there are 2 approach: heuristics and optimal approach.
	// heuristics approach will be calculated by get the maximum length of file + 1 new line, and multiply it by total file; cap = (255 + 1) * total_file.
	// optimal approach will be calculated by get length of each file + 1 new line; cap += (len(file) + 1)

	// heuristics approach.
	// calculate number of empty file
	//var totalEmptyFile int
	//for _, dir := range dirEntry {
	//	var fileInfo fs.FileInfo
	//
	//	fileInfo, err = dir.Info()
	//	if err != nil {
	//		fmt.Printf("Error when read the file info: %s", err)
	//		return
	//	}
	//
	//	if fileInfo.Size() == 0 {
	//		totalEmptyFile++
	//	}
	//}
	//fileNames := make([]byte, 0, (255+1)*totalEmptyFile)

	// optimal approach.
	// calculate length of empty file
	var fileCap int
	for _, dir := range dirEntry {
		var fileInfo fs.FileInfo

		fileInfo, err = dir.Info()
		if err != nil {
			fmt.Printf("Error when read the file info: %s", err)
			return
		}

		if fileInfo.Size() == 0 {
			name := fileInfo.Name()
			fileCap += len([]rune(name)) + 1
		}
	}

	fileNames := make([]byte, 0, fileCap)
	for _, dir := range dirEntry {
		var fileInfo fs.FileInfo

		fileInfo, err = dir.Info()
		if err != nil {
			fmt.Printf("Error when read the file info: %s", err)
			return
		}

		if fileInfo.Size() == 0 {
			name := fileInfo.Name()
			fileNames = append(fileNames, name...)
			fileNames = append(fileNames, "\n"...)
		}
	}

	fmt.Printf("total bytes: %d\n", cap(fileNames))

	// 0644 - 0-user-group-other
	err = os.WriteFile("out.txt", fileNames, 0644)
	if err != nil {
		fmt.Printf("Error on writing file: %s\n", err)
		return
	}

}
