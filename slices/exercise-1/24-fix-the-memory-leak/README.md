## EXERCISE: Fix the memory leak

### GOAL
In this exercise, your goal is to reduce the memory usage. To do that, you need to find and fix the memory leak within `main()`.


### PROBLEM
`main()` calls `api.Report()` that reports the current memory usage. 
After that, `main()` calls `api.Read()` that returns a slice with 10 million elements. But you only need the last 10 elements of the returned slice.


### WHAT YOU NEED TO DO
You only need to change the code in `main()`. Please do not touch the code in `api/api.go`.

### CURRENT OUTPUT
```
    > Memory Usage: 113 KB

    Last 10 elements: [...]

    > Memory Usage: 65651 KB
```
+ Before `api.Read()` call: It uses 113 KB of memory.

+ After `api.Read()` call : It uses  65 MB of memory.

+ This means that, `main()` never releases the memory. This is the leak.

+ Your goal is to release the unused memory. Remember, you only need 10 elements but in the current code below you have a slice with 10 million of elements.


### EXPECTED OUTPUT
```
    > Memory Usage: 116 KB

    Last 10 elements: [...]

    > Memory Usage: 118 KB
```
+ In the expected output, `main()` releases the memory.

It no longer uses 65 MB of memory. Instead, it only uses 118 KB of memory. That's why the second `api.Report()` call reports 118 KB.


### ADDITIONAL NOTE

Memory leak means: Your program is using unnecessary computer memory. It doesn't release memory that is no longer needed.

See this for more information: https://en.wikipedia.org/wiki/Memory_leak


### HINTS
Check out `hints.md` file if you get stuck.