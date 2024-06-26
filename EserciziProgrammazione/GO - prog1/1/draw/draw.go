package main

import(
  "fmt"
  "os"
  "strconv"
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
  l, _= strconv.Atoi(os.Args[1])
  n, _= strconv.Atoi(os.Args[2])
  c = []byte(os.Args[3])[0]
  if l>0 && n>0{
    for k:=1; k<=n; k++{
      fmt.Println(DrawSegment(c, (k-1)*(l-1), l))
      for i:=1; i<l; i++{
        fmt.Println(DrawPoint(c, k*(l-1)))
      }
    }
  }

}
