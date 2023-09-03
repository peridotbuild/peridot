// Copyright 2023 Peridot Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package base

import (
	"context"
	"errors"
	"fmt"
	orderedmap "github.com/wk8/go-ordered-map/v2"
	"go.ciq.dev/pika"
	"math/rand"
)

type Pika[T any] interface {
	pika.QuerySet[T]

	U(x any) error
	F(keyval ...any) Pika[T]
	D(x any) error
	Transaction(ctx context.Context) (Pika[T], error)
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

func NameGen(prefix string) string {
	// last part is a random int64
	// generate a length=18 random int
	minRan := 100000000000000000
	maxRan := 999999999999999999
	ran := minRan + rand.Intn(maxRan-minRan)

	return fmt.Sprintf("%s/%d", prefix, ran)
}

//goland:noinspection GoExportedFuncWithUnexportedType
func Q[T any](db *DB) Pika[T] {
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

func (inner *innerDB[T]) D(x any) error {
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
		return errors.New("id is nil")
	}

	qs := inner.F("name", id)
	return qs.Delete()
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

	qs := inner.F("name", id)
	return qs.Update(y)
}
