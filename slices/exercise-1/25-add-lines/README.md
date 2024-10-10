## EXERCISE: Add a newline after each sentence

  You have a slice that contains Beatles' awesome song:
  Yesterday. You want to add newlines after each sentence.

  So, create a new slice and copy every words into it. Lastly,
  after each sentence, add a newline character ('\n').


### ORIGINAL SLICE:
```
[yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday]
```

### EXPECTED SLICE (NEW):
```
[yesterday all my troubles seemed so far away \n now it looks as though they are here to stay \n oh I believe in yesterday \n]
```

### CURRENT OUTPUT
```
yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday
```

### EXPECTED OUTPUT
```
  yesterday all my troubles seemed so far away
  now it looks as though they are here to stay
  oh I believe in yesterday
```


### RESTRICTIONS

  + Don't use `append()`, use `copy()` instead.

  + Don't cheat like this:
    ```
    fmt.Println(lyric[:8])
    fmt.Println(lyric[8:18])
    fmt.Println(lyric[18:23])
    ```

  + Create a new slice that contains the sentences
    with line endings.


### NOTE

If the program does not work, please update your local copy of the prettyslice package:

   `go get -u github.com/inancgumus/prettyslice`