package dto

// Estructuras que representan lo que se recibe en las solicitudes CreateUser y UpdateUser.
type CreateUser struct {
	Name  string
	Email string
}

type UpdateUser struct {
	Name  string
	Email string
}
