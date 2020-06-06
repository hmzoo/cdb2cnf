package main

import (
	"./cdb2conf"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"strconv"
	"strings"
)

var vlans = []string{"ADMIN", "ELEVE", "PROF", "IMPRESSION", "MANAGEMENT", "DOMOTIQUE", "INVITE"}

func main() {
	fmt.Println("OK")
	fname := "CLG-FS-054-CDB-JFERRY.xlsx"
	ss := strings.Split(fname, "-")
	if len(ss) < 5 {
		fmt.Println("bad filename")
		os.Exit(0)
	}
	clgnum, err := strconv.Atoi(ss[2])
	if err != nil {
		fmt.Println("bad filename")
		os.Exit(0)
	}

	f, err := excelize.OpenFile(fname)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	ports := cdb2conf.NewPorts()
  configs := cdb2conf.NewConfigs()
	rows, err := f.GetRows("CARNET DE BRASSAGE")
	for _, row := range rows {

		if row[8] != "" {
			num, _ := strconv.Atoi(row[9])
			ss := strings.Split(row[8], "-")
			ck := checkVlans(row[10])
			if ck != 0 {
				stackid, _ := strconv.Atoi(ss[len(ss)-1])
				stack := strings.Join(ss[0:len(ss)-1], "-")
				p := cdb2conf.Port{row[6], row[8], num, row[10], clgnum + ck, stack, stackid}
				ports = append(ports, &p)
        configs.AddPort(&p)
			}
		}
	}
  ports.ReOrder();
	for _, p := range ports {
		fmt.Printf("%10s %30s %d %10s %d %10s %d\n", p.Lt, p.Comm, p.Num, p.Vlan, p.Vlanid, p.Stack, p.Stackid)
	}
  configs.Result()
}

func checkVlans(vt string) int {
	for i, v := range vlans {
		if v == vt {
			return (i + 11) * 100
		}
	}
	return 0
}
