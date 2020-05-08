package main

import (
	"benchmarkNumberGenerator/service"
)

func main() {
	numberOfLoopTest := 1000000
	service.Run(numberOfLoopTest)
}
