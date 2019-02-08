package main

func must(e error) {
	if e != nil {
		panic(e)
	}
}

func getIndex(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
