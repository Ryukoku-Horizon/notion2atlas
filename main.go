package main

import (
	"notion2atlas/presentation"

	"github.com/joho/godotenv"
)

func main() {
	var err error = nil
	godotenv.Load()
	err = presentation.HandleUpdateData()
	// aaa, err := usecase.Test("2aba501ef33780eaa8f6ca103f2d1d09")
	// filemanager.WriteJson(aaa, "notion_data/test.json")
	if err != nil {
		panic(err)
	}
}

// https://www.notion.so/1ada501ef33781e8b251df0f38a0cb23?v=1ada501ef33781c291ee000c04c13f15&source=copy_link#2aba501ef33780eaa8f6ca103f2d1d09
