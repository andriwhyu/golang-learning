package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) (*Todo, error) {
	if text == "" {
		return nil, errors.New("todo list is empty")
	}

	return &Todo{
		Text: text,
	}, nil
}

func (t *Todo) Save() error {
	todoJson, err := json.Marshal(t)
	if err != nil {
		return err
	}

	fileName := "todo.json"

	err = os.WriteFile(fileName, todoJson, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (t *Todo) Show() {
	fmt.Println(t.Text)
}
