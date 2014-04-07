package debug

import (
	"fmt"
	"os"
)

var debugFile = os.Getenv("FILE_DEBUG")

func init() {
	Debug("init")
}

func Debug(format string, args ...interface{}) {
	f, err := os.OpenFile(debugFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
		return
	}
	defer f.Close()
	fmt.Fprintf(f, format+"\n", args...)
}
