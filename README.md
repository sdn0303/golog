# golog
This is my golang logger pkg.

## Setup
```bash
go get github/sdn0303/golog
```

## Usage
```go
package main

import "github/sdn0303/golog"

var gl *golog.Logger

func init() {
	gl = golog.GetInstance()
}

func main()  {
	gl.Logging.Info("Hello World")
}
```