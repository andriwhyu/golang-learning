## EXERCISE: Print daily requests

You've got request logs of a web server. The log data contains 8-hourly totals per each day. It is stored
in the `reqs` slice.

Find and print the total requests per day, as well as the grand total.

See the `reqs` slice and the steps in the code below.


### RESTRICTIONS

   1. You need to produce the daily slice, don't just loop
      and print the element totals directly. The goal is
      gaining more experience in slice operations.

   2. Your code should work even if you add to or remove the
      existing elements from the `reqs` slice.

      For example, after solving the exercise, try it with
      this new data:
   ```
      reqs := []int{
 	      500, 600, 250,
 	      200, 400, 50,
 	      900, 800, 600,
 	      750, 250, 100,
 	      150, 654, 235,
 	      320, 534, 765,
 	      121, 876, 285,
 	      543, 642,
 	      // the last element is missing (your code should be able to handle this)
 	      // that is why you shouldn't calculate the `size` below manually.
      }
   ```
The grand total of the new data should be **10525**.


### EXPECTED OUTPUT

   Please run `solution/main.go` to see the expected
   output.

   `go run solution/main.go`