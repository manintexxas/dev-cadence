package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"testing"
)

func TestNewPerson_TableDriven(t *testing.T) {
	tests := []struct {
		name         string
		rawName      string
		rawAge       int
		expectedName string
		expectedAge  int
		expectedErr  error
	}{
		{
			name:         "Success - Standard Normalization",
			rawName:      "  Agung Pratama  ",
			rawAge:       25,
			expectedName: "AGUNG PRATAMA",
			expectedAge:  25,
			expectedErr:  nil,
		},
		{
			name:         "Success - Minimum Age Limit (18)",
			rawName:      "Siti Rahma",
			rawAge:       18,
			expectedName: "SITI RAHMA",
			expectedAge:  18,
			expectedErr:  nil,
		},
		{
			name:         "Success - Maximum Age Limit (99)",
			rawName:      "Kakek Sugiono",
			rawAge:       99,
			expectedName: "KAKEK SUGIONO",
			expectedAge:  99,
			expectedErr:  nil,
		},
		{
			name:         "Success - Unicode Diacritical Name",
			rawName:      "José Müller",
			rawAge:       30,
			expectedName: "JOSÉ MÜLLER",
			expectedAge:  30,
			expectedErr:  nil,
		},
		{
			name:        "Error - Empty Name",
			rawName:     "    ",
			rawAge:      20,
			expectedErr: ErrInvalidInput,
		},
		{
			name:        "Error - Name Too Short (< 2 chars)",
			rawName:     "A",
			rawAge:      20,
			expectedErr: ErrInvalidInput,
		},
		{
			name:        "Error - Name Too Long (> 50 chars)",
			rawName:     "Lorem ipsum dolor sit amet consectetur adipiscing elit sed do eiusmod",
			rawAge:      20,
			expectedErr: ErrInvalidInput,
		},
		{
			name:        "Error - Invalid Symbol or Emoji",
			rawName:     "Budi 🚀",
			rawAge:      25,
			expectedErr: ErrInvalidInput,
		},
		{
			name:        "Error - Age Under 18",
			rawName:     "Budi",
			rawAge:      17,
			expectedErr: ErrInvalidRange,
		},
		{
			name:        "Error - Age Over 99",
			rawName:     "Budi",
			rawAge:      100,
			expectedErr: ErrInvalidRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pm := NewPersonManager()
			person, err := pm.NewPerson(tt.rawName, tt.rawAge)

			if tt.expectedErr != nil {
				if err == nil {
					t.Fatalf("diharapkan error %v, tetapi mendapatkan nil", tt.expectedErr)
				}
				if !errors.Is(err, tt.expectedErr) {
					t.Errorf("diharapkan error type %v, tetapi mendapatkan: %v", tt.expectedErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("tidak diharapkan error, tetapi mendapatkan: %v", err)
			}

			if person.Name() != tt.expectedName {
				t.Errorf("diharapkan nama '%s', tetapi mendapatkan '%s'", tt.expectedName, person.Name())
			}

			if person.Age() != tt.expectedAge {
				t.Errorf("diharapkan umur %d, tetapi mendapatkan %d", tt.expectedAge, person.Age())
			}

			if person.ID() == "" {
				t.Error("ID tidak boleh kosong")
			}
		})
	}
}

func TestNewPerson_DuplicateNameHandling(t *testing.T) {
	pm := NewPersonManager()

	_, err := pm.NewPerson("Dedi Kurniawan", 30)
	if err != nil {
		t.Fatalf("pendaftaran pertama harus sukses: %v", err)
	}

	_, err = pm.NewPerson("  dedi kurniawan  ", 28)
	if err == nil {
		t.Fatal("diharapkan error duplikasi nama, tetapi pendaftaran lolos")
	}

	if !errors.Is(err, ErrInvalidInput) {
		t.Errorf("diharapkan ErrInvalidInput, tetapi mendapatkan: %v", err)
	}
}

func TestPersonManager_DeletePerson(t *testing.T) {
	pm := NewPersonManager()

	person, err := pm.NewPerson("Rian Ardianto", 26)
	if err != nil {
		t.Fatalf("gagal membuat person: %v", err)
	}

	err = pm.DeletePerson(person.ID())
	if err != nil {
		t.Fatalf("gagal menghapus person: %v", err)
	}

	err = pm.DeletePerson(person.ID())
	if err == nil {
		t.Fatal("diharapkan error menghapus ID yang tidak ada, tetapi mendapatkan nil")
	}

	_, err = pm.NewPerson("Rian Ardianto", 26)
	if err != nil {
		t.Fatalf("seharusnya nama bisa didaftarkan ulang setelah dihapus: %v", err)
	}
}

func TestNewPerson_ConcurrencyAndRaceCondition(t *testing.T) {
	pm := NewPersonManager()
	concurrentRequests := 20
	var wg sync.WaitGroup

	registeredIDs := make(chan string, concurrentRequests)

	for i := 0; i < concurrentRequests; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			test := strings.Repeat("I", index + 2)
			name := fmt.Sprintf("%s", test)
			person, err := pm.NewPerson(name, 25)
			if err == nil {
				registeredIDs <- person.ID()
			}
		}(i)
	}

	wg.Wait()
	close(registeredIDs)

	idMap := make(map[string]bool)
	for id := range registeredIDs {
		if idMap[id] {
			t.Errorf("ditemukan duplikasi ID hasil konkuransi: %s", id)
		}
		idMap[id] = true
	}

	if len(idMap) != concurrentRequests {
		t.Errorf("diharapkan %d orang terdaftar, tetapi hanya %d yang sukses", concurrentRequests, len(idMap))
	}
}