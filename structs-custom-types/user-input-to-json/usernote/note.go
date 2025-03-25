package usernote

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type UserNote struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (*UserNote, error) {
	if title == "" || content == "" {
		return nil, errors.New("field cannot be empty")
	}

	return &UserNote{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (n *UserNote) Save() error {
	noteJson, err := json.Marshal(n)
	if err != nil {
		return err
	}

	err = os.WriteFile("note.json", noteJson, 0644)
	if err != nil {
		fmt.Printf("Error on writing file: %s\n", err)
		return err
	}

	return nil
}

func (n *UserNote) Show() {
	fmt.Printf("Your note titled %s has the following content:\n\n%s\n\n", n.Title, n.Content)
}
