package LocalMapInfos

import "fmt"

func RecoErr() { //处理错误
	if d := recover(); d != nil {
		fmt.Println(d)
	}
}
