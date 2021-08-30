package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/Process_Go/models"
)

func checkErr(err error) {

	if err != nil {
		panic(err)
	}
}

func main() {
	arq, err := os.Open("E:\\go\\statusinvest-energia.csv")
	checkErr(err)

	leitor := csv.NewReader(bufio.NewReader(arq))
	leitor.Comma = ';'

	var model []models.Asset
	count := 0

	for {
		line, err := leitor.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			checkErr(err)
		}

		if count > 0 {
			model = append(model, models.Asset{
				Ticker: line[0],
				Preco:  line[1],
				Dy:     line[2],
			})
		}

		count++
	}

	fmt.Println(CalcularMenorPreco(model))
	fmt.Println(CalcularMaiorDY(model))
}

func CalcularMenorPreco(model []models.Asset) string {
	menorPreco := 0.0
	ticker := ""
	for _, v := range model {
		p, err := strconv.ParseFloat(strings.Replace(v.Preco, ",", ".", -1), 64)
		checkErr(err)

		if p < menorPreco || menorPreco == 0.0 {
			menorPreco = p
			ticker = v.Ticker
		}
	}
	s := fmt.Sprintf("Ticker com menor valor é: %v - R$ %v", ticker, menorPreco)
	return s
}

func CalcularMaiorDY(model []models.Asset) string {
	maiorDY := 0.0
	ticker := ""
	for _, v := range model {
		if v.Dy != "" {
			p, err := strconv.ParseFloat(strings.Replace(v.Dy, ",", ".", -1), 64)
			checkErr(err)

			if p > maiorDY {
				maiorDY = p
				ticker = v.Ticker
			}
		}
	}
	s := fmt.Sprintf("Ticker com maior DY é: %v - %v", ticker, maiorDY)
	return s
}
