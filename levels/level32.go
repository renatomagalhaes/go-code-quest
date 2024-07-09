package levels

import (
    "crypto/aes"
    "crypto/cipher"
		"crypto/rand"
    "crypto/sha256"
    "encoding/base64"
    "fmt"
    "go-code-quest/pkg/game"
		mrand "math/rand"
		"time"
    "io"
    "log"
    "net/http"

    "golang.org/x/crypto/bcrypt"
)

type Level32 struct{}

func StartLevel32() {
    game.StartLevel(Level32{})
}

func (l Level32) Start() {
    // Definir cores
    green := "\033[32m"
    yellow := "\033[33m"
    cyan := "\033[36m"
    reset := "\033[0m"

    separator := "==============================================="

    fmt.Println(separator)
    fmt.Println(string(cyan) + "Level 32: Segurança em Aplicações Go" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Desafio: Aprenda a proteger suas aplicações Go contra diversas ameaças de segurança" + string(reset))
    fmt.Println()

    // Gerenciamento de Senhas
    password := "my_secure_password"
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        log.Fatalf("Erro ao gerar hash da senha: %v", err)
    }
    fmt.Println(string(green) + "Senha: " + password + string(reset))
    fmt.Println(string(green) + "Senha Hasheada: " + string(hashedPassword) + string(reset))

    // Verificar Senha
    err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
    if err != nil {
        log.Fatalf("Erro ao comparar senha: %v", err)
    }
    fmt.Println(string(green) + "Senha verificada com sucesso!" + string(reset))

    // Criptografia
    secretMessage := "my secret message"
    key := sha256.Sum256([]byte("my_secret_key"))
    encryptedMessage, err := encrypt(secretMessage, key[:])
    if err != nil {
        log.Fatalf("Erro ao criptografar mensagem: %v", err)
    }
    fmt.Println(string(green) + "Mensagem Secreta: " + secretMessage + string(reset))
    fmt.Println(string(green) + "Mensagem Criptografada: " + encryptedMessage + string(reset))

    decryptedMessage, err := decrypt(encryptedMessage, key[:])
    if err != nil {
        log.Fatalf("Erro ao descriptografar mensagem: %v", err)
    }
    fmt.Println(string(green) + "Mensagem Descriptografada: " + decryptedMessage + string(reset))

    // Proteção contra Injeção de SQL
    http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        userID := r.URL.Query().Get("id")
        // Simulação de proteção contra injeção de SQL
        safeQuery := "SELECT * FROM users WHERE id = ?"
        fmt.Fprintf(w, "Query Segura: %s\n", safeQuery)
        fmt.Fprintf(w, "ID do Usuário: %s\n", userID)
    })

    fmt.Println(string(green) + "Servidor HTTP iniciado em http://localhost:8080" + string(reset))
    go http.ListenAndServe(":8080", nil)

    showSecuritySummary()
}

