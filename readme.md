# logger

## use
```
go get github.com/jummyliu/logger
```

## usage
```go
package main

import (
	"io"
	"os"

	"github.com/jummyliu/logger"
	"github.com/jummyliu/logger/defaultlogger"
)

func main() {
	name := "tmp.log"
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	writerList := []io.Writer{file}
	// log to console
	// writerList = append(writerList, os.Stdout)

	Logger := defaultlogger.New(io.MultiWriter(writerList...), logger.LevelDebug)

	Logger.LogDebug("debug")
	Logger.LogInfo("Info")
}

```