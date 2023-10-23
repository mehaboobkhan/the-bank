package engine

import (
	"os"
	"reflect"
	"testing"
)

// Test DecideEngineRules function with various scenarios.
func TestDecideEngineRules(t *testing.T) {
	// Test a case that should be approved.
	data := entities.RecordData{
		Income:              110000,
		Age:                 24,
		NumberOfCredhoitCards: 3,
		PoliticallyExposed:  false,
		PhoneNumber:         "886-356-0375",
	}
	file := "../data/pre_approved_phone_no.csv"
	result := DecideEngineRules(data, file)
	if result != "approved" {
		t.Errorf("Expected 'approved' but got '%s'", result)
	}

	// Test a case that should be declined.
	data = entities.RecordData{
		Income:              90000,
		Age:                 17,
		NumberOfCreditCards: 5,
		PoliticallyExposed:  true,
		PhoneNumber:         "789-356-0375",
	}
	result = DecideEngineRules(data, file)
	if result != "declined" {
		t.Errorf("Expected 'declined' but got '%s'", result)
	}

	// Test a case with pre-approved phone number.
	data = entities.RecordData{
		PhoneNumber: "486-356-0375",
	}
	result = DecideEngineRules(data, file)
	if result != "approved" {
		t.Errorf("Expected 'approved' but got '%s'", result)
	}
}

// Test ValidatePreApprovedData function.
func TestValidatePreApprovedData(t *testing.T) {
	// Test with a valid pre-approved phone number.
	data := entities.RecordData{
		PhoneNumber: "486-356-0375",
	}
	file := "../data/pre_approved_phone_no.csv"
	valid := ValidatePreApprovedData(data, file)
	if !valid {
		t.Errorf("Expected phone number to be valid but got invalid")
	}

	// Test with an invalid phone number.
	data = entities.RecordData{
		PhoneNumber: "789-356-0375",
	}
	valid = ValidatePreApprovedData(data, file)
	if valid {
		t.Errorf("Expected phone number to be invalid but got valid")
	}
}

// Test ReadCsvFile function.
func TestReadCsvFile(t *testing.T) {
	filename := "sample.csv"

	csvContent := `header1,header2,header3
value1_1,value1_2,value1_3
value2_1,value2_2,value2_3`

	err := createSampleCSVFile(filename, csvContent)
	if err != nil {
		t.Fatalf("Error creating sample CSV file: %v", err)
	}
	defer func() {
		if err := os.Remove(filename); err != nil {
			t.Errorf("Error removing sample CSV file: %v", err)
		}
	}()

	// Test reading the sample CSV file.
	lines, err := ReadCsvFile(filename)
	if err != nil {
		t.Fatalf("Error reading CSV file: %v", err)
	}

	// Expected CSV content as a slice of slices.
	expected := [][]string{
		{"header1", "header2", "header3"},
		{"value1_1", "value1_2", "value1_3"},
		{"value2_1", "value2_2", "value2_3"},
	}

	// Compare the read CSV data with the expected data.
	if !reflect.DeepEqual(lines, expected) {
		t.Errorf("Expected:\n%v\n\nGot:\n%v", expected, lines)
	}
}

// Helper function to create a sample CSV file for testing.
func createSampleCSVFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
