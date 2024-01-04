package maps

import (
	"math"
	"rpgMap/config"
	"rpgMap/global"
	"rpgMap/tool"
)

type PNode struct {
	X      int32
	Y      int32
	G      int32  // 从起始节点到当前节点的代价
	H      int32  // 启发式函数值
	Parent *PNode // 父节点
}

// Arrival 是否到达终点
func arrival(node1, node2 *PNode) bool {
	return node1.X == node2.X && node1.Y == node2.Y
}

func pNode2Point(node *PNode) global.Point {
	return global.Point{X: node.X, Y: node.Y}
}

func getNodeF(node *PNode) int32 {
	return node.G + node.H
}

func (node PNode) InSlice(NodeList []*PNode) bool {
	for _, n := range NodeList {
		if n.X == node.X && n.Y == node.Y {
			return true
		}
	}
	return false
}

// FindPath 寻路
func FindPath(state *MapState, start, end *PNode) []global.Point {
	// 先判断是否能直接到达
	if (start.X == end.X || start.Y == end.Y) && !ContainsObstacleBetween(state, start, end) {
		return []global.Point{pNode2Point(end)}
	}

	// 寻找一个中间点能否间接到达
	if path := CheckIndirect(state, start, end); len(path) > 0 {
		return path
	}
	openList := []*PNode{start}
	var closeList []*PNode
	var neighborMaps = make(map[global.Point]int32)

while:
	if len(openList) <= 0 {
		return nil
	}

	current := openList[0]
	for _, n := range openList {
		nf := getNodeF(n)
		cf := getNodeF(current)
		if nf < cf || (nf == cf && n.H < current.H) {
			current = n
		}
	}
	openList = openList[1:]
	closeList = append(closeList, current)

	if arrival(current, end) {
		return reconstructPath(current)
	}

	point := pNode2Point(current)
	step, ok := neighborMaps[point]
	if !ok {
		step = 1
	}
	neighborMaps[point] = step + 1

	neighbors := getNeighbors(state, current, step)
	if len(neighbors) <= 0 {
		return nil
	}
	for _, neighbor := range neighbors {
		if neighbor.InSlice(closeList) {
			continue
		}
		tentativeG := current.G + CalculateDistance(state, current, neighbor)
		if !neighbor.InSlice(openList) || tentativeG < neighbor.G {
			neighbor.Parent = current
			neighbor.G = tentativeG
			neighbor.H = CalculateDistance(state, neighbor, end)
			//fmt.Println("insert to open list ", current, neighbor, step)
			if !neighbor.InSlice(openList) {
				openList = append(openList, neighbor)
			}
		}
	}
	goto while
}

// CheckIndirect 检查能否通过简单的中间点间接到达
func CheckIndirect(state *MapState, start, end *PNode) (path []global.Point) {
	checkNodes := []PNode{{X: start.X, Y: end.Y}, {X: end.X, Y: start.X}} // 简单判断2个中间点
	for _, mid := range checkNodes {
		if !ContainsObstacleBetween(state, start, &mid) && !ContainsObstacleBetween(state, &mid, end) {
			path = append(path, pNode2Point(&mid), pNode2Point(end))
			return
		}
	}
	return
}

// ReconstructPath 通过node的parent返回一条路径
func reconstructPath(node *PNode) (path []global.Point) {
while:
	if node == nil {
		length := len(path)
		for i := 0; i < length/2; i++ {
			// 交换切片两端的元素
			path[i], path[length-i-1] = path[length-i-1], path[i]
		}
		return
	}
	path = append(path, pNode2Point(node))
	node = node.Parent
	goto while
}

// getNeighbors 返回一个坐标点周围的坐标点 4向
func getNeighbors(state *MapState, node *PNode, step int32) (neighbors []*PNode) {
	dx := []int32{-step, step, 0, 0}
	dy := []int32{0, 0, -step, step}
	for i := 0; i < 4; i++ {
		X := node.X + dx[i]
		Y := node.Y + dy[i]
		newNode := PNode{X: X, Y: Y}
		if IsValidCoordinate(state, X, Y) && !IsObstacleConst(state, X, Y) && !ContainsObstacleBetween(state, node, &newNode) {
			neighbors = append(neighbors, &newNode)
		}
	}
	return
}

// CalculateDistance 计算两个节点的移动代价
func CalculateDistance(state *MapState, node1, node2 *PNode) int32 {
	dx := node1.X - node2.X
	dy := node1.Y - node2.Y
	distance := int32(math.Sqrt(float64(dx*dx+dy*dy) * 10))
	if ContainsObstacleBetween(state, node1, node2) {
		distance *= 5
	}
	return distance
}

// ContainsObstacleBetween 返回两个节点直线路径是否包含障碍物
func ContainsObstacleBetween(state *MapState, start, end *PNode) bool {
	// 检查两个节点之间的直线路径是否包含障碍物
	dx := int32(math.Abs(float64(end.X - start.X)))
	dy := int32(math.Abs(float64(end.Y - start.Y)))

	xIncrement := tool.IF(start.X < end.X, int32(1), int32(-1)).(int32)
	yIncrement := tool.IF(start.Y < end.Y, int32(1), int32(-1)).(int32)

	x := start.X
	y := start.Y
	err := dx - dy

while:
	if x == end.X && y == end.Y {
		return false
	}
	if IsObstacle(state, x, y) && (x != start.X || y != start.Y) {
		return true
	}
	doubleError := err * 2
	if doubleError > -dy {
		err -= dy
		x += xIncrement
	}
	if doubleError < dx {
		err += dx
		y += yIncrement
	}

	goto while
}

// IsValidCoordinate 返回一个坐标是否合法
func IsValidCoordinate(state *MapState, x, y int32) bool {
	mapWidth := state.Config.Width
	mapHeight := state.Config.Height

	return x >= 0 && x <= mapWidth && y >= 0 && y <= mapHeight
}

// IsObstacleConst 返回一个坐标是否是固定阻挡
func IsObstacleConst(state *MapState, x, y int32) bool {
	pos := config.ConfigPos{X: x, Y: y}
	return pos.InSlice(state.Config.UnWalk)
}

// IsObstacle 返回一个坐标是否是阻挡(包括固定阻挡和移动的对象）
func IsObstacle(state *MapState, x, y int32) bool {
	if IsObstacleConst(state, x, y) {
		return true
	}
	// todo 这里需要拿到移动对象
	return false
}
