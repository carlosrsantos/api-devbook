package repositories

import (
	"api/src/models"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (p Publicacoes) Criar(publicacao models.Publicacao) (uint64, error) {
	statement, erro := p.db.Prepare(
		"insert into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (p Publicacoes) BuscarPorID(publicacaoID uint64) (models.Publicacao, error) {
	linha, erro := p.db.Query(`
		select p.*, u.nick from 
		publicacoes p inner join usuarios u
		on u.id = p.autor_id where p.id = ?`,
		publicacaoID,
	)
	if erro != nil {
		return models.Publicacao{}, erro
	}
	defer linha.Close()

	var publicacao models.Publicacao

	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick,
		); erro != nil {
			return models.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

func (p Publicacoes) Buscar(usuarioID uint64) ([]models.Publicacao, error) {
	linhas, erro := p.db.Query(`
		select distinct p.*, u.nick from publicacoes p 
		inner join usuarios u on u.id = p.autor_id
		inner join seguidores s on p.autor_id = s.usuario_id
		where u.id = ? or s.seguidor_id =?
		order by 1 desc;`,
		usuarioID, usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao

	for linhas.Next() {
		var publicacao models.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (p Publicacoes) Atualizar(publicacaoID uint64, publicacao models.Publicacao) error {
	statement, erro := p.db.Prepare("update publicacoes set titulo = ?, conteudo = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); erro != nil {
		return erro
	}

	return nil
}

func (p Publicacoes) Deletar(publicacaoID uint64) error {
	statement, erro := p.db.Prepare("delete from publicacoes where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}

func (p Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]models.Publicacao, error) {
	linhas, erro := p.db.Query(`
		select p.*, u.nick from publicacoes p 
		inner join usuarios u on u.id = p.autor_id
		where p.autor_id = ?`,
		usuarioID,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao

	for linhas.Next() {
		var publicacao models.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}

func (p Publicacoes) Curtir(publicacaoID uint64) error {
	statement, erro := p.db.Prepare("update publicacoes set curtidas = curtidas + 1 where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil

}
func (p Publicacoes) Descurtir(publicacaoID uint64) error {
	statement, erro := p.db.Prepare(`
		update publicacoes set curtidas = 
		case when curtidas > 0 then curtidas - 1
		else 0 end
		where id = ?
	`)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil

}
