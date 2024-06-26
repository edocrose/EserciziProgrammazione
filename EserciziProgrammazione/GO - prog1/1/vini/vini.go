package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
)

type BottigliaVino struct{
  nome string
	anno int
	grado float32
	quantita int
}

func CreaBottiglia(nome string, anno int, grado float32, quantita int) (BottigliaVino, bool){
  var x BottigliaVino
  if nome != "" && anno > 0 && grado > 0 && quantita > 0{
    x.nome=nome
    x.anno=anno
    x.grado=grado
    x.quantita=quantita
    return x, true
  }
  return x, false
}

func CreaBottigliaDaRiga(riga string) (BottigliaVino, bool){
  parole:=strings.Split(riga, ",")
  var nome string
  var anno int
  var grado float32
  var quantita int
  for i:=0; i<len(parole); i++{
    switch i{
    case 0:
      nome=parole[0]
    case 1:
      a,_:=strconv.Atoi(parole[1])
      anno=a
    case 2:
      g,_:=strconv.ParseFloat(parole[2], 64)
      grado=float32(g)
    case 3:
      q,_:=strconv.Atoi(parole[3])
      quantita=q
    }
  }
  return CreaBottiglia(nome, anno, grado, quantita)
}

func AggiungiBottiglia(bot BottigliaVino, cantina *[]BottigliaVino){
  *cantina=append(*cantina, bot)
}

func AggiungiBottigliaDaRiga(bot string, cantina *[]BottigliaVino){
  parole:=strings.Split(bot, ",")
  nome:=parole[0]
  anno,_:=strconv.Atoi(parole[1])
  g,_:=strconv.ParseFloat(parole[2], 64)
  grado:=float32(g)
  quantita,_:=strconv.Atoi(parole[3])
  b,_:= CreaBottiglia(nome, anno, grado, quantita)
  AggiungiBottiglia(b, cantina)
}

func ContaPerTipo (nome string, cantina []BottigliaVino) int{
  c:=0
  for i:=0; i<len(cantina); i++{
    if cantina[i].nome==nome{
      c++
    }
  }
  return c
}

func (b BottigliaVino) String() string{
  var s string
  anno:=strconv.Itoa(b.anno)
  g:=float64(b.grado)
  grado:=strconv.FormatFloat(g, 'g', -1, 64)
  quantita:=strconv.Itoa(b.quantita)
  s=b.nome+","+anno+","+grado+","+quantita
  //return fmt.Sprintf(s)
  return s
}

func main(){
  var cantina []BottigliaVino
  scanner:=bufio.NewScanner(os.Stdin)
  for scanner.Scan(){
    riga:=scanner.Text()
    if riga == ""{
      continue
    }else {
      bot, x:= CreaBottigliaDaRiga(riga)
      if x==true{
        AggiungiBottiglia(bot, &cantina)
      }
    }
  }
  for i:=0; i<len(cantina); i++{
    //fmt.Println(cantina[i].String())
    fmt.Print(cantina[i])
    fmt.Println()
  }
}
