package controllers

import "net/http"

// CriarUsuario insere um usuario na base de dados
func CriarUsuario(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Criando Usuário!"))
}

// BuscarUsuarios busca todos os usuarios na base de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Buscando todos os Usuário!"))
}

// BuscarUsuario busca um unico usuario na base de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Buscando um unico Usuário!"))
}

// AtualizarUsuario atualiza um usuario na base de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Atualizando Usuário!"))
}

// DeletarUsuario deleta um usuario na base de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Deletando Usuário!"))
}
