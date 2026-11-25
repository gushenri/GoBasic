package main //diz que esse arquivo é o executavel

import (
	"fmt"      // excreve e le os textos
	"log"      // Mostra mensagens e erros no terminal.
	"net/http" // Fornece as ferramentas para criar um servidor web.
)

func main() {
	// HandleFunc associa um caminho a uma função que atende requisições.
	// O padrão GET /{$} atende apenas a página inicial; {$} marca o fim do caminho.
	// Uma requisição GET é o pedido que o navegador faz para buscar uma página.
	http.HandleFunc("GET /{$}", paginaInicial)

	endereco := "127.0.0.1:8080"

	// ListenAndServe fica esperando requisições e chama a função correspondente.
	// nil usa o conjunto de rotas registrado acima com HandleFunc.
	// Se não conseguir iniciar ou continuar, devolve um erro.
	log.Printf("Abra http://%s no navegador !!", endereco)

	err := http.ListenAndServe(endereco, nil)
	if err != nil {
		log.Fatal(err) // Mostra o erro e encerra o programa.
	}

}

// essa função é chamada a cada requisição
// w monta a resposta --> r vai conter os pedidos do navegador
// *http.request é um ponteiro, ela vai acessar a requisição sem copia-la
func paginaInicial(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// O Header informa que a resposta será em texto e comsuporte a acentos

	_, err := fmt.Fprintln(w, "Olá, Mundo! Meu servidor está funcionando !!!! ")
	// " _ " descarta e só guarda o erro em err
	// := cria a variavel e já coloca o valor nela
	//Fprintln envia a mensagem ao navegador e acrescenta uma quebra de linha.
	// O Fprintln não escreve no terminal e sim onde foi passado para ele escrever

	if err != nil {
		log.Printf("Erro ao enviar Resposta: %v", err)
	}

}
