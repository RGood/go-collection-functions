package main

import (
	"fmt"
	"strings"

	"github.com/RGood/go-collection-functions/pkg/collection"
)

func main() {
	elements := []int{1, 2, 3, 4, 5}

	newElements := collection.Map(elements, func(e int) int {
		return e * 2
	})

	println(
		strings.Join(
			collection.Map(newElements, func(e int) string {
				return fmt.Sprintf("%d", e)
			}),
			", ",
		),
	)
}
