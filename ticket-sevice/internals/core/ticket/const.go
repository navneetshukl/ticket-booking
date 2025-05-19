package ticket

type Status string

const (
	ERROR   Status = "error"
	SUCCESS Status = "status"
)

func (s Status) ToString() string {
	return string(s)
}
