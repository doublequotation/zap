# Zap

**A build tool for injecting env variables on build**

## Goals

-   [x] All in one build tool that can do all the copying and removing its self.
-   [ ] Define your own steps for the build process.

## Getting started

for the lines, which you want to inject the variable into. Include this above the variable declaration:

```go
// @zap var: password
var databasePassword string = "fsd"
```

When you're done commenting the injection spots run the following:

```
zap -env password=1234 -in exampleFile.go
```
