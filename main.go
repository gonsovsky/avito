package main

import (
	"avito/api"
	"avito/cp"
	"avito/shared"
)

func main() {
	shared.Log("Hello!")

	defer func() {
		if r := recover(); r != nil {
			shared.Log("(main) Recovered in f", r)
		}
	}()

	go api.NewApi()
	cp.NewCP()
	shared.Log("Bye!")
}
