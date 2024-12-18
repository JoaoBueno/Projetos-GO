package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var (
	clientID     = "6cbb8a902d74a4f77cff80fe5dcddc609274478f"
	clientSecret = "1e3c430d66b567b08dd0190c70e485cc3a24e89411e9295c8c4e5a60438b"
	redirectURI  = "http://localhost:5000/callback"
	authURL      = "https://bling.com.br/Api/v3/oauth/authorize"
	tokenURL     = "https://bling.com.br/Api/v3/oauth/token"
	apiEndpoint  = "https://bling.com.br/Api/v3/produtos"
	state        = ""
	accessToken  = ""
	tokenFile    = "tokens.txt"
)

func main() {
	// Tenta ler o access_token do arquivo
	token, err := readToken()
	if err == nil && token != "" {
		accessToken = token
		fmt.Println("Usando o access_token existente.")
		makeAPICall()
	} else {
		fmt.Println("Nenhum token existente encontrado. Iniciando o fluxo de autenticação.")
		// Inicia o servidor HTTP em uma goroutine
		go func() {
			http.HandleFunc("/", handleIndex)
			http.HandleFunc("/callback", handleCallback)
			log.Fatal(http.ListenAndServe(":5000", nil))
		}()

		// Dá um tempo para o servidor iniciar
		time.Sleep(time.Second)

		// Abre o navegador para iniciar o fluxo de autenticação
		err = openBrowser("http://localhost:5000/")
		if err != nil {
			fmt.Println("Por favor, abra http://localhost:5000/ em seu navegador.")
		}

		// Aguarda a autenticação ser concluída
		for accessToken == "" {
			time.Sleep(time.Second)
		}

		// Faz a chamada à API
		makeAPICall()
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	state = generateState()
	authURL := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&response_type=code&state=%s",
		authURL, url.QueryEscape(clientID), url.QueryEscape(redirectURI), state)
	http.Redirect(w, r, authURL, http.StatusFound)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	// Verifica o estado para prevenção de CSRF
	receivedState := r.URL.Query().Get("state")
	if receivedState != state {
		http.Error(w, "Estado inválido", http.StatusBadRequest)
		return
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Código não encontrado", http.StatusBadRequest)
		return
	}

	// Troca o código pelo access_token
	err := exchangeCodeForToken(code)
	if err != nil {
		http.Error(w, "Erro ao obter o token de acesso", http.StatusInternalServerError)
		log.Println("Erro ao obter o token de acesso:", err)
		return
	}

	fmt.Fprintln(w, "Autenticação bem-sucedida! Você pode fechar esta janela.")
}

func exchangeCodeForToken(code string) error {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)

	clientCredentials := fmt.Sprintf("%s:%s", clientID, clientSecret)
	encodedCredentials := base64.StdEncoding.EncodeToString([]byte(clientCredentials))

	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+encodedCredentials)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Status da resposta do token:", resp.StatusCode)
	fmt.Println("Corpo da resposta do token:", string(bodyBytes))

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("falha ao obter o access_token, status code: %d", resp.StatusCode)
	}

	var tokenResponse map[string]interface{}
	err = json.Unmarshal(bodyBytes, &tokenResponse)
	if err != nil {
		return err
	}

	token, ok := tokenResponse["access_token"].(string)
	if !ok {
		return fmt.Errorf("access_token não encontrado na resposta")
	}

	accessToken = token

	// Salva o access_token em um arquivo
	err = saveToken(accessToken)
	if err != nil {
		log.Println("Erro ao salvar o token:", err)
	}

	return nil
}

func makeAPICall() {
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		log.Fatal("Erro ao criar a requisição:", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Erro ao fazer a chamada à API:", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Erro ao ler a resposta da API:", err)
	}

	var response interface{}
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		log.Fatal("Erro ao analisar a resposta da API:", err)
	}

	// Aqui aplicamos o unescape em strings específicas se necessário
	responseMap, ok := response.(map[string]interface{})
	if ok {
		descricaoCurta, exists := responseMap["descricaoCurta"].(string)
		if exists {
			// Desfaz a codificação HTML
			descricaoCurtaDecodificada := html.UnescapeString(descricaoCurta)
			responseMap["descricaoCurta"] = descricaoCurtaDecodificada
		}
	}

	responseFormatted, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Fatal("Erro ao formatar a resposta da API:", err)
	}

	fmt.Println(string(responseFormatted))

	// Salva em um arquivo
	err = os.WriteFile("saida.json", responseFormatted, 0644)
	if err != nil {
		log.Println("Erro ao escrever no arquivo:", err)
	}
}

// func makeAPICall() {
// 	req, err := http.NewRequest("GET", apiEndpoint, nil)
// 	if err != nil {
// 		log.Fatal("Erro ao criar a requisição:", err)
// 	}

// 	req.Header.Set("Authorization", "Bearer "+accessToken)

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatal("Erro ao fazer a chamada à API:", err)
// 	}
// 	defer resp.Body.Close()

// 	bodyBytes, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal("Erro ao ler a resposta da API:", err)
// 	}

// 	var response interface{}
// 	err = json.Unmarshal(bodyBytes, &response)
// 	if err != nil {
// 		log.Fatal("Erro ao analisar a resposta da API:", err)
// 	}

// 	responseFormatted, err := json.MarshalIndent(response, "", "    ")
// 	if err != nil {
// 		log.Fatal("Erro ao formatar a resposta da API:", err)
// 	}

// 	fmt.Println(string(responseFormatted))

// 	// Salva em um arquivo
// 	err = os.WriteFile("saida.json", responseFormatted, 0644)
// 	if err != nil {
// 		log.Println("Erro ao escrever no arquivo:", err)
// 	}
// }

func generateState() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func readToken() (string, error) {
	data, err := os.ReadFile(tokenFile)
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "access_token=") {
			return strings.TrimPrefix(line, "access_token="), nil
		}
	}
	return "", fmt.Errorf("access_token não encontrado no arquivo")
}

func saveToken(token string) error {
	data := fmt.Sprintf("access_token=%s\n", token)
	return os.WriteFile(tokenFile, []byte(data), 0644)
}

func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
		args = []string{url}
	}

	return exec.Command(cmd, args...).Start()
}
