package Dingo

import (
	"fmt"
	"os"

	"github.com/dinever/dingo/app/handler"
	"github.com/dinever/dingo/app/model"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func Init(dbPath, privKey, pubKey string) {
	model.InitializeKey(privKey, pubKey)
	if err := model.Initialize(dbPath, fileExists(dbPath)); err != nil {
		err = fmt.Errorf("failed to intialize db: %v", err)
		panic(err)
	}
	fmt.Printf("Database is used at %s\n", dbPath)
}

func Run(portNumber string) {
	app := handler.Initialize()
	fmt.Printf("Application Started on port %s\n", portNumber)
	app.Run(":" + portNumber)
}
