package Validator

// InternalSubnet parameters
type User struct {
	// (mandatory) list of subnets to be used
	Pool []string `json:"pool" yaml:"pool" validate:"max=10"`
}

func NewUser(ips []string) *User {
	return &User{Pool: ips}
}
