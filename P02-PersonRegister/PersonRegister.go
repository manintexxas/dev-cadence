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

	safeNameRegex = regexp.MustCompile(`^[\p{L}\s'\-]+$`)
)

type PersonManager struct {
	sync.Mutex
	registerName map[string]bool
	registerID   map[string]string
}

func NewPersonManager() *PersonManager {
	return &PersonManager{
		registerName: make(map[string]bool),
		registerID:   make(map[string]string),
	}
}

type Person struct {
	id   string
	name string
	age  int
}

func cleanedInput(input string) (string, error) {
	trimmed := strings.TrimSpace(input)

	if trimmed == "" {
		return "", fmt.Errorf("%w: Name cannot be empty", ErrInvalidInput)
	}

	return strings.ToUpper(trimmed), nil
}

func (p *Person) setName(input string) error {
	const minName = 2
	const maxName = 50

	name, err := cleanedInput(input)
	if err != nil {
		return err
	}

	if len(name) < minName || len(name) > maxName {
		return fmt.Errorf("%w: name length must be between %d and %d characters", ErrInvalidInput, minName, maxName)
	}

	if !safeNameRegex.MatchString(name) {
		return fmt.Errorf("%w: name contains invalid symbols or emojis", ErrInvalidInput)
	}

	p.name = name
	return nil
}

func (p *Person) setAge(input int) error {
	const minAge = 18
	const maxAge = 99

	if input < minAge || input > maxAge {
		return fmt.Errorf("%w: Age must be between %d and %d", ErrInvalidRange, minAge, maxAge)
	}

	p.age = input
	return nil
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
	p := &Person{}

	if err := p.setName(rawName); err != nil {
		return nil, err
	}

	if err := p.setAge(rawAge); err != nil {
		return nil, err
	}

	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	var id string

	for {
		generated, err := gonanoid.Generate(alphabet, 10)
		if err != nil {
			continue
		}

		m.Lock()
		if m.registerName[p.name] {
			m.Unlock()
			return nil, fmt.Errorf("%w: Name already in use", ErrInvalidInput)
		}

		if _, exists := m.registerID[generated]; !exists {
			id = generated
			m.registerName[p.name] = true
			m.registerID[id] = p.name
			m.Unlock()
			break
		}
		m.Unlock()
	}

	p.id = id
	return p, nil
}

func (m *PersonManager) DeletePerson(id string) error {
	m.Lock()
	defer m.Unlock()

	name, exists := m.registerID[id]
	if !exists {
		return fmt.Errorf("%w: Person ID not found", ErrInvalidInput)
	}

	delete(m.registerID, id)
	delete(m.registerName, name)

	return nil
}