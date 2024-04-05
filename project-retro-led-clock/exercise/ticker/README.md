## EXERCISE: Ticker: Slide the Clock

Your goal is slide the placeholders every second.
Please run the solution to see it in action.

## THIS IS A HARD EXERCISE:
+ It will take you days but it will be worth it.
+ For experienced developers, this can take an hour or so.

1. You need to determine the starting and the ending digits to create
    the sliding effect.

2. Each second, start from the next placeholder, skip the previous one.
   This means: Only draw the next placeholders.
```
  Like this:

  12:40:31
  2:40:31
  40:31
  0:31
  :31
  31
  1
```
3. After the last placeholder is displayed, fill the lines for the missing
   placeholders, and then start from the first placeholder. Draw it to the
   right part of the screen.
```
  Like this:

  12:40:31
  2:40:31
  40:31
  0:31
  :31
  31
  1
         1
        12
       12:
      12:4
     12:40
    12:40:
   12:40:3
  12:40:31
```
  As you can see, you need to draw the clock from the right part of the
  screen, beginning from the first placeholder.

## HINTS
 + You would need to clear the screen inside the loop instead of once.
   Otherwise, the previous placeholders will be left on the screen.


## EXPECTED OUTPUT
 Please run the solution to see it in action. Do not look at the source-code though.