package main

import(
  "fmt"
  "os"
  "bufio"
)

func main(){
  if len(os.Args) < 2 {
      fmt.Println("il nome del file?")
      os.Exit(1)
  }

  f, err := os.Open(os.Args[1])
  if err != nil {
      fmt.Println("file not found")
      os.Exit(2)
  }
  scanner:=bufio.NewScanner(f)
  for scanner.Scan(){
    riga:=scanner.Text()
    fmt.Println(riga)
  }

  defer f.Close()
}
