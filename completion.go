package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/schollz/progressbar/v3"
)

// Função para gerar o texto usando o modelo de linguagem
func generateCompletion(model, prompt, systemInstructionsClipboard string) (string, error) {
	payload := map[string]string{
		"model":  model,
		"prompt": prompt,
		"system": systemInstructionsClipboard,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("erro no empacotamento do JSON: %w", err)
	}

	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("erro ao fazer requisição: %w", err)
	}
	defer resp.Body.Close()

	// fmt.Println("Response Content-Type:", resp.Header.Get("Content-Type"))

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("requisição a API do llama falhou: %s", resp.Status)
	}

	// Cria a barra de progresso
	bar := progressbar.DefaultBytes(
		-1, // Tamanho total desconhecido (-1), a barra irá se ajustar automaticamente
		"Processando...",
	)

	scanner := bufio.NewScanner(io.TeeReader(resp.Body, bar)) // Conecta o leitor ao progresso
	var completionBuilder strings.Builder

	for scanner.Scan() {
		line := scanner.Text()

		var response struct {
			Completion string `json:"response"`
			Done       bool   `json:"done"`
		}

		if err := json.Unmarshal([]byte(line), &response); err != nil {
			return "", fmt.Errorf("erro ao desempacotar a resposta JSON: %w", err)
		}

		completionBuilder.WriteString(response.Completion)

		if response.Done {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("erro no escaneamento da resposta: %w", err)
	}

	return completionBuilder.String(), nil
}
