## EXERCISE: Fix the backing array problem

 Ensure that changing the elements of the `mine` slice
 does not change the elements of the `nums` slice.


### CURRENT OUTPUT (INCORRECT)
```
 Mine         : [-50 -100 -150 25 30 50]
 Original nums: [-50 -100 -150]
 ```


### EXPECTED OUTPUT
```
 Mine         : [-50 -100 -150]
 Original nums: [56 89 15]
```