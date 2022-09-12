# Goal

The goal is to create a repository with a `/examples` folder that contains a `go.mod` file. Then, another repository can do `$ go get github.com/kidfrom/learn-golang-multiple-modules` and import the `hello.go` file

# How to do it?

Use `go work` to manage multi-module.

e.g
```
// create go.sum file in the root directory.
$ go work init

// create go.mod file in the root directory / the child directory. It's depend.
$ go mod init github.com/kidfrom/learn-golang-multiple-modules
$ cd ./examples
$ go mod init github.com/kidfrom/learn-golang-multiple-modules/examples

// add module to go.work file
$ go work use .             // if there is go.mod in the root directory
$ go work use ./examples    // because there is a go.mod in the ./examples directory
```

Then, push it to the cloud `$ git push`.

After that, use this module on different repository, e.g `github.com/kidfrom/test`.

e.g
```
$ go get github.com/kidfrom/learn-golang-multiple-modules

// if you want to get the ./examples directory too.
$ go get github.com/kidfrom/learn-golang-multiple-modules/examples
```

# Update

go work is not to be used for `/examples` folder. This is because `/examples` folder dependent to the `hello` and `world` packages.

In other words, you need to do `$ go get github.com/kidfrom/learn-golang-multiple-modules@sha` in the `/examples` folder everytime you make changes to the `hello` or `world` packages.

For example, if you make changes to the `hello/hello.go` and test it with `/examples/main.go`, let's say it works. Then, you make a docker image with the `/examples` folder. In the Dockerfile, you will need to do `RUN go mod download`, it will not download the latest commit, which may result to error.