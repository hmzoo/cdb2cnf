package cdb2cnf

import (

  "strings"
  "sort"
)

const NL ="\n"

type ConfBuilder interface {
  GetSep() string
  BeforeInt(u bool) string
  Seq(stkid int,list []int) string
  ContentInt(vi int) string
  AfterInt() string
}


func Build(cf *Conf,cfb ConfBuilder) string {
   result:=""
	for _, v := range cf.VlanInts {
    if(v.Count()==1){
      result=result+cfb.BeforeInt(true)+" "
    }else{
      result=result+cfb.BeforeInt(false)+" "
    }

    keys:=[]int{}
    for k :=range v.Data {
      keys=append(keys,k)
    }
    sort.Ints(keys)

    l :=[]string{}
		for _,k := range keys {
			l=append(l,cfb.Seq(k,v.Data[k]))
		}
		result=result+strings.Join(l,cfb.GetSep())+NL
    result=result+cfb.ContentInt(v.VlanId)+NL
    result=result+cfb.AfterInt()+NL
	}

  return result

}
