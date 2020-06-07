package cdb2cnf

import(
  "fmt"
)



type Port struct{
 Lt string
 Comm string
 Num int
 Vlan string
 Vlanid int
 Stack string
 Stackid int
}

type Ports []*Port

func NewPorts() Ports{
  return Ports{}
}

func (p Port) Print(){
  fmt.Printf("%10s %30s %d %10s %d %10s %d\n", p.Lt, p.Comm, p.Num, p.Vlan, p.Vlanid, p.Stack, p.Stackid)
}

func (data Ports) ReOrder() {
data.ReOrderNum()
data.ReOrderStack()
data.ReOrderVlan()

}

func (data Ports) ReOrderNum() {

	for k, _ := range data {
		if k < len(data)-1 {
			c := data[k]
			cn := data[k+1]
			if cn.Num < c.Num {
				data[k] = cn
				data[k+1] = c
				data.ReOrderNum()
			}
		}
	}
}

func (data Ports) ReOrderStack() {
  for k, _ := range data {
    if k < len(data)-1 {
      c := data[k]
      cn := data[k+1]
      if cn.Stackid < c.Stackid {
        data[k] = cn
        data[k+1] = c
        data.ReOrderStack()
      }
    }
  }
}

func (data Ports) ReOrderVlan() {
  for k, _ := range data {
    if k < len(data)-1 {
      c := data[k]
      cn := data[k+1]
      if cn.Vlanid < c.Vlanid {
        data[k] = cn
        data[k+1] = c
        data.ReOrderVlan()
      }
    }
  }
}
