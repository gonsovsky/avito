package main

import (
	"avito/api"
	"avito/cp"
)

func main() {
	go api.NewApi()
	cp.NewCP()
}
