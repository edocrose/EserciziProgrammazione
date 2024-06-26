package main

import (
  //"fmt"
  "strconv"
)

type Point struct{
  x, y float64
}

type Rectangle struct{
  pLL, pUR Point //pLL vertice in basso a sinistra
                //pUR vertice in alto a destra
}

func NewPoint(x, y float64) Point{
  var p Point
  p.x=x
  p.y=y
  return p
}

func NewRectangle(p1, p2 Point) Rectangle{
  var r Rectangle
  if p1.x>p2.x{
    p1.x,p2.x=p2.x,p1.x
  }
  if p1.y>p2.y{
    p1.y,p2.y=p2.y,p1.y
  }
  r.pLL=p1
  r.pUR=p2
  return r

}

func Rotate(r *Rectangle, dir byte){
  p1x:=r.pLL.x
  p1y:=r.pLL.y
  p2x:=r.pUR.x
  p2y:=r.pUR.y
  if dir=='L'{
    r.pLL.x=p1x-p2y
    r.pLL.y=p1y
    r.pUR.x=p1x
    r.pUR.y=p2x-p1x

  }else if dir=='R'{
    r.pLL.x=p1x
    r.pLL.y=p1x-p2x
    r.pUR.y=p1y
    r.pUR.x=p1x+p2y
  }
}

func (r Rectangle) String() string{
  p1x:=strconv.FormatFloat(r.pLL.x, 'E', -1, 64)
  p1y:=strconv.FormatFloat(r.pLL.y, 'E', -1, 64)
  p2x:=strconv.FormatFloat(r.pUR.x, 'E', -1, 64)
  p2y:=strconv.FormatFloat(r.pUR.y, 'E', -1, 64)
  s:="(("+p1x+","+p1y+") ("+p2x+","+p2y+"))"
  return s
}

func main(){

}
