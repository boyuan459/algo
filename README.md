## Install
```
go get -u github.com/boyuan459/algo
```

## Example
```
package main

import (
	"fmt"

	"github.com/boyuan459/algo/graph"
	"github.com/boyuan459/algo/list"
)

func main() {
	graph := graph.New(8)
	list := list.New[int]()
	list.Push(10)
	fmt.Println("graph:", graph)
	fmt.Println("list:", list)
}
```

## How to publish
```
git tag "vX.X.X"
git push --tags
```
