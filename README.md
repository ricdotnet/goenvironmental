
# Go Environment Variables

Small package to help with the use of Environment Variables in Go projects.


## Usage/Examples

```cmd
go get github.com/ricdotnet/goenvironmental@master
```

The main package name is long, so I recommend using an import variable to help manage better.
```go
import (
goenv "github.com/ricdotnet/goenvironmental"
)
```

```go
func main() {
// this function starts by parsing all environment variables to a map
goenv.ParseEnv()

// to use a variable simply call its name
var1 := goenv.EnvVariables["NAME"]

// goenv.EnvVariables is also a long name so we can assign it to a smaller variable
envs := goenv.EnvVariables

// then we can use
envs["NAME"]
}
```
