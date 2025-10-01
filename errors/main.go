package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)
/*
to run this application use this commands:
1. go build -o loader.exe main.go
2. .\loader.exe -url "https://example.com/file.txt" -output "file.txt"
*/

var (
	ErrInvalidURL        = errors.New("invalid URL")
	InvalidProtocolError = errors.New("invalid protocol")
	ErrConnectionFailed  = errors.New("connection failed")
	ErrDownloadFailed    = errors.New("download failed")
	ErrFileNotFound      = errors.New("file not found")
)

func downloadFile(fileURL, fileName string) error {
	return nil
}

func main() {
	url := flag.String("url", "", "URL файла для скачивания")
	output := flag.String("output", "", "Имя файла для сохранения")
	flag.Parse()

	if *url == "" || *output == "" {
		log.Fatal("Необходимо указать параметры -url и -output")
	}

	err := downloadFile(*url, *output)
	if err == nil {
		return
	}

	switch {
	case errors.Is(err, ErrInvalidURL):
		fmt.Println("Ошибка: недействительный URL. Введите корректный URL.")
	case errors.Is(err, InvalidProtocolError):
		fmt.Printf("Ошибка: неподдерживаемый протокол в URL (%s). Используйте http или https.\n", *url)
	case errors.Is(err, ErrConnectionFailed):
		fmt.Println("Ошибка подключения: не удалось установить соединение. Повторите попытку или проверьте настройки сети.")
	case errors.Is(err, ErrDownloadFailed):
		fmt.Println("Ошибка загрузки: проверьте доступность файла для скачивания.")
	case errors.Is(err, ErrFileNotFound):
		fmt.Println("Ошибка 404: файл не найден. Проверьте правильность URL или наличие файла на сервере.")
	default:
		fmt.Printf("Произошла непредвиденная ошибка: %v\n", err)
	}
}
