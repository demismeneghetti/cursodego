package model

//Usuario representa um usuario do sistema
type Usuario struct {
	Nome  string `json:"name"`
	Email string `json:"nome"`
	Senha string `json:"password"`
}
