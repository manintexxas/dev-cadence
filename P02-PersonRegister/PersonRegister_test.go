package main

import (
	"errors"
	"testing"
)

func TestNewPerson_Success(t *testing.T) {
	name := "  Agung Pratama  "
	age := 25

	person, err := NewPerson(name, age)
	if err != nil {
		t.Fatalf("diharapkan sukses, tetapi mendapatkan error: %v", err)
	}

	if person.Name() != "AGUNG PRATAMA" {
		t.Errorf("diharapkan nama 'AGUNG PRATAMA', tetapi mendapatkan '%s'", person.Name())
	}

	if person.Age() != age {
		t.Errorf("diharapkan umur %d, tetapi mendapatkan %d", age, person.Age())
	}

	if person.ID() == "" {
		t.Error("ID tidak boleh kosong")
	}
}

func TestNewPerson_InvalidName_Empty(t *testing.T) {
	_, err := NewPerson("   ", 20)
	if err == nil {
		t.Fatal("diharapkan error untuk nama kosong, tetapi mendapatkan nil")
	}

	if !errors.Is(err, ErrInvalidInput) {
		t.Errorf("diharapkan ErrInvalidInput, tetapi mendapatkan: %v", err)
	}
}

func TestNewPerson_InvalidAge_Under18(t *testing.T) {
	_, err := NewPerson("Budi", 17)
	if err == nil {
		t.Fatal("diharapkan error untuk umur di bawah 18, tetapi mendapatkan nil")
	}

	if !errors.Is(err, ErrInvalidRange) {
		t.Errorf("diharapkan ErrInvalidRange, tetapi mendapatkan: %v", err)
	}
}

func TestNewPerson_DuplicateNameInMap(t *testing.T) {
	uniqueName := "Dedi Kurniawan"
	_, err := NewPerson(uniqueName, 30)
	if err != nil {
		t.Fatalf("pendaftaran pertama harus sukses, dapat error: %v", err)
	}

	_, err = NewPerson("  dedi kurniawan  ", 28)
	if err == nil {
		t.Fatal("diharapkan error untuk nama duplikat, tetapi pendaftaran kedua lolos")
	}

	if !errors.Is(err, ErrInvalidInput) {
		t.Errorf("diharapkan ErrInvalidInput untuk duplikasi nama, tetapi mendapatkan: %v", err)
	}
}
