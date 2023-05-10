package main

import "fmt"

func main() {
	s := "abc"
	t := "ahbgds"
	resp := isSubsequence(s, t)
	fmt.Println("Response: ", resp)
}

func isSubsequence(s string, t string) bool {

	k := 0
	for i := 0; i < len(t); i++ {
		if k < len(s) && s[k] == t[i] {
			k++
		}

	}
	return len(s) == k
}
