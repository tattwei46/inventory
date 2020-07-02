package types

type Status int

const (
	Available    Status = 1
	NotAvailable Status = 2
	Deleted      Status = 99
)

var statusMap = map[Status]string{
	Available:    "Available",
	NotAvailable: "Not Available",
	Deleted:      "Deleted",
}

func (s Status) String() string {
	if val, ok := statusMap[s]; ok {
		return val
	}

	return ""
}

func GetStatus(s string) Status {
	switch s {
	case "Available":
		return Available
	case "Not Available":
		return NotAvailable
	case "Deleted":
		return Deleted
	default:
		return 0
	}
}
