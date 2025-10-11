package main

import (
	"database/sql"
	"fmt"
	"math/rand/v2"
)

var (
	randomUsername = []string{
		"alice", "bob", "charlie", "dana", "edward",
		"fiona", "george", "hannah", "ian", "julia",
		"ken", "lisa", "mike", "nina", "oliver",
		"paula", "quentin", "rachel", "sam", "tina",
		"umar", "victoria", "wilson", "xena", "yuri",
		"zara", "alex", "bella", "cameron", "daphne",
		"evan", "faith", "gerald", "heather", "isaac",
		"jenny", "kyle", "laura", "mark", "natalie",
		"owen", "penny", "rebecca", "steve", "teri",
		"ulysses", "vanessa", "walt", "xavier", "zoe",
	}

	userComments = []string{
		"Great read! I’ve started using time-blocking and it really works.",
		"This was super helpful for someone like me just getting started with Kubernetes.",
		"Building a blog with Go sounds fun—thanks for sharing your process!",
		"Technical debt is something we all ignore until it’s too late. Thanks for the reminder.",
		"Perfect explanation for REST APIs. I’ll use this to teach my interns!",
		"I’ve been curious about PromQL. This gave me a good push to learn it.",
		"Debugging distributed systems is a nightmare—these tips are gold.",
		"I use 3 of these tools already, definitely going to try the others.",
		"Go and CLI tools are a match made in heaven. Love the examples!",
		"This article opened my eyes to how vulnerable our microservices are. Great tips!",
		"GitOps always confused me before, but your post made it very clear.",
		"Inspiring journey. I’m thinking of starting a 100-days-of-code challenge too!",
		"Very practical SRE practices. We’re going to try implementing error budgets.",
		"Just what I needed to get started with OpenTelemetry. Thanks!",
		"The rate limiting examples really helped me understand the theory.",
		"Clean code is underrated. I shared this with my whole team.",
		"Super insightful take on observability trends. Love the eBPF mention!",
		"I had no idea how powerful `context` could be in Go. Thanks for the demo.",
		"This article made me rethink how I design APIs. Great advice!",
		"Exactly what I needed to scale our logs efficiently. Appreciate the Fluent Bit section!",
	}
)

type comments struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

type Seed struct {
	db        *sql.DB
	totalData int
}

func NewSeed(db *sql.DB, totalData int) *Seed {
	return &Seed{db: db, totalData: totalData}
}

func (s *Seed) RunSeeding() error {
	generatedComments := s.commentGenerator()

	fmt.Printf("Generating %d data...\n", s.totalData)
	for _, comment := range generatedComments {
		err := s.createComments(comment)
		if err != nil {
			return err
		}

		fmt.Printf("Successfully create comment for user %s\n", comment.Username)
	}
	return nil
}

func (s *Seed) commentGenerator() []*comments {
	generatedComments := make([]*comments, 0, s.totalData)

	for i := 0; i < s.totalData; i++ {
		username := fmt.Sprintf("%s%d", randomUsername[rand.IntN(len(randomUsername))], i)
		userComment := userComments[rand.IntN(len(userComments))]

		comment := &comments{Username: username, Content: userComment}
		generatedComments = append(generatedComments, comment)
	}

	return generatedComments
}

func (s *Seed) createComments(comment *comments) error {
	query := `
		INSERT INTO comments(username, content)
		VALUES ($1, $2);
	`

	_, err := s.db.Exec(query, comment.Username, comment.Content)
	if err != nil {
		return err
	}

	return nil
}
