package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

const (
	appURL = "http://localhost:8080"
)

type responseResult struct {
	Data Post `json:"data"`
}

type updatePostReq struct {
	Title   string `json:"title" validate:"omitempty,max=255"`
	Content string `json:"content" validate:"omitempty,max=1000"`
}

type Post struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Tags      []string `json:"tags"`
	UserID    int      `json:"user_id"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

func updatePost(postID int, reqData updatePostReq, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get("http://localhost:8080/v1/posts/1")
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var respResult responseResult
	err = json.Unmarshal(body, &respResult)
	if err != nil {
		panic(err)
	}

	patchURL := fmt.Sprintf("%s/v1/posts/%d", appURL, postID)

	if reqData.Title == "" {
		reqData.Title = respResult.Data.Title
	}

	if reqData.Content == "" {
		reqData.Content = respResult.Data.Content
	}

	jsonStr, _ := json.Marshal(reqData)

	req, err := http.NewRequest(http.MethodPatch, patchURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}

	respUpdate, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	fmt.Println("response Status:", respUpdate.Status)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	newTitle := "New Title"
	newContent := "New Content"
	postID := 1

	go updatePost(postID, updatePostReq{Title: newTitle, Content: ""}, &wg)
	go updatePost(postID, updatePostReq{Title: "", Content: newContent}, &wg)
	wg.Wait()
}
