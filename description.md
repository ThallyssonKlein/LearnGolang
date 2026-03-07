Legal — como você já programa em **TypeScript, Java e Python**, a ideia é usar **projetos pequenos** que mostrem **o que é diferente no Go** (concorrência, interfaces implícitas, composição, tooling etc.).

Vou montar um **caminho progressivo**. Cada projeto pratica **um conceito principal da linguagem**.

Mas vamos fazer isso no formato de estudo: você implementa cada projeto e a gente discute as decisões depois.

---

# Plano de estudos de Go (baseado em micro-projetos)

## 1️⃣ CLI de conversão de temperatura

**Objetivo:** sintaxe básica

Projeto:

* Programa CLI que recebe um valor em Celsius e converte para Fahrenheit.

Conceitos de Go:

* `package main`
* `func main()`
* variáveis e tipos
* `fmt`
* compilação com `go build`

Extra:

* usar `os.Args`

💡 Diferença importante do Go: **binário único compilado**

---

## 2️⃣ CLI To-Do list simples

**Objetivo:** slices e structs

Projeto:

* CLI para:

  * adicionar tarefa
  * listar tarefas

Conceitos:

* `struct`
* `slice`
* métodos em structs

Exemplo de domínio:

```
Task
- id
- title
- done
```

---

## 3️⃣ To-Do persistido em JSON

**Objetivo:** serialização

Projeto:

* salvar tarefas em arquivo `tasks.json`

Conceitos:

* `encoding/json`
* tags em struct

Exemplo:

```
type Task struct {
    ID int `json:"id"`
}
```

---

## 4️⃣ Servidor HTTP simples

**Objetivo:** HTTP padrão do Go

Projeto:

* API REST simples de tasks

Rotas:

```
GET /tasks
POST /tasks
```

Conceitos:

* `net/http`
* handlers
* `http.ListenAndServe`

💡 Aqui você aprende algo muito Go:
**não precisa de framework**

---

## 5️⃣ Middleware manual

**Objetivo:** composição de funções

Projeto:

* adicionar middleware de log no servidor

Conceitos:

* funções como first class
* wrappers de handler

Exemplo conceitual:

```
logMiddleware(handler)
```

Isso mostra como Go faz middleware **sem framework**.

---

## 6️⃣ Worker concorrente

**Objetivo:** goroutines

Projeto:

* programa que baixa várias URLs em paralelo

Conceitos:

* `goroutine`
* `sync.WaitGroup`

Exemplo de comportamento:

```
urls := []string{}
baixar todas ao mesmo tempo
```

💡 Aqui começa a **principal feature da linguagem**.

---

## 7️⃣ Worker pool

**Objetivo:** channels

Projeto:

* processador de jobs concorrentes

Exemplo:

```
jobs -> channel -> workers -> results
```

Conceitos:

* `chan`
* `range chan`
* worker pool pattern

Esse é **um padrão muito usado em Go real**.

---

## 8️⃣ Sistema de cache em memória

**Objetivo:** maps + mutex

Projeto:

* cache de respostas HTTP

Conceitos:

* `map`
* `sync.Mutex`
* `sync.RWMutex`

Aqui você aprende **concorrência segura**.

---

## 9️⃣ Interface implícita

**Objetivo:** interfaces Go

Projeto:

* sistema de armazenamento de tasks

Criar interface:

```
TaskRepository
```

Implementações:

```
MemoryRepository
FileRepository
```

Conceitos:

* interfaces implícitas
* polimorfismo Go

💡 Muito diferente de Java/TS.

---

## 🔟 CLI com subcomandos

**Objetivo:** estrutura de projetos Go

Projeto:

* CLI estilo `git`

Exemplo:

```
task add
task list
task done
```

Conceitos:

* organização de pacotes
* múltiplos arquivos
* módulos Go (`go mod`)

---

# Projetos intermediários

Depois disso:

### 11️⃣ Rate limiter

Conceitos:

* `time`
* `ticker`
* channels

---

### 12️⃣ Web scraper concorrente

Conceitos:

* goroutines
* channels
* pipeline pattern

---

### 13️⃣ Chat TCP simples

Conceitos:

* `net`
* sockets
* goroutines por conexão

---

### 14️⃣ Sistema de pub/sub

Conceitos:

* channels
* fan-out
* event bus

---

# Projeto final sugerido

### 15️⃣ Mini sistema de filas (tipo SQS)

Funcionalidades:

```
enqueue
dequeue
ack
retry
```

Conceitos:

* channels
* concorrência
* persistência
* arquitetura Go

---

# Uma dica importante vindo de outras linguagens

Quando aprender Go, **não tente escrever Java em Go**.

Coisas importantes de internalizar:

* composição > herança
* interfaces pequenas
* erro retornado (`error`)
* simplicidade > abstração

---

✅ Se quiser, eu também posso te mostrar:

* **o roadmap REAL de Go que empresas esperam**
* **5 projetos que simulam código de produção**
* **os erros que devs de Java/Node mais cometem em Go**

Isso acelera MUITO o aprendizado.
