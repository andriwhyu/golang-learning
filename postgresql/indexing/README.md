# Indexing Project

## Description

This project is a practice project to apply indexing in PostgresSQL. The main functionality of this project is provide **seeding** capabilities and generate a data in `comments` table. 

In this project, the `main.go` will execute the data generator based on the `GENERATED_DATA_COUNT` environment variable, with default value 5. Below will show the project structure.

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

### Table Structure

The table contains three columns with the data type as below ⤵️

```
Table comments {
  id       integer [primary key, increment]
  username varchar(50)
  content  text [not null]
}
```

Table sample:

| id  | username    | content                                                             |
|-----|-------------|---------------------------------------------------------------------|
| 1   | fiona0      | Building a blog with Go sounds fun—thanks for sharing your process! |
| 2   | teri1       | Go and CLI tools are a match made in heaven. Love the examples!     |

### Indexing

In this project, it uses two indices:
1. GIN with pg_trgm module for text searching by regex. It applied for **content** column.
2. B-Tree indexing for standard equal search. It applied for **username** column

## How to run seeding process

You can choose one of the below steps to running the program. It depends on your purpose.

### Run all the migrations directly

This steps will run all the migrations and applied the indexing directly.

1. Initialize the postgres database using docker by running `make build`
2. Init the migrations by running `make migrate-up`
3. Run the database seeding by following command ⤵️
   ```
   # make run-generater <numdata>
   
   make run-generator 10
   ```
4. Check the created index by using following command below. It should have `comments_username_idx` and `comments_content_trgm_idx` indexes. Both indexes created by the migration file.
   ```sql
   SELECT schemaname, indexname, tablename, indexdef
   FROM pg_indexes
   WHERE tablename = 'comments';
   ```
5. (Optional) As indexes improve search duration, you can analyse the query execution time using the following command.
   ```sql
   EXPLAIN ANALYZE SELECT * FROM comments WHERE username = 'faith3265';
   EXPLAIN ANALYZE SELECT * FROM comments WHERE content LIKE '%with my%';
   ```
6. (Optional) Clean up phase.
   ```
   make migrate-down
   make clean
   ```
   
#### Command Recap

```
make build
make migrate-up
make run-generator 10

SELECT schemaname, indexname, tablename, indexdef
FROM pg_indexes
WHERE tablename = 'comments';

EXPLAIN ANALYZE SELECT * FROM comments WHERE username = 'faith3265';
EXPLAIN ANALYZE SELECT * FROM comments WHERE content LIKE '%with my%';

# optional
make migrate-down
make clean
```

### Run the migration periodically

This step is designed to compare execution times before and after indexing. Please follow the guide below.

1. Initialize the postgres database using docker by running `make build`
2. Run the migration by executing `make migrate-up 1`. This will migrate the database step by step (one level up). In the first step, it will create the `comment` database.
3. To seed the database, execute the following command. ⤵️
   ```
   # make run-generater <numdata>

   make run-generator 10
   ```
4. Analyse search username execution time **before** the indexing applied.
   ```sql
   EXPLAIN ANALYZE SELECT * FROM comments WHERE username = 'faith3265';
   ```
5. Create an index for username column by running `make migrate-up 1`. In this 2nd migration, it will create index for username column.
6. Analyse search username execution time **after** the indexing applied.
   ```sql
   EXPLAIN ANALYZE SELECT * FROM comments WHERE username = 'faith3265';
   ```
7. Analyse search comment (content column) execution time **before** the indexing applied.
   ```sql
   EXPLAIN ANALYZE SELECT * FROM comments WHERE content LIKE '%with my%';
   ```
8. Create index for content column by running `make migrate-up 1`. In this 3rd migration, it will create index for content column.
9. Analyse search comment (content column) execution time **after** the indexing applied.
   ```sql
   EXPLAIN ANALYZE SELECT * FROM comments WHERE content LIKE '%with my%';
   ```
10. (Optional) Clean up phase.
   ```
   make migrate-down
   make clean
   ```

#### Command Recap

```
make build
make migrate-up 1
make run-generator 10

EXPLAIN ANALYZE SELECT * FROM comments WHERE username = 'faith3265';
make migrate-up 1
EXPLAIN ANALYZE SELECT * FROM comments WHERE username = 'faith3265';

EXPLAIN ANALYZE SELECT * FROM comments WHERE content LIKE '%with my%';
make migrate-up 1
EXPLAIN ANALYZE SELECT * FROM comments WHERE content LIKE '%with my%';

# optional
make migrate-down
make clean
```

## Additional features

### Create new migration

You can extend this feature with your own tables and use cases by creating a new migration using this command: `make migrate-create <migration_name>`.

```
make migrate-create create_comments
```