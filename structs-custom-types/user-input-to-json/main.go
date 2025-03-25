package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// #2 approach 2

func main() {
	// #1 approach 1, without

	//text := getUserData("Todo text: ")
	//
	//newTodo, err := todo.New(text)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//newTodo.Show()
	//
	//err = newTodo.Save()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println("Success to save todo.")
	//
	//title := getUserData("Note title: ")
	//content := getUserData("Note content: ")
	//
	//note, err := usernote.New(title, content)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//note.Show()
	//
	//err = note.Save()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("Success to save notes.")

	// #2 approach 2, with interface
}

func getUserData(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	userData, _ := reader.ReadString('\n') // Reads until newline
	userData = strings.TrimSpace(userData) // trailing extra newline

	return userData
}
