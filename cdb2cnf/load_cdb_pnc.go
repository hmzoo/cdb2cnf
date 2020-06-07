package cdb2cnf

import(
  "strconv"
  "strings"
  "errors"
  "path"
"github.com/360EntSecGroup-Skylar/excelize"
)

var pncvlans = []string{"ADMIN", "ELEVE", "PROF", "IMPRESSION", "MANAGEMENT", "DOMOTIQUE", "INVITE"}
const (
  CDBSHEET = "CARNET DE BRASSAGE"
  R_NUM = 9
  R_COMM = 8
  R_VLAN = 10
  R_REP = 6
)

func LoadCDBPNC(fpath string) (error,Ports) {

ports := NewPorts()

_,fname := path.Split(fpath)

ss := strings.Split(fname, "-")
if len(ss) < 5 {
  return errors.New("Wrong file name !"),ports
}
clgnum, err := strconv.Atoi(ss[2])
if err != nil {
  return errors.New("Wrong file name !"),ports
}

f, err := excelize.OpenFile(fname)
if err != nil {
  return err,ports
}



rows, err := f.GetRows("CARNET DE BRASSAGE")
for l, row := range rows {
  num, err := strconv.Atoi(row[R_NUM])
  if err==nil  && row[R_COMM] != "" {
    ss := strings.Split(row[R_COMM], "-")
    ck := checkVlans(row[R_VLAN])
    if len(ss)==4 && ck != 0 {
      stackid, err := strconv.Atoi(ss[len(ss)-1])
      if err != nil {
        return errors.New("Wrong format for Commutateur, line "+strconv.Itoa(l+1)),ports
      }
      stack := strings.Join(ss[0:len(ss)-1], "-")
      p := Port{row[R_REP], row[R_COMM], num, row[R_VLAN], clgnum + ck, stack, stackid}
      ports = append(ports, &p)
    }
  }
}
ports.ReOrder();
return nil,ports

}

func checkVlans(vt string) int {
	for i, v := range pncvlans {
		if v == vt {
			return (i + 11) * 100
		}
	}
	return 0
}
