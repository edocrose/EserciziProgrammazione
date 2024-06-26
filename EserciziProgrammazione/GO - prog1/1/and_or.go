package main

import (
  "fmt"
  //"strconv"
  //"strings"
)

/*func main(){
  var grado float64
  grado=7.12342314
  //b:=float64(grado)
  g:=strconv.FormatFloat(grado, 'g', 2, 64)
  fmt.Println(g)

  var s string
  s=strings.Repeat("a", 3)
  fmt.Println(s)

}*/

func main(){
  var x = []int{1,2,3,4,5}
  //x=make([]int){1,2,3,4,5}
  fmt.Println(x)
  for i:=1; i<=len(x); i++{
    fmt.Println(x[i-1])
  }
  fmt.Println(x[:3])
}
