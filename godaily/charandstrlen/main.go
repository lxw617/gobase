package main

import (
	"fmt"
)

func main() {
	// str字符串和字符区别
	// fmt.Println(len("6积分:"))                    //8
	// fmt.Println(utf8.RuneCountInString("6积分:")) //4
	// fmt.Printf("%-15s%v\n", "6积分:", "6个(剩:1个)")
	// fmt.Println(utf8.RuneCountInString("支付宝5折打水券:")) //9
	// fmt.Println(len("支付宝5折打水券:"))                    //23
	fmt.Printf("%-13s%v\n", "支付宝5折打水券:", "6个(剩:1个)")
	fmt.Printf("%-18s%v\n", "18积分:", "6个(剩:1个)")
	fmt.Printf("%-13s%v\n", "支付宝1折打水券:", "6个(剩:1个)")
	fmt.Printf("%-18s%v\n", "66积分: ", "6个(剩:1个)")
	fmt.Printf("%-14s%v\n", "腾讯视频年卡:", "6个(剩:1个)")

	// fmt.Println(utf8.RuneCountInString("6积分:       1")) //12
	// fmt.Println(utf8.RuneCountInString("腾讯视频年卡:1"))     //8
	// fmt.Println(len("支付宝1折打水券"))                        //16
	// fmt.Println(len("腾讯视频年卡"))                          //18
	// fmt.Println("6积分:       1")
	// fmt.Println("腾讯视频年卡:")
}
