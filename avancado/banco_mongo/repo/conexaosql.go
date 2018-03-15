package repo

import (
	"github.com/jmoiron/sqlx"
	/*
		github.com/go-sql-driver/mysql não é usado diretamente pela aplicação
	*/
	_ "github.com/go-sql-driver/mysql"
)

//Db armazena a conexão com o banco de dados
var Db *sqlx.DB

//AbreConexaoComBancoDeDadosSQL funcao que abre a conexao com o banco MYSQL
func AbreConexaoComBancoDeDadosSQL() (err error) {
	err = nil
	db, err = sqlx.Open("mysql", "jaburunx:mateusPrestes1@tcp(jaburunx.mysql.dbass.com.br:3306)/jaburunx?parseTime=true")
	if err != nil {
		return
	}
	err = Db.Ping()
	if err != nil {
		return
	}
	return
}
