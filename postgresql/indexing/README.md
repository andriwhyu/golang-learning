# Indexing Project

## Description

This project is a practice project to apply indexing in PostgresSQL. The main functionality of this project is provide **seeding** capabilities and generate a data in `comments` table. 

In this project the `main.go` will execute the data generator based on the `GENERATED_DATA_COUNT` environment variable, with default value 5. Below will show the project structure.

### Project Structure

```
postgresql/
├── indexing/                                          # project name
│   ├── build/                                         # docker builder file
│   │   └── docker-compose.yaml
│   ├── migrations/                                    # postgres migration file
│   │   ├── 000001_create_comments_practice_up.sql
│   │   ├── 000001_create_comments_practice_down.sql
│   │   └── ...
│   ├── .envrc                                         # environment variables definition
│   ├── .envrc-sample                                  # .envrc template
│   ├── db.go                                          # establish database connnection
│   ├── seed.go                                        # seeding generator logic
│   ├── main.go                                        # main program to run the project
│   └── Makefile                                       # make operation
└── ...
```

**Environment Variables Management**

For managing environment variables, I used the `direnv` tool, which loads environment variables by reading the `.envrc` data specific to each directory. You can find a detailed explanation here: https://direnv.net/.

### Database Structure

| id | username | content                                                             |
|----|----------|---------------------------------------------------------------------|
| 1  | fiona0   | Building a blog with Go sounds fun—thanks for sharing your process! |
| 2  | teri1    | Go and CLI tools are a match made in heaven. Love the examples!     |

## How to run seeding process

You can choose one of the below steps to running the program. It depends on your purpose.

### Run all the migrations directly

This steps will run all the migrations and applied the indexing directly.

1. Initialize the postgres database using docker by running `make build`
2. Init the migrations by running `make migrate-up`
3. Run the database seeding by following command ⤵️
   ```
   # make run-generater <numdata>
   > make run-generator 10
   ```
4. Check the created index by using following command below. It should have `comments_username_idx` and `comments_content_trgm_idx` indexes. Both indexes created by the migration file.
   ```sql
   SELECT schemaname, indexname, tablename, indexdef
   FROM pg_indexes
   WHERE tablename = 'comments';
   ```
5. (Optional) As indexes improve searching duration, TBD...


## Additional features
TBD