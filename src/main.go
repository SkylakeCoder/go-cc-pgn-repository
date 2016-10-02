package main

import (
	"cc"
	"flag"
	"path/filepath"
	"os"
	"log"
	"strings"
)

var chessRepositoryPath = flag.String("path", "", "useage: -path=xxx")

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
	flag.Parse()
	if *chessRepositoryPath == "" {
		flag.Usage()
	}
	cb := cc.ChessBoard{}
	cb.Init()
	files := getAllPGNFiles(*chessRepositoryPath)
	for _, path := range files {
		cb.Reset()
		cb.ParseRecord(path)
	}
}
