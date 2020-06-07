package main

import(
  "fmt"
  "./cdb2cnf"
)



func main() {

    seq:=cdb2cnf.NewSeq(",","1/e(",")","","","","","-","")

  	s := []int{1, 4, 5, 3, 7, 8, 12, 17, 13, 11, 22, 25, 26, 31, 2}

   	cdb2cnf.PrintSequence(s)
  	fmt.Println(seq.Compact(s))
    fmt.Println()

    seq =cdb2cnf.NewSeq(",","","","port1.0.","","port1.0.","-","port1.0.","")
    fmt.Println(seq.Compact(s))
    /*
    sc :=  "3,8-10,20,50-55,1-5"
    fmt.Println(sc)
    cdb2cnf.PrintSequence(cdb2cnf.DecompactSequence(sc))
    */

}
