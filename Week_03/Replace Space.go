package main

import "fmt"

func main() {
	s:="we are good"
	fmt.Println(replaceSpace(s))
}
func replaceSpace(s string) string {
	var str string
	for _,v:=range s{
		if string(v)==" "{
			str+="%20"
		}else {
			str+=string(v)
		}
	}
	return str
}
