Para dominar a concorrência em Go, o segredo é transitar da simples execução de funções em background para a orquestração de fluxos de dados complexos e gerenciamento de ciclo de vida.

Como você já possui uma base sólida em sistemas distribuídos, este plano foca em padrões arquiteturais e segurança de memória, evitando o básico e indo direto para a implementação de sistemas resilientes.

---

## Fase 1: Sincronização e Primitivas
Antes de usar canais para tudo, é vital entender quando a memória compartilhada ainda é a ferramenta certa.

* **Mini-Algoritmo: Cache In-Memory Thread-Safe.**
    * **O Desafio:** Implemente um `Map` genérico que suporte leituras concorrentes massivas e escritas ocasionais.
    * **Foco:** Use `sync.RWMutex`. Compare a performance com `sync.Map` usando benchmarks (`go test -bench`).
* **Exercício: Singleton Pattern Concorrente.**
    * **O Desafio:** Garanta que a inicialização de um recurso pesado (ex: conexão com banco) ocorra exatamente uma vez, mesmo com 1000 goroutines tentando acessá-lo simultaneamente.
    * **Foco:** `sync.Once`.

---

## Fase 2: Orquestração com Channels
Aqui o foco é o mantra: *"Não comunique compartilhando memória; compartilhe memória se comunicando"*.

* **Padrão: Fan-In / Fan-Out.**
    * **O Desafio:** Crie um processador de imagens (ou strings) onde uma goroutine distribui tarefas para 10 workers (Fan-Out) e o resultado de todos é consolidado em um único canal de saída (Fan-In).
    * **Foco:** Buffered vs Unbuffered channels e o fechamento correto de canais para evitar deadlocks.
* **Mini-Projeto: Crawler de Diretórios.**
    * **O Desafio:** Um programa que percorre o sistema de arquivos recursivamente, calcula o hash MD5 de cada arquivo e imprime o total.
    * **Foco:** Uso de `sync.WaitGroup` para esperar a conclusão de todas as goroutines recursivas.



---

## Fase 3: Controle de Ciclo de Vida e Cancelamento
Em sistemas de produção, o maior risco é o "leak" de goroutines que nunca terminam.

* **Mini-Algoritmo: Timeout de Requisição Simulado.**
    * **O Desafio:** Implemente uma função que simula uma chamada de API demorada. Se ela não retornar em 200ms, o programa deve desistir e liberar os recursos.
    * **Foco:** `context.WithTimeout` e a cláusula `select`.
* **Exercício: Propagação de Cancelamento.**
    * **O Desafio:** Crie uma árvore de processos onde o cancelamento do "Pai" cancela automaticamente todos os "Filhos" e "Netos".
    * **Foco:** `context.WithCancel` passando o contexto via argumentos de função.

---

## Fase 4: Padrões Avançados e Resiliência
Aplicando concorrência em cenários de alta carga e instabilidade.

* **Mini-Projeto: Worker Pool com Rate Limiting.**
    * **O Desafio:** Um sistema que processa 10.000 jobs, mas não pode ultrapassar 50 requisições por segundo para não derrubar o serviço de destino.
    * **Foco:** `time.Ticker` ou bibliotecas de *token bucket* integradas ao seletor de canais.
* **Exercício: Pipeline de Dados.**
    * **O Desafio:** Implementar um pipeline `Estágio A -> Estágio B -> Estágio C`. Cada estágio é uma goroutine ligada por canais. Se um estágio falhar, o pipeline deve ser limpo graciosamente.
    * **Foco:** Padrão *Generator* e encadeamento de canais.

---

## Fase 5: Debugging e Performance
Saber escrever o código é metade do trabalho; a outra metade é garantir que ele não tem *race conditions*.

1.  **Race Detector:** Execute todos os seus projetos acima com a flag `go run -race`. Tente forçar uma race condition propositalmente para ver como o Go a detecta.
2.  **Profiling:** Use `net/http/pprof` para visualizar o dump de goroutines em um gráfico. É a melhor forma de encontrar goroutines "zumbis".
3.  **Benchmarking:** Compare a execução sequencial vs paralela (usando `runtime.GOMAXPROCS`) para entender o overhead de troca de contexto.

### Sugestão de Capstone Project
**Log Aggregator Distribuído (Local):**
Crie um programa que lê 5 arquivos de log simultaneamente (goroutines), filtra linhas que contenham "ERROR", envia para um buffer centralizado e, a cada 5 segundos ou 100 erros, faz um "flush" para um arquivo final ou banco de dados. Tudo isso respeitando cancelamento via `SIGINT` (Ctrl+C) usando `os/signal` e `context`.

Qual desses padrões ou primitivas você sente que seria mais desafiador implementar hoje na sua stack atual?