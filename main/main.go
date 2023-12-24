package main

import (
	"FTPS_pruebaGO/pkg"
	"log"
)

func main() {
	//err := pkg.Kardianos()
	//err := pkg.Webguerilla()
	err := pkg.ClienteFTPS()
	//err := pkg.DutchCoders()
	if err != nil {
		log.Fatal(err)
	}
}
