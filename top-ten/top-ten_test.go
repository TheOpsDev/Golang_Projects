package main

import (
	"log"
	"testing"
)

func TestRankByUrlValue(t *testing.T) {
	sampleMap := map[string]int{
		"http://api.tech.com/item/11986": 2,
		"http://api.tech.com/item/12919": 5,
		"http://api.tech.com/item/12940": 17,
		"http://api.tech.com/item/13482": 14,
		"http://api.tech.com/item/13780": 1,
		"http://api.tech.com/item/1528":  11,
	}

	sortedSample := rankByUrlValue(sampleMap)
	if sortedSample[0].Url != "http://api.tech.com/item/12940" {
		log.Fatal("Sample map was not sorted correctly")
	}

	if sortedSample[5].Url != "http://api.tech.com/item/13780" {
		log.Fatal("Sample map was not sorted correctly")
	}
}
