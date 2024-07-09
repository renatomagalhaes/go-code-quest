# Jogo CLI: "Go Code Quest"

````plaintext
go-code-quest/
├── cmd/
│   └── goquest/
│       └── main.go
├── levels/
│   ├── level1.go
│   ├── level2.go
│   ├── level3.go
│   ├── level4.go
│   ├── level5.go
│   ├── level6.go
│   ├── level7.go
│   ├── level8.go
│   ├── level9.go
│   ├── level10.go
│   ├── level11.go
│   ├── level12.go
│   ├── level13.go
│   ├── level14.go
│   ├── level15.go
│   ├── level16.go
│   ├── level17.go
│   ├── level18.go
│   ├── level19.go
│   ├── level20.go
│   ├── level21.go
│   ├── level22.go
│   ├── level23.go
│   ├── level24.go
│   ├── level25.go
│   ├── level26.go
│   ├── level27.go
│   ├── level28.go
│   ├── level29.go
│   ├── level30.go
│   ├── level31.go
├── pkg/
│   ├── game/
│   │   └── game.go
└── README.md
````

## Níveis Progressivos:
- **Nível 1: Variáveis e Tipos de Dados**
  - Variáveis e tipos básicos de dados.
- **Nível 2: Estruturas de Controle**
  - Estruturas de controle básicas em Golang, como if-else, switch, e for loops.
- **Nível 3: Funções**
  - Criação e uso de funções em Golang. O jogador aprenderá a definir, chamar e passar argumentos para funções.
- **Nível 4: Estruturas de Dados**
  - Manipulação de estruturas de dados em Golang, como arrays, slices e mapas. O jogador aprenderá a declarar, inicializar e manipular essas estruturas.
- **Nível 5: Estruturas e Métodos**
  - Criação e uso de estruturas (structs) e métodos em Golang. O jogador aprenderá a definir structs, criar instâncias e associar métodos a structs.
- **Nível 6: Interfaces**
  - Criação e uso de interfaces em Golang. O jogador aprenderá a definir interfaces, implementar métodos e usar interfaces para generalizar comportamentos.
- **Nível 7: Concurrency**
  - Criação e uso de goroutines e canais em Golang. O jogador aprenderá a definir goroutines para execução concorrente e a usar canais para comunicação entre goroutines.
- **Nível 8: Manipulação de Arquivos**
  - Leitura e escrita de arquivos em Golang. O jogador aprenderá a abrir, ler e escrever arquivos, bem como lidar com erros durante a manipulação de arquivos.
- **Nível 9: Testes**
  - Criação e execução de testes em Golang. O jogador aprenderá a escrever testes unitários usando o pacote testing.
- **Nível 10: Integração com Banco de Dados**
  - Conexão e manipulação de um banco de dados em Golang. O jogador aprenderá a conectar-se a um banco de dados, executar consultas SQL e manipular dados.
- **Nível 11: Integração com RabbitMQ**
  - Integração com RabbitMQ em Golang. O jogador aprenderá a se conectar a um broker RabbitMQ, enviar mensagens a uma fila e consumir mensagens da fila, além de aprender a usar ack e nack para mensagens.
- **Nível 12: Manipulação de JSON**
  - Manipulação de JSON em Golang. O jogador aprenderá a codificar e decodificar dados JSON.
- **Nível 13: Concorrência (wait groups)**
  - Concorrência avançada em Golang. O jogador aprenderá a usar goroutines, canais e grupos de espera (wait groups) para sincronização de tarefas concorrentes.
- **Nível 14: Paralelismo**
  - Paralelismo em Golang. O jogador aprenderá a usar goroutines em conjunto com a biblioteca sync/atomic para operações atômicas e runtime.GOMAXPROCS para controlar o número de threads do sistema.
- **Nível 15: Build e Execução de Programas**
  - Construir e executar programas em Golang. O jogador aprenderá a usar o comando go build para compilar programas Go e go run para executá-los diretamente.
- **Nível 16: Ponteiros**
  - Uso de ponteiros em Golang. O jogador aprenderá a declarar ponteiros, usar operadores de desreferenciação e trabalhar com funções que aceitam ponteiros como argumentos.
