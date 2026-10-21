package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

var (
	ErrInvalidInput = errors.New("Input is invalid")
	ErrInvalidRange = errors.New("Input is out of range")
)

type PersonManager struct {
	mu           sync.Mutex
	registerName map[string]bool
}

func NewPersonManager() *PersonManager {
	return &PersonManager{registerName: make(map[string]bool)}
}

type Person struct {
	id   string
	name string
	age  int
}

func generateId() string {
	alphabeth := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	id, _ := gonanoid.Generate(alphabeth, 4)

	return id
}

func cleanedInput(input string) (string, error) {
	trimmed := strings.TrimSpace(input)

	if trimmed == "" {
		return "", fmt.Errorf("%w: Name cannot be empty", ErrInvalidInput)
	}

	return strings.ToUpper(trimmed), nil

}

func setName(input string) (string, error) {
	name, err := cleanedInput(input)
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

func (m *PersonManager) NewPerson(rawName string, rawAge int) (*Person, error) {
	name, err := setName(rawName)
	if err != nil {
		return nil, err
	}

	age, err := setAge(rawAge)
	if err != nil {
		return nil, err
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	if m.registerName[name] {
		return nil, fmt.Errorf("%w: Name already in use", ErrInvalidInput)
	}

	m.registerName[name] = true

	return &Person{
		id:   generateId(),
		name: name,
		age:  age,
	}, nil

}
