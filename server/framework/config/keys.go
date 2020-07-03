package config

type Keys int

const (
	MONGODB Keys = 1
)

var keys = map[Keys]string{
	MONGODB: "mongodb",
}

func (k Keys) String() string {
	if v, ok := keys[k]; ok {
		return v
	}
	return ""
}
