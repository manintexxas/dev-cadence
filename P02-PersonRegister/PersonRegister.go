package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"sync"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

var (
	ErrInvalidInput = errors.New("Input is invalid")
	ErrInvalidRange = errors.New("Input is out of range")

	safeNameRegex = regexp.MustCompile(`^[a-zA-Z\s'\-]+$`)
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
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	id, _ := gonanoid.Generate(alphabet, 4)

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
	const minName = 2
	const maxName = 50

	name, err := cleanedInput(input)
	if err != nil {
		return "", err
	}

	if len(input) < minName || len(input) > maxName {
		return "", fmt.Errorf("%w: name length must be between %d and %d characters", ErrInvalidInput, minName, maxName)
	}

	if !safeNameRegex.MatchString(name) {
		return "", fmt.Errorf("%w: name contains invalid symbols or emojis", ErrInvalidInput)
	}

}

func setAge(input int) (int, error) {
	const minAge = 18
	const maxAge = 99

	if input <= 0 {
		return 0, fmt.Errorf("%w: Age cannot be empty or negative", ErrInvalidRange)

	} else if input < minAge || input > maxAge {
		return 0, fmt.Errorf("%w: Age must be between %d and %d", ErrInvalidRange, minAge, maxAge)

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
