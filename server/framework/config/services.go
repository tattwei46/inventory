package config

type Services int

const (
	INVENTORY Services = 1
)

var services = map[Services]string{
	INVENTORY: "inventory-service",
}

func (s Services) String() string {
	if v, ok := services[s]; ok {
		return v
	}
	return ""
}
