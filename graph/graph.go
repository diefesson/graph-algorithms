package graph

import (
	"container/list"
)

func clearVisited(g Graph) {
	g.ForVertices(func(_ Graph, index int) { g.setVisited(index, false) })
}

func ForPresentVertices(g Graph, vp VertexProcessor) {
	g.ForVertices(FilterVertices(PresentVertexFilter, vp))
}

func ForConnectedEdges(g Graph, ep EdgeProcessor) {
	g.ForEdges(FilterEdges(ConnectedEdgeFilter, ep))
}

func Neighbours(g Graph, index int) []int {
	neighbours := make([]int, 0)
	g.ForNeighbours(index, func(_ Graph, i int) { neighbours = append(neighbours, i) })
	return neighbours
}

func DepthFirst(g Graph, start int, vp VertexProcessor, ep EdgeProcessor) {
	if vp == nil {
		vp = emptyVertexProcessor
	}
	if ep == nil {
		ep = emptyEdgeProcessor
	}
	clearVisited(g)
	depthFirst(g, start, vp, ep)
}

func depthFirst(g Graph, index int, vp VertexProcessor, ep EdgeProcessor) {
	g.setVisited(index, true)
	vp(g, index)
	for _, i := range Neighbours(g, index) {
		if !g.isVisited(i) {
			ep(g, index, i)
			depthFirst(g, i, vp, ep)
		}
	}
}

func BreadthFirst(g Graph, start int, vp VertexProcessor) {
	if vp == nil {
		vp = emptyVertexProcessor
	}
	clearVisited(g)
	breadthFirst(g, start, vp)
}

func breadthFirst(g Graph, start int, vp VertexProcessor) {
	nexts := list.New()
	nexts.PushBack(start)
	g.setVisited(start, true)
	for e := nexts.Front(); e != nil; e = nexts.Front() {
		index := e.Value.(int)
		nexts.Remove(e)
		vp(g, index)
		for _, i := range Neighbours(g, index) {
			if !g.isVisited(i) {
				g.setVisited(i, true)
				nexts.PushBack(i)
			}
		}
	}
}

func IsConnectedFrom(g Graph, start int) bool {
	count := 0
	counter := func(_ Graph, _ int) { count++ }
	DepthFirst(g, start, counter, nil)
	return count == g.VertexCount()
}
