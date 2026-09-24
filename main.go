package main

import (
	"fmt"

	"github.com/goczangabor24/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	err = cfg.SetUser("Gabor")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	updatedCfg, err := config.Read()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	fmt.Println(updatedCfg)
}
