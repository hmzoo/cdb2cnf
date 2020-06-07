package cdb2cnf

import (
	"fmt"
  "sort"
)

type VlanInt struct {
	VlanId int
  Data map[int][]int
}

type Conf struct {
	name     string
	VlanInts []VlanInt
}

func NewConf(name string) *Conf {
	ncnf := Conf{name, make([]VlanInt, 0)}
	return &ncnf
}

func NewVlanInt(vid int) VlanInt {
	nvint := VlanInt{vid, make(map[int][]int)}
	return nvint
}

func (cf *Conf) AddPorts(stk, num, vid int) {
	for n, v := range cf.VlanInts {
		if v.VlanId == vid {
			cf.VlanInts[n].Data[stk] = append(cf.VlanInts[n].Data[stk], num)
      cf.VlanInts[n].reOrder()
			return
		}
	}
	cf.VlanInts = append(cf.VlanInts, NewVlanInt(vid))
  cf.reOrder()
  cf.AddPorts(stk,num,vid)
}

func (cf *Conf) reOrder() {

  for k, _ := range cf.VlanInts {
		if k < len(cf.VlanInts)-1 {
			c := cf.VlanInts[k]
			cn := cf.VlanInts[k+1]
			if cn.VlanId < c.VlanId {
				cf.VlanInts[k] = cn
				cf.VlanInts[k+1] = c
				cf.reOrder()
			}
		}
	}

}

func (vint *VlanInt) reOrder() {
  for k := range vint.Data {
      sort.Ints(vint.Data[k])
  }
}

func (vint *VlanInt) Count() int {
  r:=0
  for k := range vint.Data {
      r=r+len(vint.Data[k])
  }
  return r
}

func (cf *Conf) Print() {
	fmt.Printf("%30s :\n", cf.name)
	for _, v := range cf.VlanInts {
		fmt.Printf("VLANID  %6d:", v.VlanId)
    keys:=[]int{}
    for k :=range v.Data {
      keys=append(keys,k)
    }
    sort.Ints(keys)
		for _,k := range keys {
			fmt.Printf(" %2d: %v", k,v.Data[k])
		}
		fmt.Println()
	}
}
