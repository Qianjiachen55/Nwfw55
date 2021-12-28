package framework

import (
	"errors"
	"strings"
)

type Tree struct {
	root *node
}

type node struct {
	isLast  bool
	segment string
	handlers []ControllerHandler //中间件 && 控制器
	childs  []*node
	parent *node
}

func NewTree()*Tree  {
	root := newNode()
	return &Tree{root: root}
}

func newNode()*node  {
	return &node{
		isLast: false,
		segment: "",
		childs: []*node{},
		parent: nil,
	}
}

func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

func (n *node) filterChildNodes(segment string) []*node {
	if len(n.childs) == 0 {
		return nil
	}

	nodes := make([]*node, 0, len(n.childs))

	for _, cNodes := range n.childs {
		if isWildSegment(cNodes.segment) {
			nodes = append(nodes, cNodes)
		} else if cNodes.segment == segment {
			nodes = append(nodes, cNodes)
		}
	}
	return nodes
}

func (n *node) matchNode(uri string) *node {

	segments := strings.SplitN(uri, "/", 2)

	segment := segments[0]

	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}

	cNodes := n.filterChildNodes(segment)

	//segment无法匹配
	if cNodes == nil || len(cNodes) == 0 {
		return nil
	}

	// segments 只有一个
	if len(segments) == 1 {
		for _, tn := range cNodes {
			if tn.isLast {
				return tn
			}
		}
		return nil
	}
	for _, tn := range cNodes {
		tnMatch := tn.matchNode(segments[1])
		if tnMatch != nil {
			return tnMatch
		}
	}
	return nil

}

func (tree *Tree) AddRouter(uri string, handlers []ControllerHandler) error {
	n := tree.root

	if n.matchNode(uri) != nil{
		return errors.New("route exist: " + uri )
	}

	segments := strings.Split(uri,"/")

	for index ,segment := range segments{
		if !isWildSegment(segment){
			segment = strings.ToUpper(segment)
		}
		isLast := index == len(segments) -1

		var objNode *node

		childNodes := n.filterChildNodes(segment)

		if len(childNodes) > 0{
			for _, childNode := range childNodes{
				if childNode.segment == segment{
					objNode = childNode
					break
				}
			}
		}

		if objNode == nil{
			cNode := newNode()
			cNode.segment = segment
			if isLast{
				cNode.isLast = true
				cNode.handlers = handlers
			}
			cNode.parent = n
			n.childs = append(n.childs, cNode)
			objNode = cNode
		}

		n = objNode
	}
	return nil
}

func (n *node) parseParamsFromEndNode(uri string) map[string]string {
	ret := map[string]string{}
	segments := strings.Split(uri,"/")

	cnt := len(segments)
	cur := n
	for i:=cnt-1;i>0;i--{
		if cur.segment==""{
			break
		}
		if isWildSegment(cur.segment){
			ret[cur.segment[1:]] = segments[i]
		}
		cur = cur.parent
	}
	return ret
}



//