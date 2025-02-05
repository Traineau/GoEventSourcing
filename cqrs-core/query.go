package cqrs_core

import "fmt"

type QueryMessage interface {
	Payload() interface{}
	QueryType() string
}

type QueryBus struct {
	handlers map[string]QueryHandler
}

func NewQueryBus() *QueryBus {
	cBus := &QueryBus{
		handlers: make(map[string]QueryHandler),
	}

	return cBus
}

func (b *QueryBus) Dispatch(query QueryMessage) error {
	if handler, ok := b.handlers[query.QueryType()]; ok {
		return handler.Handle(query)
	}
	return fmt.Errorf("the query bus does not have a handler for query of type: %s", query.QueryType())
}

func (b *QueryBus) RegisterHandler(handler QueryHandler, query interface{}) error {
	typeName := typeOf(query)
	if _, ok := b.handlers[typeName]; ok {
		return fmt.Errorf("duplicate query handler registration with query bus for query of type: %s", typeName)
	}

	b.handlers[typeName] = handler

	return nil
}

type QueryDescriptor struct {
	query interface{}
}

func NewQueryMessage(query interface{}) *QueryDescriptor {
	return &QueryDescriptor{
		query: query,
	}
}

func (c *QueryDescriptor) QueryType() string {
	return typeOf(c.query)
}

// Command returns the actual command payload of the message.
func (c *QueryDescriptor) Payload() interface{} {
	return c.query
}
