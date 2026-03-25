package main

import "fmt"

func main() {
	var h, p, f, d int
	fmt.Scan(&h, &p, &f, &d)

	if d == 1 {
		for i := f; ; i++ {
			i = (i + 16) % 16

			if i == p {
				fmt.Println("N")
				break
			} else if i == h {
				fmt.Println("S")
				break
			}
		}
	} else {
		for i := f; ; i-- {
			i = (i + 16) % 16

			if i == p {
				fmt.Println("N")
				break
			} else if i == h {
				fmt.Println("S")
				break
			}
		}
	}
}
