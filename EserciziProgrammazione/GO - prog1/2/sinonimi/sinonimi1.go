package main

import(
  "fmt"
  "strings"
  "os"
  "bufio"
)

func main(){

  if len(os.Args)<3{
    fmt.Println("2 file names, please")
    os.Exit(1)
  }

  f1, err:=os.Open(os.Args[1])
  if err!=nil{
    fmt.Println("file 1 not found")
    os.Exit(2)
  }

  defer f1.Close()

  f2, err2:= os.Open(os.Args[2])
  if err2!=nil{
    fmt.Println("file 2 not found")
    os.Exit(2)
  }

  defer f2.Close()

  //mappa per sinonimi
  var sinonimi map[string][]string
  sinonimi=make(map[string][]string)

  scanner1:=bufio.NewScanner(f1)
  for scanner1.Scan(){
    riga:=scanner1.Text()
    if riga!=""{
      div:=strings.Split(riga, ": ")
      sin:=strings.Split(div[1], ", ")
      for _, x:=range sin{
        sinonimi[div[0]]=append(sinonimi[div[0]], x)
      }
    }
  }

  //cercare parola e scrivere sinonimi
  scanner2:=bufio.NewScanner(f2)
  for scanner2.Scan(){
    riga:=scanner2.Text()
    if riga!=""{
      parole:=strings.Split(riga, " ")
      for _, k:=range parole{
        fmt.Print(k)
        for a, _:=range sinonimi{
          if a==k{
            fmt.Print(sinonimi[k])
          }
        }
        fmt.Print(" ")
      }
    }
    fmt.Println()
  }
}
