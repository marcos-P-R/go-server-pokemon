package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Conexao() string {
	res, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	resData, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(resData)
}
