package cdb2cnf

import (
	"fmt"
	"strconv"
	"strings"
)

const SEP_S = ","
const SEP_T = "-"


func PrintSequence(s []int) {
	for _, v := range s {
		fmt.Printf("%3d ", v)
	}
	fmt.Println()
}

func reorder(s []int) []int {

	test := false
	for k, v := range s {
		if k+1 < len(s) {
			if s[k+1] < v {
				s[k] = s[k+1]
				s[k+1] = v
				test = true
			}
		}
	}
	if test {
		return reorder(s)
	}
	return s
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func CompactSequence(s []int) string {
	s = unique(reorder(s))

	list := make([]string, 0)

	cur := s[0]
	for n := 0; n < len(s)-1; n = n + 1 {
		if s[n+1] != s[n]+1 {
			list = append(list, bunk(cur, s[n]))
			cur = s[n+1]
		}
	}
	list = append(list, bunk(cur, s[len(s)-1]))
	return strings.Join(list, SEP_S)
}

func bunk(a, b int) string {
	if a == b {
		return strconv.Itoa(a)
	}
	return strconv.Itoa(a) + SEP_T + strconv.Itoa(b)
}

func DecompactSequence(ls string) []int {
  list := make([]int, 0)

	slist := strings.Split(ls,SEP_S)
  for _,lt := range slist {
    l:=strings.Split(lt,SEP_T)
    if(len(l)==1){
      u,err := strconv.Atoi(l[0])
      if(err!=nil){
        panic(err)
      }
      list=append(list,u)
    }
    if(len(l)==2){
      a,err := strconv.Atoi(l[0])
      if(err!=nil){
        panic(err)
      }
      b,err := strconv.Atoi(l[1])
      if(err!=nil){
        panic(err)
      }
      for n:=a;n<=b;n=n+1 {
        list=append(list,n)
      }
    }
  }

  return unique(reorder(list))

}
