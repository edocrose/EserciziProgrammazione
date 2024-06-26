package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
)

func CalcQuanteMinMax(frase string) (quante, min, max int){
  if frase==""{
    return quante, min, max
  }
  min=999
  max=0
  contenitore:=strings.Split(frase, " ")
  for _, x:=range contenitore{
    parole:=strings.Split(x, ",")
    if strings.IndexAny(parole[0], "1234567890") == -1{
      quante++
      if len(parole[0])<min{
        min=len(parole[0])
      }
      if len(parole[0])>max{
        max=len(parole[0])
      }
    }
  }

  return quante, min, max
}

func TrovaNumDopoKeyword(frase, keyWord string) (num int){
  var numString string
  parole:=strings.Split(frase, " ")
  for i:=0; i<len(parole); i++{
    if parole[i]==keyWord{
      NoVirg:=strings.Split(parole[i+1], ",")
      numString=NoVirg[0]
      break
    }
  }
  if numString==""{
    return 0
  }
  num,_=strconv.Atoi(numString)
  return num
}

func Autogramma(frase string) bool{
  quante, min, max:=CalcQuanteMinMax(frase)
  numQuante:=TrovaNumDopoKeyword(frase, "contiene:")
  numMin:=TrovaNumDopoKeyword(frase, "minima:")
  numMax:=TrovaNumDopoKeyword(frase, "massima:")
  if quante==numQuante && min==numMin && max==numMax{
    return true
  }
  return false
}

func StampaAnalisiAutogramma(frase string){
  quante, min, max:=CalcQuanteMinMax(frase)
  numQuante:=TrovaNumDopoKeyword(frase, "contiene:")
  numMin:=TrovaNumDopoKeyword(frase, "minima: ")
  numMax:=TrovaNumDopoKeyword(frase, "massima: ")

  fmt.Println("===", frase)
  fmt.Printf("minimo dichiarato %d contro minimo verificato %d\n", numMin, min)
  fmt.Printf("massimo dichiarato %d contro massimo verificato %d\n", numMax, max)
  fmt.Printf("numero parole dichiarato %d contro numero parole verificato %d\n", numQuante, quante)
  if Autogramma(frase){
    fmt.Println("onesto")
  }else{
    fmt.Println("bugiardo")
  }
}

func main(){
  var frasi []string

  if len(os.Args)<2{
    fmt.Println("file name?")
    os.Exit(1)
  }

  f,err:=os.Open(os.Args[1])

  if err!=nil{
    fmt.Println("file not found")
    os.Exit(2)
  }

  scanner:=bufio.NewScanner(f)
  for scanner.Scan(){
    riga:=scanner.Text()
    if riga!=""{
      frasi=append(frasi, riga)
    }
  }

  for _, x:=range frasi{
    StampaAnalisiAutogramma(x)
  }
}
