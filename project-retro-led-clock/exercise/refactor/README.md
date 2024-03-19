## EXERCISE: Refactor

Goal is refactoring the clock project by moving some of its parts to
a new file in the same package.

1. Create a new file: placeholders.go
2. Move the placeholder type to placeholder.go
3. Move all the placeholders (zero to nine and the colon) to placeholder.go
4. Move the digits array to placeholders.go

## HINT
+ placeholders.go file should belong to main package as well

  To remember how to do so: Rewatch the "PART I — Packages, Scopes and Importing"
  section.

+ Short declaration won't work in the package scope.
  Remember why by watching: "Short Declaration: Package Scope" lecture

+ If you receive the following error and you don't know what to do watch:
  "Packages - Learn how to run multiple files" lecture.

  #command-line-arguments
  
  undefined: placeholder

  undefined: colon

## EXPECTED OUTPUT
 The same output before the refactoring.