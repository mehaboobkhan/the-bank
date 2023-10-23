package risk

import (
	"fmt"
	"testing"
)

func TestCalculateCreditRisk(t *testing.T) {
	testCases := []struct {
		age                int
		numberOfCreditCard int
		expectedResult     string
	}{
		{20, 1, "LOW"},    // Example case with mod 0
		{23, 2, "MEDIUM"}, // Example case with mod 1
		{26, 3, "HIGH"},   // Example case with mod 2
		{18, 0, "LOW"},    // Example case with mod 0 and age 18
	}

	for _, tc := range testCases {
		t.Run(
			fmt.Sprintf("Age %d, Credit Cards %d", tc.age, tc.numberOfCreditCard),
			func(t *testing.T) {
				result := CalculateCreditRisk(tc.age, tc.numberOfCreditCard)
				if result != tc.expectedResult {
					t.Errorf("Expected result %s, but got %s", tc.expectedResult, result)
				}
			},
		)
	}
}
