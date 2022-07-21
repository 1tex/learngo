package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is ty.zhao@live.com
email1 is abc@def.org
emial2 is    kkk@qq.com.cn
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9.]+)@([a-zA-Z0-9]+)(.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}
