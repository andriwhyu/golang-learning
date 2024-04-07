## EXERCISE: Assign slice literals

1. Assign the following data using slice literals to the slices that you've declared in the first exercise.

   1. The names of your three best friends (to the names slice)

   2. The distances to five different locations (to the distances slice)

   3. Five bytes of data (to the data slice)

   4. Two currency exchange ratios (to the ratios slice)

   5. Up/Down status of four different web servers (to the alives slice)

2. Print their type, length and whether they're equal to nil value or not

3. Compare the length of the distances and the data slices; print a message
   if they are equal (use an if statement).

## EXPECTED OUTPUT
```
names    : []string 3 false
distances: []int 5 false
data     : []uint8 5 false
ratios   : []float64 2 false
alives   : []bool 4 false
The length of the distances and the data slices are the same.
```