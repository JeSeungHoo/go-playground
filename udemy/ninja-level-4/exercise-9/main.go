package main

func main() {
	// Using the code from the previous example, add a record to your map. Now print the map out
	// using the “range” loop

	m := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "Sunsets"},
	}

	m["hoo"] = []string{"game", "go", "ys"}

	for k, v := range m {
		println(k, "Likes")
		for i, val := range v {
			println("\t", i, val)
		}
	}
}
