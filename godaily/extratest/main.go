package main

import (
	"encoding/json"
	"fmt"
)

type Extra struct {
	IsSuccess   int    `json:"isSuccess"`
	Description string `json:"description"`
	BatchNo     string `json:"batchNo"`
}

func main() {
	extraStr := `{"batchNo":"221205154649036514","isSuccess":0}`
	extra := Extra{}
	err := json.Unmarshal([]byte(extraStr), &extra)
	fmt.Println(err)
	fmt.Println(extra)
}
