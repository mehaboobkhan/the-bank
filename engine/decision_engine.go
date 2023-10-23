package engine

import (
	"encoding/csv"
	"os"
)

// DecideEngineRules validates all the rules for credit approval
func DecideEngineRules(data entities.RecordData, file string) string {

	if ValidatePreApprovedData(data, file) {
		return "approved"
	}

	if data.Income > 100000 &&
		data.Age >= 18 &&
		data.NumberOfCreditCards < 4 &&
		risk.CalculateCreditRisk(data.Age, data.NumberOfCreditCards) == "LOW" &&
		data.PoliticallyExposed == false &&
		(data.PhoneNumber[0:1] == "0" || data.PhoneNumber[0:1] == "2" || data.PhoneNumber[0:1] == "5" || data.PhoneNumber[0:1] == "8") {
		return "approved"
	}
	return "declined"
}

// ValidatePreApprovedData validates pre-approved data(phone numbers)
func ValidatePreApprovedData(data entities.RecordData, file string) bool {
	csvData, err := ReadCsvFile(file)
	if err != nil {
		panic(err)
	}
	for _, line := range csvData {
		preApprovedData := entities.PreApprovedData{
			PhoneNumber: line[0],
		}
		if data.PhoneNumber == preApprovedData.PhoneNumber {
			return true
		}
	}
	return false
}

// ReadCsvFile is an helper function to read CSV data.
func ReadCsvFile(filename string) ([][]string, error) {
	fileContent, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}

	defer func() {
		if err := fileContent.Close(); err != nil {
			panic(err)
		}
	}()

	lines, err := csv.NewReader(fileContent).ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return lines, nil
}
