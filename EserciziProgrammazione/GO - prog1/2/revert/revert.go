package main

import(
  "fmt"
  "strings"
)

func RighePari(r string){
  parole:=strings.Fields(r)
  for _, x:=range parole{
    if len(x)%2==0{
      for i:=len(x)-1; i>=0; i--{
        fmt.Printf("%c", x[i])
      }
      fmt.Println()
    }
  }


func main(){
  var r string
  for{
    fmt.Scan(&r)
    if r=="stop"{
      break
    }
    RighePari(r)
  }
}
