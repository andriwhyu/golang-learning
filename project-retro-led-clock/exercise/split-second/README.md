## EXERCISE: Add Split Seconds

Your goal is adding the split second to the clock. A split second is
1/10th of a second.

1. Find the current split second
2. Add dot character to the clock (as in the expected output)
3. Add the split second digit to the clock
4. Blink the dot every two seconds (just like the separators)
5. Update the clock every 1/10th of a second, instead of every second.
     (Update the clock every 100 millliseconds)

## HINTS
  + You can find the split second using Nanosecond method of the Time type.
    https://golang.org/pkg/time/#Time.Nanosecond

  + A split second is the first digit of the Nanosecond.

  + Remember: time.Second is an integer constant, so it can be divided
              with a number.
    https://golang.org/pkg/time/#Time.Second

## EXPECTED OUTPUT
Note that, clock is updated every split second instead of a second.

Separators are displayed (second is an odd number):
```
  ██   ██        ███  ██        ██   ███       ███
   █    █    ░   █     █    ░    █     █       █ █
   █    █        ███   █         █     █       █ █
   █    █    ░     █   █    ░    █     █       █ █
  ███  ███       ███  ███       ███    █   ░   ███

  ██   ██        ███  ██        ██   ███       ██
   █    █    ░   █     █    ░    █     █        █
   █    █        ███   █         █     █        █
   █    █    ░     █   █    ░    █     █        █
  ███  ███       ███  ███       ███    █   ░   ███

  ██   ██        ███  ██        ██   ███       ███
   █    █    ░   █     █    ░    █     █         █
   █    █        ███   █         █     █       ███
   █    █    ░     █   █    ░    █     █       █
  ███  ███       ███  ███       ███    █   ░   ███

  ██   ██        ███  ██        ██   ███       ███
   █    █    ░   █     █    ░    █     █         █
   █    █        ███   █         █     █       ███
   █    █    ░     █   █    ░    █     █         █
  ███  ███       ███  ███       ███    █   ░   ███

  ██   ██        ███  ██        ██   ███       █ █
   █    █    ░   █     █    ░    █     █       █ █
   █    █        ███   █         █     █       ███
   █    █    ░     █   █    ░    █     █         █
  ███  ███       ███  ███       ███    █   ░     █

  ██   ██        ███  ██        ██   ███       ███
   █    █    ░   █     █    ░    █     █       █
   █    █        ███   █         █     █       ███
   █    █    ░     █   █    ░    █     █         █
  ███  ███       ███  ███       ███    █   ░   ███

  ██   ██        ███  ██        ██   ███       ███
   █    █    ░   █     █    ░    █     █       █
   █    █        ███   █         █     █       ███
   █    █    ░     █   █    ░    █     █       █ █
  ███  ███       ███  ███       ███    █   ░   ███

  ██   ██        ███  ██        ██   ███       ███
   █    █    ░   █     █    ░    █     █         █
   █    █        ███   █         █     █         █
   █    █    ░     █   █    ░    █     █         █
  ███  ███       ███  ███       ███    █   ░     █

  ██   ██        ███  ██        ██   ███       ███
   █    █    ░   █     █    ░    █     █       █ █
   █    █        ███   █         █     █       ███
   █    █    ░     █   █    ░    █     █       █ █
  ███  ███       ███  ███       ███    █   ░   ███

  ██   ██        ███  ██        ██   ███       ███
   █    █    ░   █     █    ░    █     █       █ █
   █    █        ███   █         █     █       ███
   █    █    ░     █   █    ░    █     █         █
  ███  ███       ███  ███       ███    █   ░   ███

  Separators are not displayed (second is an even number):

  ██   ██        ███  ██        ██   ███       ███
   █    █        █     █         █   █ █       █ █
   █    █        ███   █         █   ███       █ █
   █    █          █   █         █   █ █       █ █
  ███  ███       ███  ███       ███  ███       ███
```