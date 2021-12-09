package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv" //intへのキャスト、Atoi関数用
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords) //ここでsplitして各要素を空白区切りで認識可能にする

	for sc.Scan() { // sc.Scanの戻り値はboolでforの条件に入れることによって終端時に抜けるようにする。
		// sc.Scan時にセットされたテキストをsc.Textで得る。sc.Textをキャストし、vに代入。
		v, _ := strconv.Atoi(sc.Text())
		fmt.Println(v)
	}
}
