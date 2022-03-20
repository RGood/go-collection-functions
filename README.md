# go-collection-functions

Repository adding a few functional methods to a package. Mainly for my own personal use.

Example usage:

```golang
elements := []int{1,2,3,4,5}

newElements := Map(elments, func(e int) int {
    return e * 2
})

for _, element := range newElements {
    println(element)
}
```