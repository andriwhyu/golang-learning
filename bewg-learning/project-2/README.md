# Project 2

This document is my scratch note to learn and analyze the code.

## Prerequisite

1. docker-compose
2. Makefile
3. golang-migrate
4. direnv

app.store.Posts.Create()
|    |     |      |
|    |     |      |
|    |     |      |----------> Create function of posts
|    |     |-----------------> an interface from store struct which implement Create function
|    |-----------------------> field of application struct 
|----------------------------> main application struct

internal/store/storage.go -> store the main storage configuration
internal/store/posts.go -> struct of Post which implement interface defined in storage.go
internal/store/users.go -> struct of User which implement interface defined in storage.go

Package that can help working with database/sql:
- gorm
- sqlx
- sqlboiler 

UI for database connection:
- tableplus

dbConfig parameters/fields:
1. Address = full postgres address; `postgresql://<username>:<password>@<address>:<port>/<db_name>?sslmode=disable`
2. Max Open Connections = number of open connections (running + idle) connections in postgres.
3. Max Idle connections = number of idle connections that can be hanging in postgres.
4. Max idle time = how long that connection can be hanging until it disconnected.

Tools for migration:
- migrate https://github.com/golang-migrate/migrate
- goose https://github.com/pressly/goose

In this project I used [golang-migrate](https://github.com/golang-migrate/migrate?tab=readme-ov-file#migrate) as the migration tools. Golang migrate will handle the migration by creating two files, `.up` and `.down`

Postgres data types explanation:
1. `citext`. `citext` is not default extension for case-insensitive usage. This type will help if you've a field that case-insensitive like email. It should install first.
2. `bytea`. `bytea` is byte array to store a byte for example storing video, pdf, hash.

In this project everything about the API was store under `cmd/api/api.go`. This file contain the struct definition of config, and application. also it contain function to running the server (`mount()`, `run()`)

List of return error:
- 404 : data not found
- 400 : bad request
- 500 : internal error

The difference between http PUT and PATCH
- PUT is used for replace the entire field. Example ⤵️
```
# existing field
{
    "name": "abdul"
    "email": "abdul@email.com"
    "job": "sales"
}

# incoming request body
{
    "name": "abdullah"
}

# result
{
    "name": "abdullah"
    "email": null
    "job": null
}
```

- PATCH is used for replace partial field. Example ⤵️
```
# existing field
{
    "name": "abdul"
    "email": "abdul@email.com"
    "job": "sales"
}

# incoming request body
{
    "name": "abdullah"
}

# result
{
    "name": "abdullah"
    "email": "abdul@email.com"
    "job": "sales"
}
```

Context is immutable. That's mean to utilize the context we need to create new context with value

Concurrency Control
Concurrency control mainly used to handle race condition in database. There are two kind of concurrency control; optimistic and pessimistic concurrency control. 
1. Optimistic control is a mechanism control based on idea that concurrency is unlikely to happen. By this, the same transaction happen for the same row is allow until the step to update the DBMS data. 
    It success the update process once the version that want to be update still the same as the DBMS. If it's difference, it will conflict and fail. 
2. Pessimistic control is a mechanism control based on idea that concurrency is likely happen. By this, once there is a transaction occur, it will lock the row (no one can read/update it) until it release.

Optimistic control used version
Pessimistic control used lock

Why context is helpful?
1. With context, we can set a timeout. It means, when a process is run, e.g. database operation, it can cancel the operation if it's takes too long.
    We can save resources with this.

In Makefile, I usually used `.PHONY` command. This command is used to say to the make file that if the Make command is equal to the list in `.PHONY` then you don't need 
to look up for a file name, just do the recipe.
Pattern of make file
```
<target>: <dependencies>
    <recipe>
```

In HTTP request, every methods have its own characteristic. Some characteristics are; safe, idempotent, cacheable
**Safe**
Safe is mean the HTTP method not change the server state. All safe method is idempotent, but not all idempotent method is safe. 

**Idempotent**
Idempotent mean when the HTTP method executed it will return the same how many times it triggered. 
https://developer.mozilla.org/en-US/docs/Glossary/Idempotent

**Cacheable**
https://developer.mozilla.org/en-US/docs/Glossary/Cacheable

### Database Indexing
One of advantage of indexing is increased the Read performance. Although, the trade-off is slower Write performance.


LEARNING NOTES:
1. During implementation of Seed script, I noticed my teacher using i%len for indexing. Turns out this is effective to order result for beginning to end of len.
    It will be useful for indexing start from 0, since the result would be 0 ~ len-1.
    For example looping for 20 times from 10 data. It will sort the output as 1 - 10  
    0-10: first cycle of loop it wil generate number 0, 1, ..., 9
    1-10: second cycle of loop it wil generate number 0, 1, ..., 9
2. Composite primary key is a primary key that created from multiple columns. This approach chosen when we need set unique pairs for the columns value.

IMPROVEMENT NOTES:
1. Adjust the migrate makefile argument parser. Instead of using `make migrate-create <args>`, use `make migrate-create MIGRATION_NAME=<args>`. In additions of it, add validation if MIGRATION_NAME is empty.