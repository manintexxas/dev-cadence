package main

import (
	"errors"
	"fmt"
	"strings"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

var (
	ErrInvalidInput = errors.New("Input is invalid")
	ErrInvalidRange = errors.New("Input is out of range")
)

var registerName = make(map[string]bool)

type Person struct {
	id   string
	name string
	age  int
}

func genrateId() string {
	alphabeth := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	id, _ := gonanoid.Generate(alphabeth, 4)

	return id
}

func cleanedinput(input string) (string, error) {
	trimmed := strings.TrimSpace(input)

	if trimmed == "" {
		return "", fmt.Errorf("%w: Name cannot be empty", ErrInvalidInput)
	}

	return strings.ToUpper(trimmed), nil

}

func setName(input string) (string, error) {
	name, err := cleanedinput(input)
	if err != nil {
		return "", err
	}

	return name, nil

}

func setAge(input int) (int, error) {
	if input == 0 {
		return 0, fmt.Errorf("%w: Age cannot be empty", ErrInvalidRange)

	} else if input < 18 {
		return 0, fmt.Errorf("%w: Age under 18", ErrInvalidRange)

	}

	return input, nil
}

func (p *Person) ID() string {
	return p.id
}

func (p *Person) Name() string {
	return p.name
}

func (p *Person) Age() int {
	return p.age
}

func NewPerson(rawName string, rawAge int) (*Person, error) {
	name, err := setName(rawName)
	if err != nil {
		return nil, err
	}

	age, err := setAge(rawAge)
	if err != nil {
		return nil, err
	}

	if registerName[name] {
		return nil, fmt.Errorf("%w: Name already in use", ErrInvalidInput)
	}

	registerName[name] = true

	return &Person{
		id:   genrateId(),
		name: name,
		age:  age,
	}, nil

}
