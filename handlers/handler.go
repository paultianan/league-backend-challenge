package handlers

import (
	"fmt"
	"net/http"
	"paul-tianan/league-backend-challenge/helpers"
	"strconv"
	"strings"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	var response string
	records := helpers.FileUploader(w, r)

	if records == nil {
		helpers.ErrorHandler(w, http.StatusBadRequest, "No Records Found!")
		return
	}

	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}

	fmt.Fprint(w, response)
}

func InvertHandler(w http.ResponseWriter, r *http.Request) {
	var response string
	records := helpers.FileUploader(w, r)

	if records == nil {
		helpers.ErrorHandler(w, http.StatusBadRequest, "No Records Found!")
		return
	}

	tRecords := helpers.Transpose(records)
	for _, row := range tRecords {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}

	fmt.Fprint(w, response)
}

func FlattenHandler(w http.ResponseWriter, r *http.Request) {
	var response []string
	records := helpers.FileUploader(w, r)

	if records == nil {
		helpers.ErrorHandler(w, http.StatusBadRequest, "No Records Found!")
		return
	}

	for _, row := range records {
		response = append(response, row...)
	}

	fmt.Fprint(w, strings.Join(response, ","))
}

func SumHandler(w http.ResponseWriter, r *http.Request) {
	var sum int
	records := helpers.FileUploader(w, r)

	if records == nil {
		helpers.ErrorHandler(w, http.StatusBadRequest, "No Records Found!")
		return
	}

	for _, row := range records {
		for _, r := range row {
			i, err := strconv.Atoi(r)
			if err != nil {
				helpers.ErrorHandler(w, http.StatusBadRequest, "This record is not a number.")
				return
			}
			sum = sum + i
		}
	}

	fmt.Fprint(w, sum)
}

func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	product := 1
	records := helpers.FileUploader(w, r)

	if records == nil {
		helpers.ErrorHandler(w, http.StatusBadRequest, "No Records Found!")
		return
	}

	for _, row := range records {
		for _, r := range row {
			i, err := strconv.Atoi(r)
			if err != nil {
				helpers.ErrorHandler(w, http.StatusBadRequest, "This record is not a number.")
				return
			}
			product = product * i
		}
	}

	fmt.Fprint(w, product)
}
