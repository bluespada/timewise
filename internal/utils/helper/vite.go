package helper

import (
	"log"
	"os"
)

type ViteMetadata struct{}

func GetViteMetadata(path string) []ViteMetadata {
	_, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
		return []ViteMetadata{}
	}
	return []ViteMetadata{}
}
