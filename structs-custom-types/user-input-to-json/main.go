package main

import (
	"bufio"
	"example.com/user-input-to-json/usernote"
	"fmt"
	"os"
	"strings"
)

func main() {
	title := getUserData("Note title: ")
	content := getUserData("Note content: ")

	note, err := usernote.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	note.Show()

	err = note.Save()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Success to save notes.")
}

func getUserData(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	userData, _ := reader.ReadString('\n') // Reads until newline
	userData = strings.TrimSpace(userData) // trailing extra newline

	return userData
}
