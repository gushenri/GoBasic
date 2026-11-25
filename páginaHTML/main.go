package main

import (
	"html/template" // Monta páginas HTML com dados do Go e protege a inserção de texto.
	"log"           // Mostra mensagens e erros no terminal.
	"net/http"      // Fornece as ferramentas para criar um servidor web.
)

const htmlPagina = `<!DOCTYPE html>
<html lang="pt-BR">
<head> 
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<title>Minha lista de tarefas</title>
</head>
<body>
	<main>
		<h1>Minha lista de tarefas</h1>
		<p>Estas são as tarefas que estou aprendendo a mostrar com Go:</p>
		<ul>
			{{range . }}
				<li>{{.}}</li>
			{{else}}
				<li>Nenhuma tarefa por enquanto.</li>
			{{end}}
		</ul>
	</main>
</body>
</html>`

// agora precisamos de uma variavel para entender o html
var modeloPagina = template.Must(template.New("inicio").Parse(htmlPagina))

// Esta função é chamada a cada requisição que corresponde à rota inicial.
// w permite montar a resposta; r contém os dados do pedido do navegador.
// *http.Request é um ponteiro: permite acessar a requisição sem copiá-la inteira.
func paginaInicial(w http.ResponseWriter, r *http.Request) {
	// []string é um slice: uma sequência de textos. Cada texto é uma tarefa.
	// Por enquanto, a lista é recriada a cada pedido; ainda não há cadastro.

	tarefas := []string{
		"Aprender a criar servidor em Go",
		"Mostrar uma página HTML",
		"Exibir tarefas usando um template",
	}

	//Agora o navegador deve interpretar a resposta como HTML, em vez de texto puro.
	w.Header().Set("Content-type", "text/html; charset=utf-8")

	// Execute combina o modelo com o slice e escreve o HTML resultante em w.
	// No modelo, {{range .}} percorre a lista recebida.
	// Dentro do range, {{.}} representa a tarefa atual; {{end}} fecha o bloco.
	// {{else}} mostra uma mensagem quando a lista está vazia.
	err := modeloPagina.Execute(w, tarefas)
	if err != nil {
		log.Printf("Erro ao enviar a resposta: %v", err)
	}

}

func main() {
	http.HandleFunc("GET /{$}", paginaInicial)

	endereco := "127.0.0.1:8080"
	log.Printf("Abra http://%s no navegador. Para parar, pressione Ctrl+C.", endereco)

	err := http.ListenAndServe(endereco, nil)
	// nil significa ausência de erro; != significa "diferente de".
	if err != nil {
		log.Fatal(err) // Mostra o erro e encerra o programa.
	}
}
