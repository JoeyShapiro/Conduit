package main

import (
	"encoding/binary"
	"encoding/json"

	"github.com/boltdb/bolt"
)

type Statement struct {
	store  *Conduit
	bucket string
	where  func(entry any) bool
}

type Conduit struct {
	db *bolt.DB
}

type View struct {
}

// tolist

// as

func New(file string) (store Conduit, err error) {
	store.db, err = bolt.Open(file, 0600, nil)

	return store, err
}

func (store *Conduit) Create(bucket string) error {
	err := store.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(bucket))
		return err
	})
	return err
}

func (store *Conduit) From(bucket string) *Statement {
	s := &Statement{
		store:  store,
		bucket: bucket,
	}

	return s
}

func (s *Statement) Where(where func(entry any) bool) *Statement {
	s.where = where
	return s
}

func (s *Statement) Insert(object any) error {
	content, err := json.Marshal(object)
	if err != nil {
		return err
	}

	err = s.store.db.Update(func(tx *bolt.Tx) (err error) {
		bucket := tx.Bucket([]byte(s.bucket))

		primaryKey, err := bucket.NextSequence()
		if err != nil {
			return err
		}

		// convert pk to unsigned int byte array
		bePK := make([]byte, 8)
		binary.BigEndian.PutUint64(bePK, uint64(primaryKey))

		// store in the bucket and return the error, if any
		return bucket.Put(bePK, content)
	})
	if err != nil {
		return err
	}

	return err
}

func (s *Statement) Update(pk uint64, object any) (err error) {
	content, err := json.Marshal(object)
	if err != nil {
		return err
	}

	err = s.store.db.Update(func(tx *bolt.Tx) (err error) {
		bucket := tx.Bucket([]byte(s.bucket))

		// convert pk to unsigned int byte array
		bePK := make([]byte, 8)
		binary.BigEndian.PutUint64(bePK, uint64(pk))

		// store in the bucket and return the error, if any
		return bucket.Put(bePK, content)
	})
	if err != nil {
		return err
	}

	return err
}

func (s *Statement) Execute() (results []any, err error) {
	// create the view for the db
	err = s.store.db.View(func(tx *bolt.Tx) error {
		// open the bucket using its name
		bucket := tx.Bucket([]byte(s.bucket))

		// for every row
		bucket.ForEach(func(k, v []byte) error {
			// convert the byte array into an entry
			var entry any
			err := json.Unmarshal(v, &entry)
			if err != nil {
				return err
			}

			// if the entry matches the where condition specified by the user
			if s.where(entry) {
				// append to the results
				results = append(results, entry)
			}

			return nil
		})
		return nil
	})

	return
}
