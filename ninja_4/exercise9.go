package main

import "fmt"

func main() {
	s := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	s["new_key"] = []string{`James Bond`, `Literature`, `Computer Science`}

	for k, v := range s {
		fmt.Println("Record", k)
		for i, v1 := range v {
			fmt.Println(i, v1)
		}
	}

}
