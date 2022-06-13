package main

import (
	ct "dijkstra/colourful-terminal"
	nd "dijkstra/nodes"
	hp "dijkstra/helper-functions"
	"fmt"
)

var graph = [len(nd.Node_list)][len(nd.Node_list)]float64{}

func main(){
	ct.Success("Graph initialized")
	for _,n1 := range nd.Node_list{
		for _,n2 := range nd.Node_list{
			var in_routes bool = false
			for _,way := range n1.OpenRoutes{
				if n2.Id == way{in_routes = true}
			}
			if !in_routes{ graph[n1.Id][n2.Id] = 0.0 
			}else { 
				graph[n1.Id][n2.Id] = hp.RoundFloat(
					(hp.DistanceXY(n1.Position,n2.Position)),3) }
		}
		fmt.Println(graph[n1.Id])
	}
}

