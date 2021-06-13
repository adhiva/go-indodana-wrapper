package indodana

type Env int8

const (
	_ Env = iota

	// Sandbox : represent sandbox environment
	Sandbox

	// Production : represent production environment
	Production
)

var typeString = map[Env]string{
	Sandbox:    "https://sandbox01-api.indodana.com/chermes/merchant",
	Production: "https://api.indodana.com/chermes/merchant",
}

// implement stringer
func (e Env) String() string {
	for k, v := range typeString {
		if k == e {
			return v
		}
	}
	return "undefined"
}
