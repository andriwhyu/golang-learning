## EXERCISE: Adjust the width and height automatically

   Instead of setting the width and height manually,
   you need to get the width and height of the terminal
   screen from your operating system.

 1. Update your program to use my screen package.
    It offers an easy way to get the width and height.

    `go get -u https://github.com/inancgumus/screen`


 2. Read the package's documentation and find a way to
    get the screen size: width and height.

    The documentation is here:
    https://godoc.org/github.com/inancgumus/screen


 3. Use it to set the board's dimensions.


## OPTIONAL EXERCISE

 1. When you set the width, you may see that the ball
    goes beyond the left and right borders. This happens
    because the ball emoji spans to multiple console
    columns (or cells). Ordinary characters have a
    single column.

    1. Get the width of the ball emoji using a function
       from the following package:

       `go get -u github.com/mattn/go-runewidth`

    2. Divide the width using the rune width of the
       ball emoji.

 2. Your terminal may have borders, so reduce the
    height by taking into account the height of
    your terminal borders.


## EXPECTED OUTPUT

 When you run the program, the ball should start
 animating across the total width and height of your
 terminal screen dynamically.

 Currently, you set width and height manually, so it
 wasn't matter whether your terminal was bigger or
 smaller, but now it will be!