package main

import (
	"cig/internal/core"
	"cig/internal/pkg/action"
	"log"
)

func main() {
	brand := (new(core.Brand)).Init()
	brand.Name = "测试工作室"
	status, err := action.AddBrand(brand)
	if err != nil {
		log.Fatalf("Failed to add a brand: %v", err)
	}

	if status != 1 {
		log.Fatalf("Failed to add a brand!")
	}

	fetchedBrand, err := action.GetBrand(brand.Id)
	if err != nil {
		log.Fatalf("Failed to fetch the brand!")
	}

	log.Printf("The fetched brand is %#v", fetchedBrand)
}
