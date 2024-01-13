## EXERCISE: Declare empty arrays
1. Declare and print the following arrays with their types:
   a. The names of your three best friends (names array)
   b. The distances to five different locations (distances array)
   c. A data buffer with five bytes of capacity (data array)
   d. Currency exchange ratios only for a single currency (ratios array)
   e. Up/Down status of four different web servers (alives array)
   f. A byte array that doesn't occupy memory space (zero array)

2. Print only the types of the same arrays.
3. Print only the elements of the same arrays.

## HINT
When printing the elements of an array, you can use the usual Printf verbs.
Example: When printing a string array, you can use "%q" verb as usual.

## EXPECTED OUTPUT
```
names    : [3]string{"", "", ""}
distances: [5]int{0, 0, 0, 0, 0}
data     : [5]uint8{0x0, 0x0, 0x0, 0x0, 0x0}
ratios   : [1]float64{0}
alives   : [4]bool{false, false, false, false}
zero     : [0]uint8{}

names    : [3]string
distances: [5]int
data     : [5]uint8
ratios   : [1]float64
alives   : [4]bool
zero     : [0]uint8

names    : ["" "" ""]
distances: [0 0 0 0 0]
data     : [0 0 0 0 0]
ratios   : [0.00]
alives   : [false false false false]
zero     : []
```