package main

import(
  "fmt"
  "./cdb2cnf"
)



func main() {


  	s := []int{1, 4, 5, 3, 7, 8, 12, 17, 13, 11, 22, 25, 26, 31, 2}
    sc :=  "3,8-10,20,50-55,1-5"
   	cdb2cnf.PrintSequence(s)
  	fmt.Println(cdb2cnf.CompactSequence(s))
    fmt.Println()
    fmt.Println(sc)
    cdb2cnf.PrintSequence(cdb2cnf.DecompactSequence(sc))

}
