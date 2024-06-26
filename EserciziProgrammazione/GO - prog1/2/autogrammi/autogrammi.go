package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
	"bufio"
)

func CalcQuanteMinMax(frase string) (quante, min, max int) {
	if frase==""{
		return quante, min, max
	}
	contenitoreParole := strings.Split(frase, " ")
	max = 0
	min = 999

	for _,j := range contenitoreParole {
		contenitoreNoVirg:= strings.Split(j,",")
		contenitore:= strings.Split(contenitoreNoVirg[0],":")

		if strings.IndexAny(contenitore[0], "1234567890") == -1 {
			quante++
			c := len(contenitore[0])
			if min > c {
				min = c
			}
			if c > max {
				max = c
			}

		}
	}
	return quante, min, max
}

func TrovaNumDopoKeyword(frase, keyWord string) (num int) {
	var t string
	contenitore := strings.Split(frase, " ")

	for i:= 0; i < len(contenitore); i++ {
		if contenitore[i] == keyWord {
			t = contenitore[i + 1]
			break
		}
	}

  if strings.IndexAny(t, ",")!=-1{
    virgole:=strings.Split(t, ",")
    t=virgole[0]
  }

	n, _ := strconv.Atoi(t)
	num = n
	return num
}

func Autogramma(frase string) bool {
	var c int = 0
	quante,min,max := CalcQuanteMinMax(frase)
	numMax:= TrovaNumDopoKeyword(frase, "massima:")
	numMin:= TrovaNumDopoKeyword(frase, "minima:")
	numQuante:= TrovaNumDopoKeyword(frase, "contiene:")

	if quante==numQuante{
		c++
	}
	if min==numMin{
		c++
	}
	if max==numMax{
		c++
	}
	if c == 3 {
		return true
	}
	return false
}

func StampaAnalisiAutogramma(frase string) {
	quante,min,max := CalcQuanteMinMax(frase)
	numMax:= TrovaNumDopoKeyword(frase, "massima:")
	numMin:= TrovaNumDopoKeyword(frase, "minima:")
	numQuante:= TrovaNumDopoKeyword(frase, "contiene:")

	fmt.Println("=== ", frase)

	fmt.Printf("minimo dichiarato %d contro minimo verificato %d\n",numMin,min)

	fmt.Printf("massimo dichiarato %d contro massimo verificato %d\n",numMax,max)

	fmt.Printf("numero parole dichiarato %d contro numero parole verificato %d\n",numQuante,quante)
}

func main () {
	var succ bool
	var frasi []string

	if len(os.Args) <2 {
		fmt.Println("file name?")
		os.Exit(1)
	}

	f,err := os.Open(os.Args[1])
	if err != nil {
	fmt.Println("file not found")
	os.Exit(2)
	}

	scanner:= bufio.NewScanner(f)
	for scanner.Scan() {
		riga:= scanner.Text()
		if riga!=""{
			frasi = append(frasi, riga)
		}
	}

	for _,x:= range frasi {
		succ = Autogramma(x)
		StampaAnalisiAutogramma(x)
		if succ == true {
			fmt.Println("onesto")
		} else {
			fmt.Println("bugiardo")
		}
	}

}
