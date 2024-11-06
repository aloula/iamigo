package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// Struct para armazenar as configurações
type Config struct {
	Modelos map[string]string
}

// Função para carregar as configurações do arquivo CSV
func LoadConfig(filename string) (*Config, error) {
	// Abre o arquivo CSV
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir o arquivo de configuração: %w", err)
	}
	defer file.Close()

	// Cria um novo leitor CSV com suporte a diferentes fins de linha
	reader := csv.NewReader(bufio.NewReader(file))
	reader.FieldsPerRecord = -1 // Permite que o leitor infira o número de campos

	// Lê os dados do arquivo
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o arquivo de configuração: %w", err)
	}

	// Verifica se há dados suficientes no arquivo
	if len(records) < 2 {
		return nil, fmt.Errorf("o arquivo de configuração deve ter pelo menos duas linhas")
	}

	// Cria o mapa de modelos
	config := &Config{
		Modelos: make(map[string]string),
	}

	// Começa a partir da segunda linha (índice 1) para pular o cabeçalho
	for _, row := range records[1:] {
		if len(row) >= 3 { // Garante que temos 3 colunas
			config.Modelos[strings.TrimSpace(row[0])] = strings.TrimSpace(row[1]) + " " + strings.TrimSpace(row[2])
		}
	}

	return config, nil
}
