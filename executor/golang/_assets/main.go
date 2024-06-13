package main

import "fmt"

func main() {
	//整数hを入力
	var h int
	fmt.Scan(&h)

	h_ := 0
	i := 0
	for {
		//h_に2^iを加える
		h_ += 1 << i

		if h_ > h {
			break
		}

		i++
	}

	fmt.Println(i + 1)
}
