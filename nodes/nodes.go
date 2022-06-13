package nodes

type Node struct {
  Position [2]float64
  Id  uint
	OpenRoutes []uint
}


var node_0 = Node{
  Position: [2]float64{0.0,0.0},
  Id: 0,
	OpenRoutes: []uint{1,5},
}

var node_1 = Node{
  Position: [2]float64{2.0,1.0},
  Id: 1,
	OpenRoutes: []uint{0,2,3,4,5},
}

var node_2 = Node{
  Position: [2]float64{5.0,1.0},
  Id: 2,
	OpenRoutes: []uint{1,3,6,8},
}

var node_3 = Node{
  Position: [2]float64{4.0,2.0},
  Id: 3,
	OpenRoutes: []uint{1,2,4,6},
}

var node_4 = Node{
  Position: [2]float64{3.0,3.0},
  Id: 4,
	OpenRoutes: []uint{1,3,5,6,7},
}

var node_5 = Node{
  Position: [2]float64{1.0,4.0},
  Id: 5,
	OpenRoutes: []uint{0,1,4,7},
}

var node_6 = Node{
  Position: [2]float64{4.0,4.0},
  Id: 6,
	OpenRoutes: []uint{2,3,4,7,8},
}

var node_7 = Node{
  Position: [2]float64{3.0,5.0},
  Id: 7,
	OpenRoutes: []uint{4,5,6,8},
}

var node_8 = Node{
  Position: [2]float64{5.0,5.0},
  Id: 8,
	OpenRoutes: []uint{2,6,7},
}

var Node_list = [9]Node{
	node_0,
	node_1,
	node_2,
	node_3,
	node_4,
	node_5,
	node_6,
	node_7,
	node_8,
}
