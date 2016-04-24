package auth

type NoAuthenticator struct {

}

func (na NoAuthenticator) GetPermissions(authorization string) map[string]Permission {
	d := map[string]Permission{};
	d["*"] = Permission{true, true}
	return d
}

func init() {
	na := NoAuthenticator{}
	RegisterAuthenticator("None", na)
}
