package cdb2cnf

import (
	"fmt"
	"strconv"
	"strings"
)

const SEP_S = ","
const SEP_T = "-"

type Seq struct {
  Sep string
  Befs string
  Afts string
  Befu string
  Aftu string
  Befa string
  Afta string
  Befb string
  Aftb string
}

func NewSeq(sep,befs,afts,befu,aftu,befa,afta,befb,aftb string) *Seq{
  s :=Seq{sep,befs,afts,befu,aftu,befa,afta,befb,aftb}
  return &s
}


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

func (seq *Seq) Compact(s []int) string {
	s = unique(reorder(s))
  if(len(s)==1){
  return seq.Befu+strconv.Itoa(s[0])+seq.Aftu
}

	list := make([]string, 0)

	cur := s[0]
	for n := 0; n < len(s)-1; n = n + 1 {
		if s[n+1] != s[n]+1 {
			list = append(list, seq.bunk(cur, s[n]))
			cur = s[n+1]
		}
	}
	list = append(list, seq.bunk(cur, s[len(s)-1]))


	return seq.Befs+strings.Join(list, seq.Sep)+seq.Afts
}

func (seq *Seq) bunk(a, b int) string {
	if a == b {
		return seq.Befa+strconv.Itoa(a)+seq.Afta
	}
	return seq.Befa+strconv.Itoa(a) +seq.Afta+ seq.Befb + strconv.Itoa(b)+seq.Aftb
}

/*
func Decompact(ls string) []int {
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
*/
