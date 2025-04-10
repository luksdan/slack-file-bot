package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar .env: %v", err)
	}

	token := os.Getenv("SLACK_BOT_TOKEN")
	channelID := os.Getenv("CHANNEL_ID")
	filePath := "APENAS FALE.pdf"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("Erro ao obter informações do arquivo: %v", err)
	}

	if fileInfo.Size() == 0 {
		log.Fatal("Arquivo está vazio.")
	}

	api := slack.New(token)

	params := slack.UploadFileV2Parameters{
		Filename: fileInfo.Name(),
		Reader:   file,
		Channel:  channelID,
		FileSize: int(fileInfo.Size()),
	}

	result, err := api.UploadFileV2(params)
	if err != nil {
		log.Fatalf("Erro ao enviar arquivo: %v", err)
	}

	fmt.Printf("Arquivo enviado com sucesso!\nID: %s\nTítulo: %s\n", result.ID, result.Title)
}
