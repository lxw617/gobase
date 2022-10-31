package main

import (
	"fmt"
)

func main() {
	terminalParams := make([]map[string]interface{}, 0)
	params := make([]map[string]interface{}, 0)
	terminalParams = append(terminalParams, map[string]interface{}{
		"merchantId":   "merchantId",
		"terminalId":   "terminalId",
		"deviceSerial": "deviceSerial",
	})
	terminalParams = append(terminalParams, map[string]interface{}{
		"merchantId":   "111111111",
		"terminalId":   "2222222",
		"deviceSerial": "33333",
	})
	terminalParams = append(terminalParams, map[string]interface{}{
		"merchantId":   "44444444",
		"terminalId":   "55555555",
		"deviceSerial": "6666666",
	})
	fmt.Println(terminalParams)
	for _, terminal := range terminalParams {
		merchantId, _ := terminal["merchantId"]
		terminalId, _ := terminal["terminalId"]
		deviceSerial, _ := terminal["deviceSerial"]
		if merchantId == "44444444" && terminalId == "55555555" && deviceSerial == "6666666" {
			fmt.Println(merchantId, terminalId, deviceSerial)
		}
	}
	params = append(params, map[string]interface{}{
		"merchantId":   "merchantId",
		"terminalId":   "terminalId",
		"deviceSerial": "deviceSerial",
	})

}
func (nums []map[string]interface{})  string {
	if len(nums) == 0 {
        return 0
    }
    left,right := 1,1
    for right < len(nums){
        if nums[right] != nums[right-1]{
            nums[left] = nums[right]
            left ++
        }
        right ++
    }
    return left
	
}