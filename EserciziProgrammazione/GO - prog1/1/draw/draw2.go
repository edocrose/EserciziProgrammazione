package main

import(
  "fmt"
  "strconv"
  "os"
)

func DrawPoint(c byte, k int)string{
  var s string
  for i:=0; i<k; i++{
    s=s+" "
  }
  s=s+string(c)
  return s
}

func DrawSegment(c byte, k, l int) string{
  var s string
  for i:=0; i<k; i++{
    s=s+" "
  }
  for i:=0; i<l; i++{
    s=s+string(c)
  }
  return s
}

func main(){
  var l, n int
  var c byte
  l, _=strconv.Atoi(os.Args[1])
  n, _=strconv.Atoi(os.Args[2])
  c=[]byte(os.Args[3])[0]
  for i:=0; i<n; i++{
    fmt.Println(DrawSegment(c, i*(l-1), l))
    for j:=0; j<l-1; j++{
      fmt.Println(DrawPoint(c, (i+1)*(l-1)))
    }

  }

}
