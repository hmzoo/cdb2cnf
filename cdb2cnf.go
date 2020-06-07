package main

import (
	"./cdb2cnf"
	"fmt"
  "flag"
  "os"
  "strings"
  "path/filepath"
)

var confs = make(map[string]*cdb2cnf.Conf)
var format string
var outputdir string
var builder cdb2cnf.ConfBuilder

func main() {
  flag.Usage = func() {
        fmt.Println("\ncdb2cnf  [-f FORMAT] [-w [-o REPERTOIRE]] [FICHIER_CDB]")
        flag.PrintDefaults()
    }
  flag.StringVar(&format,"f","8000S", "type de formatage : 8000S , FS980 ")
  writeFlag  :=flag.Bool("w",false, "ecriture des fichiers  ")
  outputFlag  :=flag.String("o",".", "choix du repertoire ")
  flag.Parse()
  format =strings.ToTitle(format)
 

  if filename := flag.Arg(0); filename != "" {
	   err,ports:= cdb2cnf.LoadCDBPNC (filename)
     if(err !=nil){
       fmt.Println(err)
       os.Exit(1)
     }
     err=buildconfs(ports)
     if(err !=nil){
       fmt.Println(err)
       os.Exit(1)
     }
   }else{
     flag.Usage()
     os.Exit(1)
   }


  switch format {
  case "FS980":
      builder= cdb2cnf.T_FS980{}
      default :
      builder= cdb2cnf.T_8000S{}
  }

if(*writeFlag ){
  outputdir=*outputFlag
  save(builder)
}else{
  print(builder)
}
}

func buildconfs(ports cdb2cnf.Ports) error {
  for _, p := range ports {
    if _, ok := confs[p.Stack]; !ok {
      confs[p.Stack]=cdb2cnf.NewConf(p.Stack)
    }
      confs[p.Stack].AddPorts(p.Stackid, p.Num, p.Vlanid)
	}
  return nil
}

func print(builder cdb2cnf.ConfBuilder){
  for k,v := range confs {
    r := cdb2cnf.Build(v,builder)
    fmt.Println("--------------",k)
    fmt.Println(r)

  }
}

func save(builder cdb2cnf.ConfBuilder){
  if err := os.MkdirAll(outputdir, 0777 ); err != nil {
    fmt.Println(err)
    os.Exit(1)
	}
  for k,v := range confs {
    data := cdb2cnf.Build(v,builder)
    filename :=filepath.Join(outputdir, k+"_INT.txt")
    saveinfile(filename,data)

  }
}

func saveinfile(filename,data string){
  file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  defer file.Close()
  fmt.Fprintf(file, data)
}
