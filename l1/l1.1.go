package l1

type Human struct {
	Name string
}

func (h *Human) GetName() string {
	return h.Name
}

type Action struct {
	Human
	Type string
}

func (a *Action) GetType() string {
	return a.Type
}
