package main

import(
	"dijkstra/nodes"
	"fmt"
	"math"
)

var graph = [len(nodes.Node_list)][len(nodes.Node_list)]float64{}

func main(){
	for _,n1 := range nodes.Node_list{
		for _,n2 := range nodes.Node_list{
			var in_routes bool = false
			for _,way := range n1.OpenRoutes{
				if n2.Id == way{in_routes = true}
			}
			if !in_routes{
				graph[n1.Id][n2.Id] = 0.0
			}else { graph[n1.Id][n2.Id] = Distance(n1.Position,n2.Position) }
		}
	}
	fmt.Println(graph)
}

// Returns the distance between two x,y coordinates
func Distance(a [2]float64, b [2]float64) float64 {
  dist := math.Sqrt(math.Pow((a[0]-b[0]),2) + math.Pow((a[1]-b[1]),2))
	return dist
}