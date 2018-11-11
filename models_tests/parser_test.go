package models_tests

import (
	"fmt"
	"testing"

	"github.com/mwhagedorn/gorocksim/models"
)

func TestParserSelection(t *testing.T) {

	parser := models.NewEngineParser("../engines/Estes_A8.rse")
	equals(t, "../engines/Estes_A8.rse", parser.FileName)
	parser.Parse()
	fmt.Printf("Datapoints are %v\n", parser.Data)
}
