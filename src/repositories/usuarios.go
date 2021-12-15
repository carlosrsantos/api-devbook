package repositorios

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

//repositorio de usuarios
type usuarios struct {
	db *sql.DB
}

//cria um repositorio de usuarios
func NovoRepositorioDeUsuario(db *sql.DB) *usuarios {
	return &usuarios{db}
}

func (u usuarios) Criar(usuario models.Usuario) (uint64, error) {
	statement, erro := u.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values(?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (u usuarios) Buscar(nomeOuNick string) ([]models.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%

	linhas, erro := u.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil

}

func (u usuarios) BuscarPorID(ID uint64) (models.Usuario, error) {
	linhas, erro := u.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = ?",
		ID,
	)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario models.Usuario
	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return models.Usuario{}, erro
		}
	}
	return usuario, nil
}

func (u usuarios) Atualizar(ID uint64, usuario models.Usuario) error {
	statement, erro := u.db.Prepare(
		"updte usuarios into set nome = ?, nick = ?, email = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}

	return nil
}
