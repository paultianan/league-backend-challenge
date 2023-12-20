package helpers

import (
	"encoding/csv"
	"net/http"
)

func FileUploader(w http.ResponseWriter, r *http.Request) [][]string {
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error: "+err.Error())
		return nil
	}

	file, fH, err := r.FormFile("file")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Error retrieving the file: "+err.Error())
		return nil
	}
	defer file.Close()

	// Validate the file size
	maxSize := int64(5 << 20) // 5 MB limit
	if fH.Size > maxSize {
		ErrorHandler(w, http.StatusBadRequest, "File size exceeds the limit")
		return nil
	}

	if !isValidFileType(fH.Filename, ".csv") {
		ErrorHandler(w, http.StatusBadRequest, "Invalid file type. Only .csv files are allowed.")
		return nil
	}

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest, "Error while reading the csv: "+err.Error())
		return nil
	}

	return records
}

func Transpose(records [][]string) [][]string {
	// Know the number of columns and rows
	col := len(records[0]) // column
	row := len(records)    // rows

	// create new 2d slice. We will going to store the datas here
	result := make([][]string, col)
	for i := range result {
		result[i] = make([]string, row)
	}
	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			result[i][j] = records[j][i]
		}
	}
	return result
}

func ErrorHandler(w http.ResponseWriter, statusCode int, errorMessage string) {
	http.Error(w, errorMessage, statusCode)
}

func isValidFileType(filename, extension string) bool {
	return len(filename) > len(extension) && filename[len(filename)-len(extension):] == extension
}