func encrypt(text string, key []byte) (string, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }
    b := base64.StdEncoding.EncodeToString([]byte(text))
    ciphertext := make([]byte, aes.BlockSize+len(b))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))

    return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func decrypt(cryptoText string, key []byte) (string, error) {
    ciphertext, err := base64.URLEncoding.DecodeString(cryptoText)
    if err != nil {
        return "", err
    }

    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }

    if len(ciphertext) < aes.BlockSize {
        return "", fmt.Errorf("ciphertext too short")
    }
    iv := ciphertext[:aes.BlockSize]
    ciphertext = ciphertext[aes.BlockSize:]

    stream := cipher.NewCFBDecrypter(block, iv)
    stream.XORKeyStream(ciphertext, ciphertext)

    data, err := base64.StdEncoding.DecodeString(string(ciphertext))
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func showSecuritySummary() {
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
    fmt.Println("package main")
    fmt.Println()
    fmt.Println("import (")
    fmt.Println("    \"crypto/aes\"")
    fmt.Println("    \"crypto/cipher\"")
    fmt.Println("    \"crypto/rand\"")
    fmt.Println("    \"crypto/sha256\"")
    fmt.Println("    \"encoding/base64\"")
    fmt.Println("    \"fmt\"")
    fmt.Println("    \"io\"")
    fmt.Println("    \"log\"")
    fmt.Println("    \"net/http\"")
    fmt.Println()
    fmt.Println("    \"golang.org/x/crypto/bcrypt\"")
    fmt.Println(")")
    fmt.Println()
    fmt.Println("func main() {")
    fmt.Println("    password := \"my_secure_password\"")
    fmt.Println("    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatalf(\"Erro ao gerar hash da senha: %v\", err)")
    fmt.Println("    }")
    fmt.Println("    fmt.Println(\"Senha: \" + password)")
    fmt.Println("    fmt.Println(\"Senha Hasheada: \" + string(hashedPassword))")
    fmt.Println()
    fmt.Println("    err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatalf(\"Erro ao comparar senha: %v\", err)")
    fmt.Println("    }")
    fmt.Println("    fmt.Println(\"Senha verificada com sucesso!\")")
    fmt.Println()
    fmt.Println("    secretMessage := \"my secret message\"")
    fmt.Println("    key := sha256.Sum256([]byte(\"my_secret_key\"))")
    fmt.Println("    encryptedMessage, err := encrypt(secretMessage, key[:])")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatalf(\"Erro ao criptografar mensagem: %v\", err)")
    fmt.Println("    }")
    fmt.Println("    fmt.Println(\"Mensagem Secreta: \" + secretMessage)")
    fmt.Println("    fmt.Println(\"Mensagem Criptografada: \" + encryptedMessage)")
    fmt.Println()
    fmt.Println("    decryptedMessage, err := decrypt(encryptedMessage, key[:])")
    fmt.Println("    if err != nil {")
    fmt.Println("        log.Fatalf(\"Erro ao descriptografar mensagem: %v\", err)")
    fmt.Println("    }")
    fmt.Println("    fmt.Println(\"Mensagem Descriptografada: \" + decryptedMessage)")
    fmt.Println()
    fmt.Println("    http.HandleFunc(\"/users\", func(w http.ResponseWriter, r *http.Request) {")
    fmt.Println("        userID := r.URL.Query().Get(\"id\")")
    fmt.Println("        safeQuery := \"SELECT * FROM users WHERE id = ?\"")
    fmt.Println("        fmt.Fprintf(w, \"Query Segura: %s\\n\", safeQuery)")
    fmt.Println("        fmt.Fprintf(w, \"ID do Usuário: %s\\n\", userID)")
    fmt.Println("    })")
    fmt.Println()
    fmt.Println("    fmt.Println(\"Servidor HTTP iniciado em http://localhost:8080\")")
    fmt.Println("    go http.ListenAndServe(\":8080\", nil)")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func encrypt(text string, key []byte) (string, error) {")
    fmt.Println("    block, err := aes.NewCipher(key)")
    fmt.Println("    if err != nil {")
    fmt.Println("        return \"\", err")
    fmt.Println("    }")
    fmt.Println("    b := base64.StdEncoding.EncodeToString([]byte(text))")
    fmt.Println("    ciphertext := make([]byte, aes.BlockSize+len(b))")
    fmt.Println("    iv := ciphertext[:aes.BlockSize]")
    fmt.Println("    if _, err := io.ReadFull(rand.Reader, iv); err != nil {")
    fmt.Println("        return \"\", err")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    stream := cipher.NewCFBEncrypter(block, iv)")
    fmt.Println("    stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))")
    fmt.Println()
    fmt.Println("    return base64.URLEncoding.EncodeToString(ciphertext), nil")
    fmt.Println("}")
    fmt.Println()
    fmt.Println("func decrypt(cryptoText string, key []byte) (string, error) {")
    fmt.Println("    ciphertext, err := base64.URLEncoding.DecodeString(cryptoText)")
    fmt.Println("    if err != nil {")
    fmt.Println("        return \"\", err")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    block, err := aes.NewCipher(key)")
    fmt.Println("    if err != nil {")
    fmt.Println("        return \"\", err")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    if len(ciphertext) < aes.BlockSize {")
    fmt.Println("        return \"\", fmt.Errorf(\"ciphertext too short\")")
    fmt.Println("    }")
    fmt.Println("    iv := ciphertext[:aes.BlockSize]")
    fmt.Println("    ciphertext = ciphertext[aes.BlockSize:]")
    fmt.Println()
    fmt.Println("    stream := cipher.NewCFBDecrypter(block, iv)")
    fmt.Println("    stream.XORKeyStream(ciphertext, ciphertext)")
    fmt.Println()
    fmt.Println("    data, err := base64.StdEncoding.DecodeString(string(ciphertext))")
    fmt.Println("    if err != nil {")
    fmt.Println("        return \"\", err")
    fmt.Println("    }")
    fmt.Println()
    fmt.Println("    return string(data), nil")
    fmt.Println("}")
    fmt.Println("```")

    fmt.Println()
    fmt.Println(separator)
    fmt.Println(string(cyan) + "Resumo do que foi aprendido" + string(reset))
    fmt.Println(separator)
    fmt.Println(string(yellow) + "Neste nível, você aprendeu sobre o gerenciamento de senhas, proteção contra injeção de SQL e XSS, e criptografia de dados em Go." + string(reset))
    fmt.Println()

    // Dicas e curiosidades
    dicas := []string{
        "Dica: Use `bcrypt` para armazenar senhas de forma segura.",
        "Curiosidade: `bcrypt` adiciona um valor de 'sal' para proteger contra ataques de rainbow table.",
        "Dica: Evite construir queries SQL diretamente concatenando strings. Use consultas preparadas.",
        "Curiosidade: AES (Advanced Encryption Standard) é um dos algoritmos de criptografia mais seguros e amplamente utilizados.",
        "Dica: Sempre use HTTPS para proteger a comunicação entre cliente e servidor.",
        "Curiosidade: O XSS (Cross-Site Scripting) permite que um invasor injete scripts maliciosos em páginas da web visualizadas por outros usuários.",
    }
    mrand.Seed(time.Now().UnixNano())
    dica := dicas[mrand.Intn(len(dicas))]

    // Exibir dica ou curiosidade
    fmt.Println(separator)
    fmt.Println(string(yellow) + dica + string(reset))
    fmt.Println(separator)

    // Parabenizar o jogador
    fmt.Println("\n" + string(green) + "Parabéns por concluir o trigésimo segundo nível!" + string(reset))
}
