package types

import "fmt"

type CmdArgRegisterFunc[T any] func(p *T, name string, shortHand string, value T, description string)

type CmdArg interface {
	String() string

	GetName() string
	GetShortHand() string
	GetDescription() string

	IsSecret() bool
}

type TypedCmdArg[T any] interface {
	CmdArg

	GetValue() T
	SetValue(value T)

	RegisterWith(fn CmdArgRegisterFunc[T])
}

type cmdArg[T any] struct {
	name        string
	shortHand   string
	description string
	value       T
	isSecret    bool
}

func NewCmdArg[T any](name, short, description string, value T, isSecret bool) TypedCmdArg[T] {
	return &cmdArg[T]{
		name:        name,
		shortHand:   short,
		description: description,
		value:       value,
		isSecret:    isSecret,
	}
}

func (c *cmdArg[T]) String() string {
	if c.isSecret {
		return fmt.Sprintf("%s: [MASKED]", c.name)
	}
	return fmt.Sprintf("%s: %v", c.name, c.value)
}

func (c *cmdArg[T]) GetName() string {
	return c.name
}

func (c *cmdArg[T]) GetShortHand() string {
	return c.shortHand
}

func (c *cmdArg[T]) GetDescription() string {
	return c.description
}

func (c *cmdArg[T]) IsSecret() bool {
	return c.isSecret
}

func (c *cmdArg[T]) GetValue() T {
	return c.value
}

func (c *cmdArg[T]) SetValue(value T) {
	c.value = value
}

func (c *cmdArg[T]) RegisterWith(fn CmdArgRegisterFunc[T]) {
	fn(&c.value, c.name, c.shortHand, c.value, c.description)
}
