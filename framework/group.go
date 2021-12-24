package framework

type IGroup interface {
	Get(string,ControllerHandler)
	Put(string,ControllerHandler)
	Post(string,ControllerHandler)
	Delete(string,ControllerHandler)

	Group(string) IGroup
}

type Group struct {
	core *Core
	parent *Group
	prefix string
}

func NewGroup(core *Core,prefix string) *Group {
	return &Group{
		core:   core,
		parent: nil,
		prefix: prefix,
	}
}

func (g *Group) Get(uri string,handler ControllerHandler) {
	uri = g.prefix + uri
	g.core.Get(uri,handler)
}

func (g *Group) Post(uri string,handler ControllerHandler){
	uri = g.prefix + uri
	g.core.Post(uri,handler)
}

func (g *Group) Put(uri string,handler ControllerHandler){
	uri = g.prefix + uri
	g.core.Put(uri,handler)
}

func (g *Group) Delete(uri string,handler ControllerHandler){
	uri = g.prefix + uri
	g.core.Delete(uri,handler)
}

func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil{
		return g.prefix
	}
	return g.parent.getAbsolutePrefix() + g.prefix
}

func (g *Group) Group(uri string) IGroup {

	cgroup := NewGroup(g.core,uri)
	cgroup.parent = g

	return cgroup
}