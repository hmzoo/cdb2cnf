package cdb2cnf

import(
  "strconv"
)



type T_FS980 struct {

}

func (b T_FS980) GetSep() string {
  return ","
}


func (b T_FS980) BeforeInt(u bool) string {
  if(u) {
    return "interface"
  }
  return "interface"
}

func (b T_FS980) Seq(stkid int,list []int) string {
  stackid := strconv.Itoa(stkid)
  seq:= NewSeq(",","","","port"+stackid+".0.","","port"+stackid+".0.","","-"+stackid+".0.","")
  return seq.Compact(list)
}

func (b T_FS980) ContentInt(vi int) string {
  vlan := strconv.Itoa(vi)
  return "switchport"+NL+"switchport mode access"+NL+"switchport access vlan "+vlan

}

func (b T_FS980) AfterInt() string {
  return "!"
}
