# Zap

**A build tool for injecting env variables on build**

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

This will write to the file and create a new file with the extension .copy. In this case we would have a file named `exampleFile.go` and `exampleFile.go.copy`. Then you can build your go program with the new file that has the environment variables injected. When are done building remove the file injected into and rename the `.copy` file to the original name. Now you have a binary with the injected variables and your source code looks like it never changed.
