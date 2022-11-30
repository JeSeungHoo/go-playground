package main

import "fmt"

// Take the code from the previous exercise, then store the values of type person in a map with the
// key of last name. Access each value in the map. Print out the values, ranging over the slice.

type person struct {
	first  string
	last   string
	flavor []string
}

func main() {
	p1 := person{
		first:  "SeungHoo",
		last:   "Je",
		flavor: []string{"Chocolate", "banana"},
	}
	p2 := person{
		first:  "James",
		last:   "Bond",
		flavor: []string{"banila", "milk"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for k, v := range m {
		fmt.Println("key is", k)
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, f := range v.flavor {
			fmt.Println(i, f)
		}
	}

}
