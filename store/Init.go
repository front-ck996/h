package store

import (
	"encoding/json"
	"fmt"
	"github.com/front-ck996/csy"
	bolt "go.etcd.io/bbolt"
	"sync"
)

type Store struct {
	DbName string
	Db     *bolt.DB
	Init   StoreInit
	Bucket string
}
type StoreInit struct {
	DbDir      string
	DbName     string
	DbFullFile string
}

var dbs = map[string]Store{}
var lock sync.Mutex

func init() {
	dbs = map[string]Store{}
}

// GetStore 获取存储对象
func GetStore(init StoreInit) (Store, error) {
	if _, ok := dbs[init.DbName]; !ok {
		if init.DbDir == "" {
			init.DbDir = "_dbs/"
		}
		_s := Store{
			Init:   init,
			DbName: init.DbName,
			Bucket: "default",
		}
		lock.Lock()
		init.DbFullFile = fmt.Sprintf("%s/%s", init.DbDir, init.DbName)
		if err := csy.NewFile().FileExistsCreateDir(init.DbFullFile); err != nil {
			return Store{}, err
		}
		db, err := bolt.Open(init.DbDir+init.DbName, 0666, nil)
		if err != nil {
			return Store{}, err
		}
		_s.Db = db
		dbs[init.DbName] = _s
		lock.Unlock()
	}
	return dbs[init.DbName], nil
}
func (s *Store) ClearBucket() {
	s.Bucket = "default"
}
func (s *Store) SetBucket(bucket string) {
	s.Bucket = bucket
}

func _decodeValue(value interface{}) []byte {
	marshal, err := json.Marshal(value)
	if err == nil {
		return marshal
	}
	// 使用类型断言将接口变量转换为 []byte 类型
	if bytes, ok := value.([]byte); ok {
		return bytes
	}
	return nil
}

func _encodeValue[T any](value []byte) T {
	var result T
	err := json.Unmarshal(value, &result)
	if err == nil {
		return result
	}
	return result
}
func _get[T any | string](db *bolt.DB, bucket string, key string) T {
	var result T

	if db == nil {
		return result
	}
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket))
		if bucket == nil {
			return fmt.Errorf("桶 mybucket 不存在")
		}
		valueCopy := bucket.Get([]byte(key))
		result = _encodeValue[T](valueCopy)
		return nil
	})
	return result
}

func _set(db *bolt.DB, bucket string, key string, value interface{}) error {
	err := db.Update(func(tx *bolt.Tx) error {
		// 打开或创建一个名为 "mybucket" 的桶（Bucket）
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		bytes := _decodeValue(value)
		err = bucket.Put([]byte(key), bytes)
		return err
	})
	return err
}
