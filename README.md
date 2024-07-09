# Jogo CLI: "Go Code Quest"

O **Go Code Quest** é um jogo educativo baseado em linha de comando (CLI) projetado para ensinar habilidades básicas e avançadas de programação em Go. Através de diversos níveis, os jogadores são desafiados a resolver problemas e aprender conceitos fundamentais e avançados da linguagem Go. Cada nível aborda um tópico específico, com exemplos práticos e explicações, permitindo que os jogadores aprimorem suas habilidades de programação de forma interativa e divertida.

## Como Jogar

Para jogar o Go Code Quest, você precisa ter o Go instalado em seu sistema. Siga as instruções abaixo para iniciar o jogo:

1. Clone o repositório:
    ```sh
    git clone https://github.com/renatomagalhaes/go-code-quest.git
    cd go-code-quest
    ```

2. Compile o jogo:
    ```sh
    go build -o goquest cmd/goquest/main.go
    ```

3. Execute o jogo:
    ```sh
    ./goquest start <level>
    ```

    Substitua `<level>` pelo nível que você deseja iniciar, por exemplo:
    ```sh
    ./goquest start level1
    ```

### Níveis Disponíveis

Os seguintes níveis estão disponíveis:

1. **level1:** Variáveis e Tipos de Dados
2. **level2:** Estruturas de Controle
3. **level3:** Funções
4. **level4:** Estruturas e Métodos
5. **level5:** Arrays e Slices
6. **level6:** Mapas
7. **level7:** Concurrency (Concorrência)
8. **level8:** Channels
9. **level9:** Testes
10. **level10:** Pacotes
11. **level11:** Interfaces
12. **level12:** Errors (Erros)
13. **level13:** Leitura e Escrita de Arquivos
14. **level14:** JSON
15. **level15:** HTTP
16. **level16:** Context
17. **level17:** Templates
18. **level18:** Banco de Dados SQL
19. **level19:** Banco de Dados NoSQL
20. **level20:** Variáveis de Ambiente
21. **level21:** CLI (Command Line Interface)
22. **level22:** Ponteiros
23. **level23:** Reflection
24. **level24:** Go Modules
25. **level25:** Goroutines e WaitGroups
26. **level26:** WebSockets
27. **level27:** Middleware
28. **level28:** Logging
29. **level29:** gRPC
30. **level30:** Caching
31. **level31:** OAuth
32. **level32:** GraphQL
33. **level33:** JWT (JSON Web Tokens)
34. **level34:** Prometheus para Monitoramento
35. **level35:** Kubernetes
36. **level36:** Terraform
37. **level37:** Plugins Dinâmicos
38. **level38:** Integração com Serviços Externos
39. **level39:** Channels e Select
40. **level40:** Singleton
41. **level41:** Tarefas Assíncronas e WaitGroups
42. **level42:** Conexão Avançada com RabbitMQ



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
- **Nível 32: Segurança em Aplicações Go**
  - Práticas de segurança para desenvolver aplicações seguras em Go. O jogador aprenderá sobre o gerenciamento de senhas e autenticação, proteção contra injeção de SQL e XSS, e uso de pacotes para criptografia.
- **Nível 33: Trabalhando com Canais Avançados**
  - Uso avançado de canais para sincronização e comunicação entre goroutines. O jogador aprenderá sobre padrões de comunicação com canais, buffers e timeouts com canais, e casos de uso avançados.
- **Nível 34: Monitoramento e Logging em Go**
  - Implementar monitoramento e logging eficaz em aplicações Go. O jogador aprenderá sobre a integração com ferramentas de monitoramento, a coleta e o armazenamento de logs, e o uso de pacotes de logging.
- **Nível 35: Configurações de Aplicações com Viper**
  - Usar o pacote Viper para gerenciar configurações em aplicações Go. O jogador aprenderá a carregar configurações de arquivos, variáveis de ambiente e outros fontes, e a acessar esses valores no código.
- **Nível 36: Manipulação de JSON em Go**
  - Trabalhar com JSON em Go. O jogador aprenderá a serializar e desserializar dados JSON, lidar com estruturas aninhadas e usar tags de struct para customizar a conversão.
- **Nível 37: Utilizando WebSockets em Go**
  - Implementar WebSockets em aplicações Go, permitindo tanto o envio quanto o recebimento de mensagens. O jogador aprenderá a configurar um servidor WebSocket, lidar com conexões e mensagens, e usar a biblioteca gorilla/websocket.
- **Nível 38: Utilizando Plugins Dinâmicos em Go**
  - Criar e utilizar plugins dinâmicos em Go. O jogador aprenderá a compilar código como plugin, carregar o plugin em tempo de execução e usar as funcionalidades do plugin. ````go build -buildmode=plugin -o plugin38.so plugin38.go````
- **Nível 39: Utilizando Channels e Select para Comunicação Concorrente em Go**
  - Utilizar channels e a instrução select para comunicação concorrente em Go. O jogador aprenderá a criar channels, enviar e receber dados através deles, e usar a instrução select para múltiplas operações de comunicação.
- **Nível 40: Padrão de Projeto Singleton em Go**
  - Implementar o padrão de projeto Singleton em Go para garantir que uma classe tenha apenas uma instância e fornecer um ponto global de acesso a essa instância.
- **Nível 41: Trabalhando com Tarefas Assíncronas e Go Routines**
  - Gerenciar tarefas assíncronas usando Go Routines. O jogador aprenderá a iniciar Go Routines, sincronizá-las com WaitGroups, e gerenciar a concorrência de forma eficiente.
- **Nível 42: Conexão Avançada com RabbitMQ e Consumidores**
  - Implementar uma conexão avançada com RabbitMQ, incluindo a criação de consumidores, leitura de uma fila de alto movimento, e uso de ack e nack para gerenciamento de mensagens.

## Conclusão

O **Go Code Quest** oferece uma ampla gama de tópicos que cobrem desde conceitos básicos até avançados na linguagem Go. Ao completar os níveis, os jogadores terão uma compreensão sólida das melhores práticas de programação em Go, além de experiência prática com exemplos reais. Aproveite a jornada e bons estudos!