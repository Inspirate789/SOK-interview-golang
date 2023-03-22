package main

import (
	"fmt"
	"log"
)

func main() {
	var n uint
	fmt.Print("Enter the column number: ")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		log.Fatal(err)
	}

	n--
	strLen := int(n/26) + 1
	strOffset := n % 26

	res := make([]byte, 0, strLen)
	for i := 0; i < strLen-1; i++ {
		res = append(res, 'A')
	}
	res = append(res, byte('A'+strOffset))

	fmt.Println("Column name:", string(res))
}
