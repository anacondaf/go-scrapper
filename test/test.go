package test

import (
	"fmt"
	"runtime"
)

func RunTest() {
	_, filePath, _, _ := runtime.Caller(0)

	fmt.Println(filePath)
}
