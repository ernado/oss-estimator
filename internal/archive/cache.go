package archive

import (
	"encoding/binary"

	"github.com/go-faster/errors"
	lru "github.com/hashicorp/golang-lru/v2"
	"go.etcd.io/bbolt"
)

type UserCache struct {
	db  *bbolt.DB
	lru *lru.Cache[int64, struct{}]
}

var (
	_bucket        = []byte("id-to-actor")
	_bucketInverse = []byte("actor-to-id")
)

type key [8]byte

func (u *UserCache) inDB(k key) (found bool, err error) {
	if err := u.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(_bucket)
		if b == nil {
			return nil
		}
		found = b.Get(k[:]) != nil
		return nil
	}); err != nil {
		return false, errors.Wrap(err, "view")
	}
	return found, nil
}

func (u *UserCache) key(id int64) key {
	var k [8]byte
	binary.BigEndian.PutUint64(k[:], uint64(id))
	return k
}

func (u *UserCache) put(k key, v []byte) error {
	// Set both index values.
	if err := u.db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(_bucket)
		if err != nil {
			return errors.Wrap(err, "create bucket")
		}
		if err := b.Put(k[:], v); err != nil {
			return errors.Wrap(err, "put")
		}
		bi, err := tx.CreateBucketIfNotExists(_bucketInverse)
		if err != nil {
			return errors.Wrap(err, "create inverse bucket")
		}
		if err := bi.Put(v, k[:]); err != nil {
			return errors.Wrap(err, "put")
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "db")
	}
	return nil
}

func (u *UserCache) Close() error {
	u.lru.Purge()
	if err := u.db.Sync(); err != nil {
		return errors.Wrap(err, "sync")
	}
	if err := u.db.Close(); err != nil {
		return errors.Wrap(err, "close")
	}
	return nil
}

func (u *UserCache) hasOrAdd(id int64) bool {
	if _, ok := u.lru.Get(id); ok {
		return true
	}
	u.lru.Add(id, struct{}{})
	return false
}

func (u *UserCache) HasName(id int64) (bool, error) {
	if u.hasOrAdd(id) {
		return true, nil
	}
	var hit bool
	k := u.key(id)
	if err := u.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(_bucket)
		if b == nil {
			return nil
		}
		hit = b.Get(k[:]) != nil
		return nil
	}); err != nil {
		return false, errors.Wrap(err, "view")
	}
	return hit, nil
}

func (u *UserCache) Get(name []byte) (int64, error) {
	var id int64
	if err := u.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(_bucketInverse)
		if b == nil {
			return nil
		}
		v := b.Get(name)
		if v == nil {
			return nil
		}
		id = int64(binary.BigEndian.Uint64(v))
		return nil
	}); err != nil {
		return 0, errors.Wrap(err, "view")
	}
	return id, nil
}

func (u *UserCache) Add(id int64, v []byte) error {
	if id == 0 {
		// No-op.
		return nil
	}
	if u.hasOrAdd(id) {
		return nil
	}
	k := u.key(id)
	if err := u.put(k, v); err != nil {
		return errors.Wrap(err, "put")
	}
	return nil
}

type UserCacheOptions struct {
	Path     string
	Size     int
	ReadOnly bool
}

func NewUserCache(opt UserCacheOptions) (*UserCache, error) {
	if opt.Size <= 0 {
		opt.Size = 1000
	}
	db, err := bbolt.Open(opt.Path, 0666, &bbolt.Options{
		NoSync:         true,
		NoFreelistSync: true,
		ReadOnly:       opt.ReadOnly,
	})
	if err != nil {
		return nil, errors.Wrap(err, "db open")
	}
	cache, err := lru.New[int64, struct{}](opt.Size)
	if err != nil {
		return nil, errors.Wrap(err, "cache")
	}
	return &UserCache{
		db:  db,
		lru: cache,
	}, nil
}
