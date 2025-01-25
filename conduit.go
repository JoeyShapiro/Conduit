package main

import "github.com/boltdb/bolt"

type Conduit struct {
	db *bolt.DB
}

type Statement struct {
	bucket string
	where  func(entry any) bool
}

type View struct {
}

// as

func New(file string) (store Conduit, err error) {
	store.db, err = bolt.Open(file, 0600, nil)

	return store, err
}

func (store *Conduit) From(bucket string) (s *Statement) {
	s.bucket = bucket

	return s
}

func (s *Statement) Where(where func(entry any) bool) *Statement {
	s.where = where
	return s
}

func (s *Statement) Execute() (results []any, err error) {
	return
}
