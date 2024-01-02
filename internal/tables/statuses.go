package tables

type status int

const (
	Todo status = iota
	InProgress
	Done
)

func (s status) String() string {
	return [...]string{"todo", "in progress", "done"}[s]
}

func (s status) Next() int {
	if s == Done {
		return int(Todo)
	}
	return int(s + 1)
}

func (s status) Previous() int {
	if s == Todo {
		return int(Done)
	}

	return int(s - 1)
}

func (s status) Int() int {
	return int(s)
}
