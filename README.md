# go-collection-functions

Repository adding a few functional methods to a package. Mainly for my own personal use.

Example usage:

Map:

```golang
elements := []int{1, 2, 3, 4, 5}

// Multiply each element by 2
newElements := collections.Map(elements, func(e int) int {
    return e * 2
})

//Convert the elments to a string and join them on ", "
println(
    strings.Join(
        collections.Map(newElements, func(e int) string {
            return fmt.Sprintf("%d", e)
        }),
        ", ",
    ),
)
// => "2, 4, 6, 8, 10"
```