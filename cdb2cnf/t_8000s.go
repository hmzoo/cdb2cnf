package cdb2cnf

import(
  "strconv"
)



type T_8000S struct {

}

func (b T_8000S) GetSep() string {
  return ","
}


func (b T_8000S) BeforeInt(u bool) string {
  if(u) {
    return "interface ethernet"
  }
  return "interface range ethernet"
}

func (b T_8000S) Seq(stkid int,list []int) string {
  stackid := strconv.Itoa(stkid)
  seq:= NewSeq(",",stackid+"/e(",")",stackid+"/e","","","","-","")
  return seq.Compact(list)
}

func (b T_8000S) ContentInt(vi int) string {
  vlan := strconv.Itoa(vi)
  return "switchport access vlan "+vlan

}

func (b T_8000S) AfterInt() string {
  return "exit"
}
