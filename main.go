package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	middleware "github.com/ueverson/ProcessingWorksheetGO/Middleware"
	"github.com/ueverson/ProcessingWorksheetGO/models"
)

type Configuration struct {
	UrlPlanilha string `json:"UrlPlanilha"`
}

func main() {

	config := ConfigRead()
	arq, err := os.Open(config.UrlPlanilha)
	middleware.Handler(err)

	leitor := csv.NewReader(bufio.NewReader(arq))
	leitor.Comma = ';'

	var model []models.Asset
	count := 0

	for {
		line, err := leitor.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			middleware.Handler(err)
		}

		if count > 0 {
			model = append(model, models.Asset{
				Ticker:          line[0],
				Name:            line[1],
				Type:            line[2],
				Sector:          line[3],
				Price:           middleware.ConvertFloat(line[4]),
				Dy:              line[5],
				Max:             middleware.ConvertFloat(line[6]),
				Min:             middleware.ConvertFloat(line[7]),
				DPA:             middleware.ConvertFloat(line[8]),
				FairPrice:       0,
				WithinFairPrice: false,
			})
		}

		count++
	}

	totalModel := 1
	for i, v := range model {
		calcFairPrice(&v)
		model[i] = v
		totalModel++
	}

	for _, s := range model {
		if s.WithinFairPrice {
			fmt.Println(s.Ticker, "com Price R$", s.Price, "onde deveria ser R$", fmt.Sprintf("%.2f", s.FairPrice), "DPA", s.DPA)
		}
	}

	fmt.Println("\nTotal: ", totalModel)
}

//Calcula o Preço Justo
func calcFairPrice(model *models.Asset) {
	if model.DPA != 0 {
		p := (model.DPA * 100) / 6
		b := p > model.Price

		(*model).FairPrice = p
		(*model).WithinFairPrice = b
	} else {
		(*model).FairPrice = 0
		(*model).WithinFairPrice = false
	}
}

//Faz a leitura da configuração do config
func ConfigRead() Configuration {
	file, err := os.Open("configs/config.json")
	middleware.Handler(err)
	defer file.Close()

	configBytes, err := ioutil.ReadAll(file)
	middleware.Handler(err)

	var config Configuration
	err = json.Unmarshal(configBytes, &config)
	middleware.Handler(err)

	return config
}
