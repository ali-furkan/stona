package tools

import (
	"fmt"
	"log"
	"math/rand"
)

func RandomKey(i int) string {
	b := make([]byte, 64)
	n, err := rand.Read(b)
	if err != nil {
		log.Fatalln(err.Error())
		return ""
	}
	return fmt.Sprint(n)
}
