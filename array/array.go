package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	capacity, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	inputElements := strings.Fields(scanner.Text())
	if len(inputElements) != capacity {
		fmt.Println("Jumlah elemen tidak sesuai dengan kapasitas array.")
		return
	}

	for _, elem := range inputElements {
		num, _ := strconv.Atoi(elem)

		if num > 2 && num%2 == 0 {
			fmt.Printf("%d ", num)
		}
	}
}
