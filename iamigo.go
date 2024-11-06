package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
)

// Define cores personalizadas para as fontes
var (
	azul     = color.New(color.FgBlue).SprintFunc()
	verde    = color.New(color.FgGreen).SprintFunc()
	vermelho = color.New(color.FgRed).SprintFunc()
	amarelo  = color.New(color.FgYellow).SprintFunc()
	cinza    = color.New(color.FgHiBlack).SprintFunc()
)

func main() {
	// Ascii Art
	myFigure := figure.NewColorFigure("IAmigo", "lean", "green", true)
	myFigure.Print()
	fmt.Println()

	// Verifica argumentos da linha de comando
	if len(os.Args) < 3 {
		fmt.Println(vermelho("Uso: iamigo <perfil> <pergunta> <usar_clipboard>"))
		fmt.Println(vermelho("Exemplo: ./iamigo python-qa 'Como usar o pytest?' 0"))
		os.Exit(1)
	}

	perfilEscolhido := os.Args[1]
	prompt := os.Args[2]

	// Obter o diretório do executável
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println(vermelho("Erro ao obter o caminho do executável:", err))
		os.Exit(1)
	}
	exeDir := filepath.Dir(exePath)

	// Caminho completo para o arquivo de configuração
	configPath := filepath.Join(exeDir, "config.csv")

	// Carrega as configurações do arquivo CSV
	config, err := LoadConfig(configPath)
	if err != nil {
		fmt.Println(vermelho("Erro ao carregar as configurações:", err))
		os.Exit(1)
	}

	// Verifica se o valor do <usar_clipboard> é "0" ou "1"
	var clipboardContent string
	var clipboardWithContext string

	if len(os.Args) > 3 { // Verifica se o terceiro argumento existe
		strClip := os.Args[3]

		if strClip == "1" {
			var err error
			clipboardContent, err = clipboard.ReadAll()
			clipboardWithContext = " Contexto adicional: \n\n" + clipboardContent + "\n"
			if err != nil {
				fmt.Println(vermelho("Erro ao ler a área de transferência:", err))
				os.Exit(1)
			}
		} else if strClip == "0" {
			clipboardWithContext = ""
		} else {
			fmt.Println(vermelho("Valores válidos para <usar_clipboard> são 0 ou 1:", strClip))
			os.Exit(1)
		}
	}

	// Obter o modelo e instruções da linguagem escolhida
	modelo, instrucoes, ok := obterModeloInstrucoes(config, perfilEscolhido)
	if !ok {
		fmt.Printf(vermelho("Perfil '%s' não encontrada no arquivo de configuração.\n"), perfilEscolhido)
		os.Exit(1)
	}

	systemInstructionsClipboard := instrucoes + clipboardWithContext
	fmt.Println(amarelo("- Instruções: ", systemInstructionsClipboard))

	// Chamando a função do módulo 'completion' para requisição ao Ollama
	completion, err := generateCompletion(modelo, prompt, systemInstructionsClipboard)
	if err != nil {
		fmt.Println(vermelho("Erro na geração da resposta:", err))
		os.Exit(1)
	}

	fmt.Println(verde("\n\n- Resposta do IAmigo:\n"), cinza(completion))
}

// Função auxiliar para buscar modelo e instruções por linguagem
func obterModeloInstrucoes(config *Config, perfil string) (string, string, bool) {
	perfil = strings.TrimSpace(perfil)
	fmt.Println(cinza("- Perfil: ", perfil))
	for lang, data := range config.Modelos {
		if strings.TrimSpace(lang) == perfil {
			parts := strings.SplitN(data, " ", 2)
			if len(parts) == 2 {
				fmt.Println(azul("- Modelo: ", parts[0]))
				return parts[0], parts[1], true
			}
		}
	}
	return "", "", false
}
