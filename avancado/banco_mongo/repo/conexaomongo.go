package repo

import (
	"gopkg.in/mgo.v2"
)

//SessaoMongo armazena a sessao de conexao com o Mongo
var SessaoMongo *mgo.Session

//AbreSessaoComMongo faz a conexao com o Mongo
func AbreSessaoComMongo() (err error) {
	err = nil
	SessaoMongo, err = mgo.Dial("mongodb://alunocursodego:Gopher2012@ds040014,mlab.com:40017/cursodego?authSource=cursodego")
	return
}
