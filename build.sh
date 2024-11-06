#!/bin/bash

# Diretório de saída para os binários
OUTPUT_DIR="dist"

# Criar o diretório de saída, se não existir
mkdir -p "$OUTPUT_DIR"

# Função para construir o binário para uma plataforma específica
build_binary() {
  local GOOS="$1"
  local GOARCH="$2"
  local OUTPUT_NAME="$3"

  echo "Build para $GOOS/$GOARCH..."
  env GOOS="$GOOS" GOARCH="$GOARCH" go build -o "$OUTPUT_DIR/$OUTPUT_NAME" .
  if [ $? -ne 0 ]; then
    echo "Erro na build para $GOOS/$GOARCH"
    exit 1
  fi
}

# Construir para Windows (x86)
build_binary windows 386 "iamigo-windows-386.exe"

# Construir para Linux (x86)
build_binary linux 386 "iamigo-linux-386"

# Construir para Linux (ARM)
build_binary linux arm "iamigo-linux-arm"

# Construir para macOS (x86)
build_binary darwin amd64 "iamigo-darwin-amd64"

# Construir para macOS (ARM)
build_binary darwin arm64 "iamigo-darwin-arm64"

echo "Build concluída com sucesso!"
