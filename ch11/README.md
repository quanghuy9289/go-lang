# Create and use module in Go application

https://golang.org/doc/tutorial/create-module

## 1. Create local module

- Create folder `math`
- Init for math module:
  
  > go mod init `module_path`

module_path: `golang-book/ch11/math`

- Define functions in `math module`

Note that function name start with a capital letter that mean this function is published (namning convention) - other packages and programs are able to see it - if not our main program would not have been able to see it

## 2. Using local module in program

- Enable dependency tracking for the code you're about to write. (create your code as module)

> go mod init `golang-book/ch11`

- Edit `golang-book/ch11/math` module to use your local `golang-book/ch11/math` module. By default Go tools could find module from its repository (its published location). For now, because you haven't published the module yet, you need to adapt the `golang-book/ch11/math` module so it can find the `golang-book/ch11/math` code on your local file system
  
> go mod edit -replace `golang-book/ch11/math=./math`

With `./math` is the path to `golang-book/ch11/math` module on local machine

- Synchronize the `golang-book/ch11/math` module's dependencies

> go mod tity

- Write code to call functions of `golang-book/ch11/math` module

```go
package main

import (
  "fmt"
  "golang-book/ch11/math"
)

func main() {
  xs := []float64{1, 2, 3, 4, 5}
  avg := math.Average(xs)
  fmt.Println(avg)
}

```

### 3. Using your packages that not export to module

- Define your module name in `go.mod` file

```go
module github.com/quanghuy9289/gqlgen-samples/todos

go 1.17
```

- Refer to package in your code in module project by `module name/package_path`

```go
package graph

import "github.com/quanghuy9289/gqlgen-samples/todos/graph/model"

type Resolver struct {
  todos []*model.Todo
}
```

In this: `github.com/quanghuy9289/gqlgen-samples/todos` - module name and `/graph/model` - package path
Model is the package that defined in `model.go` of `todos/graph/model` folder

```go
package model
type Todo struct {
  ID   string `json:"id"`
  Text string `json:"text"`
  Done bool   `json:"done"`
  User *User  `json:"user"`
}
```
