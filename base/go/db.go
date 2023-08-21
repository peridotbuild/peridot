package base

import (
	"context"
	"fmt"
	orderedmap "github.com/wk8/go-ordered-map/v2"
	"go.ciq.dev/pika"
)

type Pika[T any] interface {
	pika.QuerySet[T]

	F(keyval ...any) Pika[T]
	Transaction(ctx context.Context) (Pika[T], error)
	U(x any) error
}

type DB struct {
	*pika.PostgreSQL
}

type innerDB[T any] struct {
	pika.QuerySet[T]
	*DB
}

type idInterfaceForInt interface {
	GetID() int64
}

type idInterfaceForString interface {
	GetID() string
}

func NewDB(databaseURL string) (*DB, error) {
	db, err := pika.NewPostgreSQL(databaseURL)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func NewDBArgs(keyval ...any) *orderedmap.OrderedMap[string, any] {
	args := pika.NewArgs()
	for i := 0; i < len(keyval); i += 2 {
		args.Set(keyval[i].(string), keyval[i+1])
	}

	return args
}

//goland:noinspection GoExportedFuncWithUnexportedType
func Q[T any](db *DB) *innerDB[T] {
	return &innerDB[T]{pika.Q[T](db.PostgreSQL), db}
}

func (inner *innerDB[T]) F(keyval ...any) Pika[T] {
	var qs pika.QuerySet[T] = inner
	args := pika.NewArgs()
	for i := 0; i < len(keyval); i += 2 {
		args.Set(keyval[i].(string), keyval[i+1])
		qs = qs.Filter(fmt.Sprintf("%s=:%s", keyval[i].(string), keyval[i].(string)))
	}

	inner.QuerySet = qs.Args(args)
	return inner
}

func (inner *innerDB[T]) Transaction(ctx context.Context) (Pika[T], error) {
	ts := pika.NewPostgreSQLFromDB(inner.DB.DB())
	err := ts.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return &innerDB[T]{pika.Q[T](ts), inner.DB}, nil
}

func (inner *innerDB[T]) U(x any) error {
	y := x.(*T)

	// Check if x has GetID() method
	var id any
	idInterface, ok := x.(idInterfaceForInt)
	if ok {
		intID := idInterface.GetID()
		if intID == 0 {
			return fmt.Errorf("id is 0")
		}
		id = intID
	}

	idInterface2, ok := x.(idInterfaceForString)
	if ok {
		stringID := idInterface2.GetID()
		if stringID == "" {
			return fmt.Errorf("id is empty")
		}
		id = stringID
	}

	if id == nil {
		// Fallback to normal Update
		return inner.Update(y)
	}

	qs := inner.F("id", id)
	return qs.Update(y)
}
