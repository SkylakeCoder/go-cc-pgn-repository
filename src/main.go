package main

import (
	"flag"
	"path/filepath"
	"os"
	"log"
	"strings"
	"repository"
)

var chessRepositoryPath = flag.String("path", "", "usage: -path=xxx")
var dbPath = flag.String("dbpath", "", "usage: -dbpath=xxx")
var host = flag.String("host", "localhost", "usage: -host=xxx")
var port = flag.String("port", "8688", "usage: -port=xxx")

func getAllPGNFiles(path string) []string {
	result := []string {}
	err := filepath.Walk(path, func(path string, file os.FileInfo, err error) error {
		if file == nil {
			return err
		}
		if file.IsDir() {
			return nil
		}
		if strings.Contains(path, ".pgn") {
			result = append(result, path)
		}
		return nil
	})
	if err != nil {
		log.Fatalln("get files failed...path=>", path)
	}
	return result
}

func main() {
	repository.Init()

	flag.Parse()
	if (*dbPath != "") {
		repository.Load(*dbPath)
	} else {
		if *chessRepositoryPath == "" {
			flag.Usage()
		}
		cb := repository.ChessBoard{}
		cb.Init()
		files := getAllPGNFiles(*chessRepositoryPath)
		totalCount := len(files)
		index := 1
		for _, path := range files {
			log.Printf("current: %s\n", path)
			cb.Reset()
			cb.ParseRecord(path)
			log.Printf("[%d/%d]\n", index, totalCount)
			index++
		}
		repository.Save()
	}

	repository.StartServer(*host, *port)
}
