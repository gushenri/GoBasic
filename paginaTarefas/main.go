package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"sync"
)

var modeloPagina = template.Must(template.ParseFiles("index.html"))

var tarefasMu sync.Mutex

var tarefas = []string{
	"Aprender a criar um servidor em Go",
	"Mostrar uma página HTML",
	"Exibir tarefas usando um template",
}

// O servidor pode atender vários pedidos ao mesmo tempo.
// O mutex impede que dois pedidos leiam e alterem a lista simultaneamente.

func paginaInicial(w http.ResponseWriter, r *http.Request) {

	tarefasMu.Lock()
	lista := append([]string(nil), tarefas...)
	tarefasMu.Unlock()
	// []string é um slice: uma sequência de textos. Cada texto é uma tarefa.
	// Copiamos a lista protegida para montar a página sem bloquear outros pedidos.

	// Agora o navegador deve interpretar a resposta como HTML, em vez de texto puro.
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Execute combina o modelo com o slice e escreve o HTML resultante em w.
	// No modelo, {{range .}} percorre a lista recebida.
	// Dentro do range, {{.}} representa a tarefa atual; {{end}} fecha o bloco.
	// {{else}} mostra uma mensagem quando a lista está vazia.
	err := modeloPagina.Execute(w, lista)
	if err != nil {
		log.Printf("Erro ao enviar a resposta: %v", err)
	}
}

func adicionarTarefa(w http.ResponseWriter, r *http.Request) {
	// adicionarTarefa recebe os dados enviados pelo formulário.

	// ParseForm interpreta os campos enviados pelo navegador.
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Não foi possível ler o formulário.", http.StatusBadRequest)
		return // Encerra esta função sem cadastrar a tarefa.
	}

	tarefa := strings.TrimSpace(r.PostForm.Get("tarefa"))
	// PostForm.Get busca o campo pelo name; TrimSpace remove espaços nas pontas.
	if tarefa == "" {
		http.Error(w, "Digite uma tarefa antes de add.", http.StatusBadRequest)
		return
	}

	// append devolve um slice com o novo item; guardamos esse resultado na lista.

	tarefasMu.Lock()
	tarefas = append(tarefas, tarefa)
	tarefasMu.Unlock()

	// Após o POST, o navegador faz um GET da página inicial.
	// Assim, atualizar a página não envia o formulário novamente.
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func main() {
	http.HandleFunc("GET /{$}", paginaInicial)
	http.HandleFunc("POST /tarefas", adicionarTarefa)

	endereco := "127.0.0.1:8080"
	log.Printf("Abra http://%s no navegador. Para parar, pressione Ctrl+C.", endereco)

	// ListenAndServe fica esperando requisições e chama a função correspondente.
	// nil usa o conjunto de rotas registrado acima com HandleFunc.
	// Se não conseguir iniciar ou continuar, devolve um erro.
	err := http.ListenAndServe(endereco, nil)
	// nil significa ausência de erro; != significa "diferente de".
	if err != nil {
		log.Fatal(err) // Mostra o erro e encerra o programa.
	}
}
