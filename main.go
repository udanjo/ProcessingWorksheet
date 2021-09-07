package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/ueverson/ProcessingWorksheetGO/configs"
	"github.com/ueverson/ProcessingWorksheetGO/database"
	"github.com/ueverson/ProcessingWorksheetGO/middleware"
	"github.com/ueverson/ProcessingWorksheetGO/models"
)

func main() {
	database.StartDB()

	config := configs.Config()
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

	ResetDatabase()

	totalModel := 1
	totalFairPrice := 1
	for i, v := range model {
		calcFairPrice(&v)
		model[i] = v
		totalModel++
	}

	for _, s := range model {
		if s.WithinFairPrice {
			result := Create(s)
			fmt.Println(s.Ticker, "com Price R$", s.Price, "onde deveria ser R$", s.FairPrice, " Gravou ?", result)
			totalFairPrice++
		}
	}

	fmt.Println("\nTotal: ", totalModel)
	fmt.Println("Total com valor justo: ", totalFairPrice)
}

//Calcula o Preço Justo
func calcFairPrice(model *models.Asset) {
	if model.DPA != 0 {
		s := fmt.Sprintf("%.2f", (model.DPA*100)/6)
		p := middleware.ConvertFloat(s)
		b := p > model.Price

		(*model).FairPrice = p
		(*model).WithinFairPrice = b
	} else {
		(*model).FairPrice = 0
		(*model).WithinFairPrice = false
	}
}

func Create(m models.Asset) bool {
	db := database.GetDatabase()

	err := db.Create(&m).Error
	if err != nil {
		middleware.Handler(err)
		return false
	}

	return true
}

func ResetDatabase() bool {
	db := database.GetDatabase()

	err := db.Exec("DELETE FROM assets").Error

	if err != nil {
		middleware.Handler(err)
		return false
	}
	return true
}
