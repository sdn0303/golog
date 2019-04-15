# golog
This is my golang logger pkg.

## Setup
```bash
go get github.com/sdn0303/golog
```

## Usage
```go
package main

import "github.com/sdn0303/golog"

var golg *golog.Logger

func init() {
	golg = golog.GetInstance()
}

func main()  {
	golg.Info("Hello World")
}
```
