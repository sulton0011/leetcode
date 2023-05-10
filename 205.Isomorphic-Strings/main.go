package main

import "fmt"

func main() {

	s := "egg"
	t := "adda"
	resp := isIsomorphic(s, t)

	fmt.Println("resp::", resp)
}

func isIsomorphic(s string, t string) bool {
	mapS := [256]byte{}
	mapT := [256]byte{}

	if len(s) != len(t) {
		return false
	}
	for i, v := range s {
		if mapS[v] != mapT[t[i]] {
			return false
		}
		mapS[v] = byte(i + 1)
		mapT[t[i]] = byte(i + 1)
	}
	return true
}
