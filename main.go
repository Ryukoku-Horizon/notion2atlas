package main

import (
	"notion2atlas/presentation"

	"github.com/joho/godotenv"
)

func main() {
	var err error = nil
	godotenv.Load()
	err = presentation.HandleUpdateData()
	// aaa, err := usecase.Test("2d0a501ef3378050b050f0389a743d37")
	// filemanager.WriteJson(aaa, "notion_data/test.json")
	if err != nil {
		panic(err)
	}
}
