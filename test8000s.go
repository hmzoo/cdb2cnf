package main

import(
  "fmt"
  "./cdb2cnf"
)


func main(){
  cnf:=cdb2cnf.NewConf("TEST-N2")

  cnf.AddPorts(2,4,50)
  cnf.AddPorts(1,3,50)
  cnf.AddPorts(1,3,60)
  cnf.AddPorts(2,8,60)
  cnf.AddPorts(1,12,50)
  cnf.AddPorts(1,13,50)
  cnf.AddPorts(2,23,50)
  cnf.AddPorts(2,22,50)
  cnf.AddPorts(2,20,20)
  cnf.AddPorts(1,22,20)
  cnf.AddPorts(1,22,88)

  builder:= cdb2cnf.T_8000S{}

  r := cdb2cnf.Build(cnf,builder)

  fmt.Println(r)

}
