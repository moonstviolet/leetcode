/*
 * Created by Reverie
 * 2020.07.28
 * 22:19:11
 */

package main

import "fmt"

//遇到左括号入栈,遇到右括号看栈顶是否匹配,最后看栈是否为空
func isValid(s string) bool {
	var st []rune
	for _, v := range s {
		switch v {
		case '(', '[', '{':
			st = append(st, v)
		case ')':
			if len(st) == 0 || st[len(st)-1] != '(' {
				return false
			}
			st = st[:len(st)-1]
		case ']':
			if len(st) == 0 || st[len(st)-1] != '[' {
				return false
			}
			st = st[:len(st)-1]
		case '}':
			if len(st) == 0 || st[len(st)-1] != '{' {
				return false
			}
			st = st[:len(st)-1]
		}
	}
	return len(st) == 0
}

func main() {
	fmt.Println(isValid("{[]}"))
}
