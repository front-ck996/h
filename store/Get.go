package store

func Get[T any](s Store, key string) T {
	return _get[T](s.Db, s.Bucket, key)
}

func Set(s Store, key string, value interface{}) error {
	return _set(s.Db, s.Bucket, key, value)
}
