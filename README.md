[![Go Report Card](https://goreportcard.com/badge/github.com/ricdotnet/goenvironmental)](https://goreportcard.com/report/github.com/ricdotnet/goenvironmental)

# Go Environment Variables

Small package to help with the use of Environment Variables in Go projects.


## Usage/Examples

```cmd
go get github.com/ricdotnet/goenvironmental@latest
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
  // we can also pass a custom file (ie. dev mode envs)
  goenv.ParseEnv(".env.development")

  // to use a variable simply call its name
  var1 := goenv.EnvVariables["NAME"]

  // goenv.EnvVariables is also a long name so we can assign it to a smaller variable
  envs := goenv.EnvVariables

  // then we can use
  envs["NAME"]
}
```
