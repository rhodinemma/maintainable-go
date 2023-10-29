package calculator_test

import (
	"log"
	"os"
	"testing"

	"github.com/rhodinemma/maintainable-go/terminal-calculator/calculator"
)

func TestMain(m *testing.M) {
	// setup statements
	setup()

	// run tests
	e := m.Run()

	// cleanup statements
	teardown()

	// report the exit code
	os.Exit(e)
}

func setup() {
	log.Println("Setting up.")
}

func teardown() {
	log.Println("Tearing down.")
}

func init() {
	log.Println("Initializing.")
}

func TestAdd(t *testing.T) {
	// Arrange
	e := calculator.Engine{}
	x, y := 2.5, 3.5
	want := 6.0

	// Act
	got := e.Add(2.5, 3.5)

	// Assert
	if got != 6.0 {
		t.Errorf("Add(%.2f,%.2f) incorrect, got: %.2f, want: %.2f.", x, y, got, want)
	}
}
