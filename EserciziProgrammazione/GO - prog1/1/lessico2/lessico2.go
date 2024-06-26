package main

import(
  "fmt"
  "bufio"
  "os"
  "strings"
  "sort"
)

func PrintMenu(){
  fmt.Println("+ (legge e memorizza)")
  fmt.Println("? (numeri di riga in cui è comparsa la parola data)")
  fmt.Println("m (stampa le lunghezze min e max)")
  fmt.Println("p (stampa le parole memorizzate)")
}

func PrintDict(m map[string][]int){
  var c []string
  for x, _ := range m { //prendo le chiavi della mappa e le metto in una slice
    c = append(c, x)
  }

  sort.Strings(c) //metto in ordine lessicografico i componenti della slice

  for i, j:= range c{ //parole stampate in modo ordinato
    fmt.Print(c[i], " : ", m[j], "\n")
  }
}

func AddToDict(line string, n int, dict map[string][]int){
  x:=strings.Split(line, " ")
  x=x[1:]
  for i:=0; i<len(x); i++{
    dict[x[i]]= append(dict[x[i]], n)
  }
}

func Stats(dict map[string][]int) (int, int){
  var c []string
  var min, max int
  min=100
  max=0

  for x, _ := range dict { //prendo le chiavi della mappa e le metto in una slice
    c = append(c, x)
  }

  for i:=0; i<len(c); i++{
    //per il minore
    if len(c[i])<min{
      min=len(c[i])
    }
    //per il maggiore
    if len(c[i])>max{
      max=len(c[i])
    }
  }

  return min, max
}

func main(){
  var dict map[string][]int
  dict=make(map[string][]int)
  i:=1

  PrintMenu()

  scanner:=bufio.NewScanner(os.Stdin)
  for scanner.Scan(){
    riga:=scanner.Text()
    parole:=strings.Split(riga, " ")

    switch parole[0]{
    case "+":
      fmt.Println("alimento il dizionario")
      AddToDict(riga, i, dict)

    case "?":
      fmt.Println("consulto il dizionario")
      riga=riga[2:]
      fmt.Println("parola:", riga)
      fmt.Println("righe", dict[riga])

    case "m":
      fmt.Println("lunghezza min e max")
      min, max:=Stats(dict)
      fmt.Println(min, max)

    case "p":
      fmt.Println("stampo il dizionario ordinato")
      PrintDict(dict)

    default:
      fmt.Println("opzione non valida")
    }

    i++
  }


}