- **Nível 17: Uso de Templates SQL**
  - Uso de templates SQL em Golang. O jogador aprenderá a usar arquivos de template para criar consultas SQL dinâmicas e como utilizar esses templates no código Go, incluindo exemplos de consultas simples e consultas mais complexas com INNER JOIN.
- **Nível 18: Concorrência Avançada**
  - Uso avançado de concorrência em Golang. O jogador aprenderá a usar sync.WaitGroup, canais (chan), e o pacote context para gerenciar goroutines e comunicação entre elas de forma eficaz.
- **Nível 19: Integração com MongoDB**
  - Integração com o banco de dados NoSQL MongoDB em Golang. O jogador aprenderá a se conectar ao MongoDB, inserir, consultar e manipular documentos.
- **Nível 20: Criação do Arquivo .env**
  - Criar um arquivo .env no diretório raiz do seu projeto e adicionar variavéis de ambiente ex:
    - RABBITMQ_URI=amqp://guest:guest@localhost:5672/
    - QUEUE_NAME=test_queue
- **Nível 21: Uso de Templates para Geração de Código**
  - Usar templates em Go para gerar código dinamicamente. O jogador aprenderá a usar o pacote text/template para criar templates que podem ser preenchidos com dados e gerar arquivos de código Go.
- **Nível 22: Uso Avançado de Ponteiros**
  - Uso avançado de ponteiros em Golang. O jogador aprenderá a manipular ponteiros, passando o endereço de variáveis em vez de seus valores, e entenderá como isso pode ser eficiente e útil.
- **Nível 23: Integração com APIs REST usando HTTP**
  - Integrações com APIs REST em Golang usando o pacote net/http. O jogador aprenderá a fazer requisições GET, POST, PUT e DELETE, além de processar respostas JSON.
- **Nível 24: Uso de Contextos para Controle de Execução e Cancelamento**
  - Uso do pacote context em Golang para controlar o tempo de execução e cancelamento de goroutines. O jogador aprenderá a criar contextos com timeout e cancelamento, e a usá-los para controlar a execução de goroutines.
- **Nível 25: Uso de WebSockets para Comunicação em Tempo Real**
  - Usar WebSockets em Golang para permitir comunicação em tempo real entre um servidor e clientes. O jogador aprenderá a configurar um servidor WebSocket, enviar e receber mensagens em tempo real.
- **Nível 26: Testes Unitários e de Integração com Go**
  - Escrever testes unitários e de integração em Golang usando o pacote testing. O jogador aprenderá a escrever testes para funções individuais e a integrar testes que verificam o comportamento de múltiplos componentes.
- **Nível 27: Profiling e Otimização de Código**
  - Usar ferramentas de profiling para identificar gargalos de desempenho e otimizar o código. O jogador aprenderá a usar o pacote pprof para fazer profiling de CPU e memória, analisar os resultados e aplicar técnicas de otimização.
- **Nível 28: Implementação de Middlewares em Go**
  - Criar e usar middlewares para gerenciar requisições HTTP. O jogador aprenderá a criar middlewares básicos, usar middlewares para autenticação, logging e tratamento de erros, e integrar middlewares em um servidor HTTP.
- **Nível 29: Integração com gRPC**
  - Usar gRPC para comunicação eficiente entre serviços. O jogador aprenderá a criar um servidor e cliente gRPC, definir serviços usando Protocol Buffers, e testar chamadas de métodos gRPC.
- **Nível 30: Aplicação de Clean Architecture em Go**
  - Aplicar princípios de Clean Architecture em projetos Go. O jogador aprenderá a estruturar projetos seguindo a Clean Architecture, separando as camadas de domínio, aplicação e infraestrutura, e implementando casos de uso e interfaces.
- **Nível 31: Manipulação de Arquivos e Diretórios**
  - Manipular arquivos e diretórios no sistema de arquivos. O jogador aprenderá a ler e escrever arquivos, criar e remover diretórios, e usar os pacotes os e io/ioutil.
