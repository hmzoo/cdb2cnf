package cdb2conf

import (
	"fmt"
	"strconv"
)

type Config struct {
	Name  string
	Ports Ports
}

type Configs []*Config

func NewConfigs() Configs {
	return Configs{}
}

func (cfs *Configs) AddPort(p *Port) {

	for _, c := range *cfs {

		if c.Name == p.Stack {
			c.Ports = append(c.Ports, p)
			return
		}
	}
	ncf := Config{p.Stack, Ports{}}

	*cfs = append(*cfs, &ncf)
	cfs.AddPort(p)
	fmt.Println(len(*cfs))
}

func (cf Configs) Result() {

	for _, c := range cf {
		fmt.Println(c.Name)
		c.Ports.ReOrder()
		curVlan := 0
		curStack := 0
		curNums := make([]int, 0)
		curText := ""

		for _, p := range c.Ports {
			if p.Vlanid != curVlan {
				if curVlan != 0 {
					curText = curText + fmtNums(curNums)
					curNums = make([]int, 0)
					curText = curText + "\nswitchport access vlan " + strconv.Itoa(curVlan) + "\nexit\n"
				}
				curText = curText + "\ninterface range ethernet "
				curStack = 0
				curVlan = p.Vlanid
			}

			if p.Stackid != curStack {
				if curStack != 0 {
					curText = curText + fmtNums(curNums) + ", "
				}
				curText = curText + strconv.Itoa(p.Stackid) + "/e"
				curStack = p.Stackid
				curNums = make([]int, 0)
			}

			curNums = append(curNums, p.Num)
			//curText =curText+strconv.Itoa(p.Num)+"-"
		}
		curText = curText + fmtNums(curNums)
		curText = curText + "\nswitchport access vlan " + strconv.Itoa(curVlan) + "\nexit\n"
		fmt.Println(curText)
	}

}

func fmtNums(nums []int) string {

	if len(nums) == 0 {
		return ""
	}
	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}
	r := ""
	curnum := 0
	for n, i := range nums {
		if n == 0 {
			curnum = i
			r = r + strconv.Itoa(i)
		} else {
			if curnum+1 == i {
				r = r + "-"
				curnum = curnum + 1
			} else {
				if curnum != i {
					r = r + strconv.Itoa(curnum)
				}
				r = r + "," + strconv.Itoa(i)
				curnum = i
			}
		}

	}
	return "(" + r + ")"
}
