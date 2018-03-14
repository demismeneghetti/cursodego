package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/demismeneghetti/cursodego/avancado/web_post/model"
)

var orquestrador sync.WaitGroup

func main() {
	orquestrador.Add(4)
	go post("Demis Meneghetti", "demis@meneghetti.com", "S3cr3t")
	go post("Giovanni Meneghetti", "giovanni@meneghetti.com", "S3cr3t")
	go post("Antonella Meneghetti", "antonella@meneghetti.com", "S3cr3t")
	go post("Valentim Meneghetti", "valentim@meneghetti.com", "S3cr3t")
	orquestrador.Wait()
}

func post(nome, email, pass string) {

	fmt.Println("oi")

	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	usuario := model.Usuario{}
	usuario.Nome = nome
	usuario.Email = email
	usuario.Senha = pass

	conteudoEnviar, err := json.Marshal(usuario)
	if err != nil {
		fmt.Println("[main] Erro ao gerar o JSON do objeto usuario. Erro: ", err.Error())
		return
	}

	request, err := http.NewRequest("POST", "https://nbooks.herokuapp.com/api/accounts", bytes.NewBuffer(conteudoEnviar))
	if err != nil {
		fmt.Println("[main] Erro ao criar um request para o request bin. Erro: ", err.Error())
		return
	}
	// request.SetBasicAuth("fizz", "buzz")
	request.Header.Set("content-type", "application/json; charset=utf-8")
	resposta, err := cliente.Do(request)
	if err != nil {
		fmt.Println("[main] Erro ao chamar o servico do request bin. Erro: ", err.Error())
		return
	}
	defer resposta.Body.Close()

	corpo, err := ioutil.ReadAll(resposta.Body)
	if err != nil {
		fmt.Println("[main] Erro ao ler o conteudo retornado pelo request bin. Erro: ", err.Error())
		return
	}

	if resposta.StatusCode == 200 {
		fmt.Println("Passou >>>")
		fmt.Println(string(corpo))
	} else {
		fmt.Println("Fodeu... ")
		fmt.Println(string(corpo))
	}
}
