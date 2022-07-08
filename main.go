package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	swissknife "github.com/Sagleft/swiss-knife"
)

func newSolution() solution {
	return solution{}
}

func (app *solution) loadConfig() error {
	if _, err := os.Stat(configJSONPath); os.IsNotExist(err) {
		return errors.New("failed to find config json")
	}

	jsonBytes, err := ioutil.ReadFile(configJSONPath)
	if err != nil {
		return errors.New("failed to read config file: " + err.Error())
	}

	err = json.Unmarshal(jsonBytes, &app.Config)
	if err != nil {
		return errors.New("failed to load config: " + err.Error())
	}
	return nil
}

func (app *solution) exchangeConnect() error {
	// TODO
	return nil
}

func (app *solution) loadPrices() error {
	// TODO
	return nil
}

func (app *solution) do() error {
	// TODO
	return nil
}

func main() {
	app := newSolution()

	err := swissknife.CheckErrors(
		app.loadConfig,
		app.exchangeConnect,
		app.loadPrices,
		app.do,
	)
	if err != nil {
		log.Fatalln(err)
	}
}
