package db

import (
	"context"
	"fmt"
	"github.com/andriwhyu/golang-learning/bewg-learning/project-2/internal/store"
	"log"
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

	firstNames = []string{
		"Alice", "Bob", "Charlie", "Dana", "Edward",
		"Fiona", "George", "Hannah", "Ian", "Julia",
		"Ken", "Lisa", "Mike", "Nina", "Oliver",
		"Paula", "Quentin", "Rachel", "Sam", "Tina",
		"Umar", "Victoria", "Wilson", "Xena", "Yuri",
		"Zara", "Alex", "Bella", "Cameron", "Daphne",
		"Evan", "Faith", "Gerald", "Heather", "Isaac",
		"Jenny", "Kyle", "Laura", "Mark", "Natalie",
		"Owen", "Penny", "Rebecca", "Steve", "Teri",
		"Ulysses", "Vanessa", "Walt", "Xavier", "Zoe",
	}

	lastNames = []string{
		"Smith", "Johnson", "Williams", "Brown", "Jones",
		"Miller", "Davis", "Garcia", "Martinez", "Hernandez",
		"Lopez", "Gonzalez", "Wilson", "Anderson", "Thomas",
		"Taylor", "Moore", "Jackson", "Martin", "Lee",
		"Perez", "Thompson", "White", "Harris", "Sanchez",
		"Clark", "Ramirez", "Lewis", "Robinson", "Walker",
		"Young", "Allen", "King", "Wright", "Scott",
		"Torres", "Nguyen", "Hill", "Flores", "Green",
		"Adams", "Nelson", "Baker", "Hall", "Rivera",
		"Campbell", "Mitchell", "Carter", "Roberts", "Gomez",
	}

	postTitles = []string{
		"10 Tips to Boost Your Productivity Today",
		"Understanding Kubernetes for Beginners",
		"How I Built a Blog Using Go and Markdown",
		"The Hidden Costs of Technical Debt",
		"A Beginner’s Guide to RESTful APIs",
		"Why You Should Learn PromQL in 2025",
		"Debugging Distributed Systems: Lessons Learned",
		"Top 5 Monitoring Tools for Modern DevOps",
		"From Zero to CLI App in Go",
		"How to Secure Your Microservices Architecture",
		"Deploying Applications with GitOps",
		"What I Learned After 100 Days of Coding",
		"Improving System Reliability with SRE Practices",
		"Using OpenTelemetry to Trace Your Services",
		"Understanding Rate Limiting with Practical Examples",
		"Writing Clean Code: Habits and Mindsets",
		"The Future of Observability in Cloud-Native Environments",
		"Exploring the Power of Context in Go",
		"Common Pitfalls in API Design (And How to Avoid Them)",
		"Scaling Log Ingestion with Fluent Bit and Amazon S3",
	}

	postContents = []string{
		"Struggling with procrastination? Here are 10 actionable tips to stay focused and get more done every day, from setting clear goals to using time-blocking methods.",
		"Kubernetes can be intimidating at first. In this guide, we'll demystify core concepts like Pods, Deployments, and Services so you can confidently deploy your first containerized app.",
		"I built a static blog engine using Go and parsed content from Markdown files. This post shares the architecture, libraries used, and challenges I faced.",
		"Ignoring tech debt can slow teams and introduce bugs. Learn to identify, prioritize, and manage it before it affects your product’s quality and team morale.",
		"Learn the basics of building RESTful APIs: endpoints, HTTP methods, status codes, and JSON responses. Perfect for new backend developers.",
		"Prometheus is everywhere, and knowing how to write queries with PromQL is a valuable skill. Here's why it's worth learning and how to get started.",
		"Distributed systems are hard to debug. This post walks through real-world examples, tips, and tools that helped me track down elusive issues.",
		"Explore five essential monitoring tools DevOps teams use daily: Prometheus, Grafana, Datadog, Loki, and New Relic, and how they complement each other.",
		"Go is perfect for building CLI tools. This post walks through creating one from scratch using Cobra, Viper, and Go's built-in libraries.",
		"Security often takes a backseat in microservices. Learn how to implement mTLS, OAuth2, and zero-trust principles for safer communication between services.",
		"GitOps simplifies Kubernetes deployments by treating Git as the source of truth. Learn the tools and workflows to adopt GitOps today.",
		"After 100 days of consistent coding, I gained more than just skills. Here are the habits, lessons, and surprises from the journey.",
		"Site Reliability Engineering brings engineering practices to ops. Learn how to apply SLIs, SLOs, and error budgets to improve uptime and user satisfaction.",
		"OpenTelemetry is the future of observability. See how to instrument Go services and visualize traces with Grafana or Jaeger.",
		"Rate limiting protects your services. This post covers token buckets, leaky buckets, and how to implement them in Go.",
		"Clean code isn’t just pretty—it’s maintainable. Learn practical habits to improve readability, reduce bugs, and keep your team happy.",
		"As systems grow complex, observability must evolve. Explore trends like eBPF, distributed tracing, and AI-driven insights.",
		"The context package is critical for timeouts and cancellations in Go. Learn how to use it properly in HTTP handlers and goroutines.",
		"Bad APIs hurt users. Avoid common mistakes like inconsistent naming, leaking internal errors, and ignoring versioning best practices.",
		"Need to handle large log volumes? See how Fluent Bit and Amazon S3 can build a scalable, low-cost log pipeline with minimal ops overhead.",
	}

	postComments = []string{
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

	postTags = []string{
		"golang",
		"kubernetes",
		"devops",
		"productivity",
		"api",
		"rest",
		"cli",
		"tracing",
		"observability",
		"prometheus",
		"promql",
		"sre",
		"microservices",
		"gitops",
		"security",
		"distributed-systems",
		"openTelemetry",
		"logging",
		"clean-code",
		"software-engineering",
	}
)

const (
	numData = 10
)

func Seed(store store.Storage) error {
	ctx := context.Background()
	users := generateUsers(numData)

	for _, user := range users {
		err := store.Users.Create(ctx, user)
		if err != nil {
			log.Println("Error creating user:", err)
			return err
		}
	}

	posts := generatePosts(numData, users)

	for _, post := range posts {
		err := store.Posts.Create(ctx, post)
		if err != nil {
			log.Println("Error creating post:", err)
			return err
		}
	}

	return nil
}

func generateUsers(userCount int) []*store.User {
	users := make([]*store.User, 0, userCount)
	for i := 0; i < userCount; i++ {
		username := fmt.Sprintf("%s%d", randomUsername[i%len(randomUsername)], i)
		users = append(users, &store.User{
			Email:     fmt.Sprintf("%s@example.com", username),
			FirstName: firstNames[i%len(firstNames)],
			LastName:  lastNames[i%len(lastNames)],
			Username:  username,
			Password:  fmt.Sprintf("%s%d", username, i+1),
		})
	}

	return users
}

func generatePosts(postCount int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, 0, postCount)
	randTagsId := rand.IntN(len(postTags))

	for i := 0; i < postCount; i++ {
		posts = append(posts, &store.Post{
			UserID:  users[rand.IntN(len(users))].ID,
			Title:   postTitles[i%len(postTitles)],
			Content: postContents[i%len(postContents)],
			Tags: []string{
				postTags[randTagsId],
				postTags[len(postTags)-1-randTagsId],
			},
		})
	}

	return posts
}
