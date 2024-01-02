package tables

type status int

const (
	todo status = iota
	inProgress
	done
)

func (s status) String() string {
	return [...]string{"todo", "in progress", "done"}[s]
}

func (s status) Next() int {
	if s == done {
		return int(todo)
	}
	return int(s + 1)
}

func (s status) Previous() int {
	if s == todo {
		return int(done)
	}

	return int(s - 1)
}

func (s status) Int() int {
	return int(s)
}
