package main

import(
  "fmt"
  "bufio"
  "os"
  "strings"
)



func main(){
  fmt.Print("+ (legge e memorizza)\n? (numeri di riga in cui è comparsa la parola data)\nm (stampa le lunghezze min e max)\np (stampa le parole memorizzate)\n")
  scanner:=bufio.NewScanner(os.Stdin)
  for scanner.Scan(){
    riga:=scanner.Text()
    parole:=strings.Split(riga, " ")
    switch parole[0]{
    case "+":
      fmt.Println("alimento il dizionario")
    case "?":
      fmt.Println("consulto il dizionario")
    case "m":
      fmt.Println("lunghezza min e max")
    case "p":
      fmt.Println("stampo il dizionario ordinato")
    default:
      fmt.Println("opzione non valida")
    }
  }
  /*i:=0
  for{
    p, err := fmt.Scan(&var)
    if err != nil{
      break
    }
    frasi=append(frasi, p)
    i++
  }*/


}
