package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// purpose := []string{`{"address":{"id":12051,"name":"深圳大学","type":1},"type":0}`, "", `{"address":{"id":12051,"name":"深圳大学","type":1}}`}
	purpose := []string{`{}`}
	for _, v := range purpose {
		var manageConfigMap map[string]interface{}
		err := json.Unmarshal([]byte(v), &manageConfigMap)
		if err != nil {
			continue
		}
		manage := manageConfigMap["type"].(float64)
		// manage, ok := manageConfigMap["type"].(float64)
		// if !ok {
		// 	fmt.Println(manage)
		// 	fmt.Println(ok)
		// }
		fmt.Println(manage)
	}
}
