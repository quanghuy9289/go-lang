# Go programming tutorial

## 1. Go workspace

In Go, programs are kept in a directory hierarchy that is called a workspace.
A workspace is simply a root directory of your Go applications.
A workspace contains three subdirectories at its root:
  - `src` – This directory contains source files organized as packages. You will write your Go applications inside this directory.
  - `pkg` - This directory contains Go package objects.
  - `bin` - This directory contains executable programs.

You have to specify the location of the workspace before developing Go programs => set `GOPATH` environment to working directory of your program

## 2. Packages

In Go, source files are organized into system directories called packages, which enable code reusability across the Go applications.

In Go packages, all identifiers will be exported to other packages if the first letter of the identifier name starts with an uppercase letter. The functions and types will not be exported to other packages if we start with a lowercase letter for the identifier name.

### Package type

- `Go's standard library` - build-in packages when install Go: `net/http, fmt, math...`, are available at the `src` subdirectory of the `GOROOT directory` - `/usr/local/go/src` in Linux base system

- `Third-party packages` - packages that are created by yourself or GO community, are available in the `pkg` subdirectory of the `GOPATH location` - `/Users/username/go/pkg` in Linux base system

### Import packages

```go
    package main
    
    import (
      "fmt"
    )
    func main(){
    fmt.Println("Hello, Gopher!")
    }
```

When you import packages, the Go compiler will look on the locations specified by the environment variable `GOROOT` and `GOPATH`.

### Install 3rd packages

Using `go get` command. The Go get command will fetch the packages from the source repository and put the packages in the `pkg` subdirectory of the `GOPATH location`.

> go get github.com/99designs/gqlgen

### Alias names for packages

We can use alias names for packages to avoid package name ambiguity.

```go
package main
import (
        mongo "mywebapp/libs/mongodb/db"
        mysql "mywebapp/libs/mysql/db"
)
func main() {
    mongodata :=mongo.Get() //calling method of package  "mywebapp/libs/mongodb/db"
    sqldata:=mysql.Get() //calling method of package "mywebapp/libs/mysql/db"  
    fmt.Println(mongodata )
    fmt.Println(sqldata )
}
```

We can note that the name of both packages are the same - `db`. We can use an alias name for one package, and can use the alias name whenever we need to call a method of that package.

## 3. Init function

When you write Go packages, you can provide a `optional function` `init` that will be called at the beginning of the execution time.
The init method can be used for adding initialization logic into the package.

```go
  package db
  import (
        "gopkg.in/mgo.v2"
          "gopkg.in/mgo.v2/bson"
  )
  func init {
    // initialization code here    
  }
```

In some contexts, we may need to import a package only for invoking it’s init method, where we don’t need to call forth other methods of the package. If we imported a package and are not using the package identifier in the program, Go compiler will show an error. In such a situation, we can use a blank identifier ( _ ) as the package alias name, so the compiler ignores the error of not using the package identifier, but will still invoke the init function.

```go
package main
import (
  _ "mywebapp/libs/mongodb/db" // not error
	"fmt"
	"log"
)
func main() {
  //implementation here
}
```

## 4. Creating and Reusing packages

- Set `GOPATH` environment to your working space - define location for downloading 3rd packages

> export GOPATH=/path/desired/here

- Create your package - crete folder `lib` and add file `languages.go`

```go
package lib
var languages map[string]string
func init(){
  languages= make(map[string]string)
  languages["cs"] = "C #"
  languages["js"] = "JavaScript"
  languages["rb"] = "Ruby"
  languages["go"] = "Golang"
}
func Get(key string) (string){
  return languages[key]
}
func Add(key,value string){
  languages[key]=value
}
func GetAll() (map[string]string){
  return languages
}
```

- Reuse package in your program - create `test` program folder and add `main.go` file

```go
package main
 
import (
  "fmt"
  "your_path/lib"
)
 
func main() {
  lib.Add("dr","Dart")
  fmt.Println(lib.Get("dr"))
  languages:=lib.GetAll()
  for _, v := range languages {
    fmt.Println(v)
  }
}
```

- Run `go install` - build the package `lib` which will be available at the `pkg` subdirectory of `GOPATH`, and create executable file in `$GOPATH/bin` and that can be called by typing program name (`test`) in terminal.

## 5. Create your program with Go module

With Go module - from Go version `1.11` we don't need to set `GOPATH` to your working directory - we are no longer limited to `GOPATH`

Just open your IDE at your project folder - contains `go.mod` file

Now we can import your package by `module_name/package_path`

Refer: `https://github.com/quanghuy9289/go-lang/tree/master/ch11`
