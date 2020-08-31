package main

import "fmt"

type Node struct {
	row int
	col int
	value int
}

func main()  {
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	var sparsearray []Node

	valNode := Node{
		row:   11,
		col:   11,
		value: 0,
	}
	sparsearray = append(sparsearray, valNode)

	for i,value :=range chessMap{
		for j,v2 := range value {
			if v2 != 0 {
				node := Node{
					row:i,
					col:j,
					value:v2,
				}
				sparsearray = append(sparsearray, node)
			}
		}
	}
	for i,node := range sparsearray{
		fmt.Printf("%d: %d %d %d\n",i,node.row,node.col,node.value)
	}
}
