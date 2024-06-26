package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sort"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2 file names, please")
		os.Exit(1)
	}

	file1 := os.Args[1]

	file2 := os.Args[2]

	f, err := os.Open(file1)

	if err != nil {
		fmt.Println("file 1 not found")

		os.Exit(2)
	}

	defer f.Close()

	f2, err2 := os.Open(file2)

	if err2 != nil {
		fmt.Println("file 2 not found")

		os.Exit(2)
	}

	defer f2.Close()

	var dizionario map[string][]string
	dizionario = make(map[string][]string)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		contenitore := strings.Split(line, ":")
		m := strings.Split(contenitore[1], ",")

		for i := range m {
			m[i] = strings.TrimSpace(m[i])
		}
		sort.Strings(m)
		dizionario[contenitore[0]] = m
	}
	scanner2 := bufio.NewScanner(f2)
	for scanner2.Scan() {

		line := scanner2.Text()

		if line == " " {
			fmt.Println()
			continue
		}

		parole := strings.Split(line, " ")


		for _, k := range parole {
			fmt.Print(k)
			for i, _ := range dizionario {

				if k == i {
					fmt.Print(dizionario[i])
				}
			}
			fmt.Print(" ")

		}
		fmt.Println()
	}
}
