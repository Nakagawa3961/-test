package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// 個数があってないときの表示方法
	if len(os.Args) != 3 {
		fmt.Println("引数の数が正しくありません。2つの引数を指定してください。")
		return
	}

	num1, err1 := strconv.Atoi(os.Args[1])
	num2, err2 := strconv.Atoi(os.Args[2])

	// エラーチェック?
	if err1 != nil ||
		err2 != nil {
		fmt.Println("引数の形式が正しくありません。整数を指定してください。")
		return
	}

	// 足し算を実行
	fmt.Println(num1 + num2)

	// 結果を出力
	fmt.Println()
}
