package main

import (
	"encoding/json"
	"os"

	"github.com/emejotaw/product-api/config"
)

func main() {

	config := config.New(".")
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(config)
}
