## IAmigo

Um CLI para interagir com a IA da Meta (Ollama), funcionando como assistente para tarefas variadas.


#### Pré-requisitos

- Ollama em execução localmente: O IAmigo depende do Ollama para funcionar. Certifique-se de ter o Ollama instalado e em execução na sua máquina. Veja as instruções de instalação em [https://ollama.ai/]().
- Golang (opcional): O IAmigo é escrito em Go, então você precisa tê-lo instalado em seu sistema caso queira compilá-lo. Você pode baixá-lo em [https://go.dev/dl/](). Caso preferir, use as versões pré-compiladas na pasta [dist](dist).


#### Configuração

O IAmigo usa  o arquivo [config.csv](config.csv) para armazenar suas configurações de Perfil, Modelo e Instruções de Sistema. O usuário pode editar ou adicionar novas configurações conforme sua necessidade. O Modelo deve ter sido instalado no Ollama previamente: ``ollama pull <modelo>``


#### Build e Uso

1) Baixe o código-fonte

2) Compile o código (opcional), você pode usar as versões pré-compiladas da pasta [dist](dist):

```
go build iamigo.go completion.go config.go
chmod +x iamigo.go
```

3) Uso

```
./iamigo <perfil> <pergunta> <usar_clipboard>
```

4) Exemplo:
   Substitua `<perfil>` pelo nome do perfil que você deseja usar (ex: "python-qa", "react-native").
   Substitua `<pergunta>` pela sua pergunta sobre desenvolvimento ou testes de software.
   Substitua `<usar_clipboard` por 0 (zero) para não usar ou por 1 (um) para usar o conteúdo do clipboard como contexto adicional da pergunta. Esse parâmetro é opcional.

```
./iamigo python-qa "Como escrever um teste de performance com Locust?"

./iamigo react-native "Como adicionar a funcionalidade de buscar produtos dado esse contexto?" 1
```
