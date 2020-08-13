package main

import "fmt"

func main() {
	s := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	for k, v := range s {
		fmt.Println("Record", k)
		for i, v1 := range v {
			fmt.Println(i, v1)
		}
	}
	fmt.Println(s)
}
