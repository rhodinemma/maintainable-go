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

	actAssert := func(x, y, want float64) {
		// Act
		got := e.Add(x, y)

		// Assert
		if got != want {
			t.Errorf("Add(%.2f,%.2f) incorrect, got: %.2f, want: %.2f.", x, y, got, want)
		}
	}

	t.Run("positive input", func(t *testing.T) {
		x, y := 2.5, 3.5
		want := 6.0
		actAssert(x, y, want)
	})

	t.Run("negative input", func(t *testing.T) {
		x, y := -2.5, -3.5
		want := -6.0
		actAssert(x, y, want)
	})
}
