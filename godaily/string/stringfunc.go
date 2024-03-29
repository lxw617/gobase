package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	// 统计字符串的长度,按字节len(str)
	// golang的编码统一为utf-8(ascii的字符(字母和数字)占一个字节，汉字占用3个字节)
	str := "hello北"
	fmt.Println("str len=", len(str)) // 8

	str2 := "he1lo北京"
	// 字符串遍历,同时处理有中文的问题r:=[]rune(str)
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}

	// 字符串转整数:n,err := strconv.Atoi("12")
	n, err := strconv.Atoi("he11o")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转成的结果是", n)
	}

	// 整数转字符串 str = strconv.Itoa(12345)
	str = strconv.Itoa(12345)
	fmt.Printf("str=%v,str=%T\n", str, str)

	// 字符串转[]byte: var bytes=[]byte("hello go")
	bytes := []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)

	//[]byte转字符串:str = string([]byte{97, 98,99})
	str = string([]byte{97, 98, 99})
	fmt.Printf("str=%v\n", str)

	// 10进制转2，8，16进制:str = strconv.FormatInt(123，2),返回对应的字符串
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制是=%v\n", str)
	str = strconv.FormatInt(123, 16)
	fmt.Printf("123对应的16进制是=%v\n", str)

	// 查找子串是否在指定的字符串中:strings.Contains("seafood", "foo")//true
	b := strings.Contains("seafood", "maryt")
	fmt.Printf("b=%v\n", b)

	// 统计一个字符串有几个指定的子串 :strings.Count(" ceheese", "e")//4
	num := strings.Count("ceheese", "e")
	fmt.Printf("num=%v\n", num)

	// 不区分大小写的字符串比较(=是区分字母大小写的):fmt.Println(strin
	b = strings.EqualFold("abc", "Abc")
	fmt.Printf("b=%v\n", b)           // true
	fmt.Println("结果", "abc" == "Abc") // false//区分字母大小写

	// 返回子串在字符串第一次出现的index值,如果没有返回-1://strings.Index("NLT_abc","abc")1/4
	index := strings.Index("NLT_abcabcabc", "abc") // 4
	fmt.Printf("index=%v\n", index)

	// 返回子串在字符串最后一次出现的index,
	// 如没有返回-1 : strings.LastIndex("go golang",“go")
	index = strings.LastIndex("go golang", "go") // 3
	fmt.Printf("index=%v\n", index)

	// 将指定的子串替换成另外一个子串:strings.Replace( "go go hello", "go", "go语言",n)//n可以指定你希望替换几个,如果n=-1表示全部替换
	str2 = "go go hello"
	str = strings.Replace(str2, "go", "北京", -1)
	fmt.Printf("str=%v str2=%v\n", str, str2)

	// 按照指定的某个字符,为分割标识，将一个字符串拆分成字符串数组://strings.Split("hello,wrold,ok",","）
	strArr := strings.Split("he1lo,wrold,ok", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v]=%v\n", i, strArr[i])
		fmt.Printf("strArr=%v\n", strArr)
	}

	// 15)将字符串的字母进行大小写的转换:
	// strings.ToLower("co")l/ go strings.ToUpper("Go") //Go
	str = "goLang He1lo"
	str = strings.ToLower(str)
	str = strings.ToUpper(str)
	fmt.Printf("str=%v\n", str) // golang hello

	// 将字符串左右两边的空格去掉
	str = strings.TrimSpace(" tn a lone gopher ntrn ")
	fmt.Printf("str=%q\n", str)

	// 将字符串左右两边指定的字符去掉
	// strings. Trim("! hello! ",“I")//["hello"]//将左右两边!和”“去掉
	str = strings.Trim("! hel11o! ", "!")
	fmt.Printf("str=%q\n", str)
	// 将字符串左边指定的字符去掉 ： strings.TrimLeft("! hello! ", " !") // ["hello"] //将左边! 和""去掉
	// 将字符串右边指定的字符去掉 ：strings.TrimRight("! hello! ", " !") // ["hello"] //将右边! 和""去掉

	// 判断字符串是否以指定的字符串开头: strings.HasPrefix("ftp://192.168.10.1", "ftp") // true
	b = strings.HasPrefix("ftp://192.168.10.1", "ftp")
	fmt.Printf("b=%v\n", b)
	// 判断字符串是否以指定的字符串结束: strings.HasSuffix("NLT_abc.jpg", "abc") //false

	fmt.Println(len("01510001汉字"))                    // 14
	fmt.Println(utf8.RuneCountInString("01510001汉字")) // 10

	// str字符串和字符区别
	// fmt.Println(len("6积分:"))                    //8
	// fmt.Println(utf8.RuneCountInString("6积分:")) //4
	// fmt.Printf("%-15s%v\n", "6积分:", "6个(剩:1个)")
	// fmt.Println(utf8.RuneCountInString("支付宝5折打水券:")) //9
	// fmt.Println(len("支付宝5折打水券:"))                    //23

	// fmt.Printf("%-13s%v\n", "支付宝5折打水券:", "6个(剩:1个)")
	// fmt.Printf("%-18s%v\n", "18积分:", "6个(剩:1个)")
	// fmt.Printf("%-13s%v\n", "支付宝1折打水券:", "6个(剩:1个)")
	// fmt.Printf("%-18s%v\n", "66积分: ", "6个(剩:1个)")
	// fmt.Printf("%-14s%v\n", "腾讯视频年卡:", "6个(剩:1个)")

	// fmt.Println(utf8.RuneCountInString("6积分:       1")) //12
	// fmt.Println(utf8.RuneCountInString("腾讯视频年卡:1"))     //8
	// fmt.Println(len("支付宝1折打水券"))                        //16
	// fmt.Println(len("腾讯视频年卡"))                          //18
	// fmt.Println("6积分:       1")
	// fmt.Println("腾讯视频年卡:")
}
