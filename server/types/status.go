package types

type Status int

const (
	Available    Status = 1
	NotAvailable        = 2
)

var statusMap = map[Status]string{
	Available:    "Available",
	NotAvailable: "Not Available",
}

func (s Status) String() string {
	if val, ok := statusMap[s]; ok {
		return val
	}

	return ""
}
