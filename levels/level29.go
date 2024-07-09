package levels

import (
    "fmt"
    "go-code-quest/pkg/game"
    "os/exec"
)

type Level29 struct{}

func StartLevel29() {
    game.StartLevel(Level29{})
}

func (l Level29) Start() {
    // Definir cores
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 29: Integração com gRPC" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Crie um servidor e cliente gRPC para comunicação eficiente entre serviços" + string(reset))
    fmt.Println()

    // Iniciar o servidor gRPC
    fmt.Println(string(green) + "Iniciando servidor gRPC..." + string(reset))
    cmdServer := exec.Command("go", "run", "levels/level29/server/server.go")
    if err := cmdServer.Start(); err != nil {
        fmt.Println("Erro ao iniciar o servidor gRPC:", err)
        return
    }
    defer cmdServer.Process.Kill()

    // Iniciar o cliente gRPC
    fmt.Println(string(green) + "Iniciando cliente gRPC..." + string(reset))
    cmdClient := exec.Command("go", "run", "levels/level29/client/client.go")
    if err := cmdClient.Run(); err != nil {
        fmt.Println("Erro ao iniciar o cliente gRPC:", err)
        return
    }

    showGrpcSummary()
}

func showGrpcSummary() {
    green := "\033[32m"
    cyan := "\033[36m"
    yellow := "\033[33m"
    reset := "\033[0m"
    separator := "==============================================="

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Exemplo de solução em Go" + string(reset))
    fmt.Println(separator)
    fmt.Println("```go")
    fmt.Println("syntax = \"proto3\";")
    fmt.Println()
    fmt.Println("package greet;")
    fmt.Println()
    fmt.Println("service Greeter {")
    fmt.Println("  rpc SayHello (HelloRequest) returns (HelloReply);")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("message HelloRequest {")
    fmt.Println("  string name = 1;")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("message HelloReply {")
    fmt.Println("  string message = 1;")
    fmt.Println("}")
    fmt.Println("```")
    fmt.Println()
    fmt.Println(separator)
    fmt.Println("```go")
    fmt.Println("package main")
    fmt.Println()
    fmt.Println("import (")
    fmt.Println("    \"context\"")
    fmt.Println("    \"log\"")
    fmt.Println("    \"net\"")
    fmt.Println()
    fmt.Println("    \"go-code-quest/levels/level29/proto\"")
    fmt.Println()
    fmt.Println("    \"google.golang.org/grpc\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("type server struct {")
    fmt.Println("    proto.UnimplementedGreeterServer")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {")
    fmt.Println("    return &proto.HelloReply{Message: \"Hello, \" + req.Name}, nil")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    lis, err := net.Listen(\"tcp\", \":50051\")")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatalf(\"failed to listen: %v\", err)")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    s := grpc.NewServer()")
    fmt.Println("    proto.RegisterGreeterServer(s, &server{})")
    fmt.Println()
    fmt.Println("    log.Println(\"Server is running on port 50051...\")")
    fmt.Println("    if err := s.Serve(lis); err != nil {")
    fmt.Println("        log.Fatalf(\"failed to serve: %v\", err)")
    fmt.Println("    }")
    fmt.Println("}")
    fmt.Println("```")
    fmt.Println()
    fmt.Println(separator)
    fmt.Println("```go")
    fmt.Println("package main")
    fmt.Println()
    fmt.Println("import (")
    fmt.Println("    \"context\"")
    fmt.Println("    \"log\"")
    fmt.Println("    \"time\"")
    fmt.Println()
    fmt.Println("    \"go-code-quest/levels/level29/proto\"")
    fmt.Println()
    fmt.Println("    \"google.golang.org/grpc\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    conn, err := grpc.Dial(\"localhost:50051\", grpc.WithInsecure(), grpc.WithBlock())")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatalf(\"did not connect: %v\", err)")
    fmt.Println("    }")
    fmt.Println("    defer conn.Close()")
    fmt.Println()
    fmt.Println("    c := proto.NewGreeterClient(conn)")
    fmt.Println()
    fmt.Println("    ctx, cancel := context.WithTimeout(context.Background(), time.Second)")
    fmt.Println("    defer cancel()")
    fmt.Println()
    fmt.Println("    r, err := c.SayHello(ctx, &proto.HelloRequest{Name: \"World\"})")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatalf(\"could not greet: %v\", err)")
    fmt.Println("    }")
    fmt.Println("    log.Printf(\"Greeting: %s\", r.GetMessage())")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu a usar gRPC para comunicação eficiente entre serviços. Você criou um servidor e um cliente gRPC, definiu serviços usando Protocol Buffers, e testou chamadas de métodos gRPC." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use Protocol Buffers para definir seus serviços gRPC.",
        "Curiosidade: gRPC é amplamente usado em sistemas distribuídos devido à sua eficiência e desempenho.",
        "Dica: Use o comando `protoc` para gerar código a partir de arquivos .proto.",
        "Curiosidade: gRPC suporta várias linguagens de programação, permitindo a interoperabilidade entre serviços escritos em diferentes linguagens.",
        "Dica: Use `context` para gerenciar prazos e cancelamentos em chamadas gRPC.",
        "Curiosidade: gRPC é baseado no protocolo HTTP/2, que oferece melhor desempenho em comparação com HTTP/1.1.",
    }
    rand.Seed(time.Now().UnixNano())
    dica := dicas[rand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o vigésimo nono nível!" + string(reset))
}
