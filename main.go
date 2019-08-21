package main

import (
	"fmt"
	j "github.com/TheStarBoys/implement-data-structures-and-algorithms-with-golang/algorithms/tag/jsonHandle"
)

func main() {
	json := `{"code":0,"message":"","data":[{"account_id":"1008958616510007","balance":580500,"available_balance":580500,"frozen_balance":0,"unconfirmed_balance":0,"currency":"BTC"}]}`
	fmt.Println(j.CompareJson([]byte(json), []byte(json)))
}

