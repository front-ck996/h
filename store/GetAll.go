package store

func GetAll[T any](s Store) []T {
	return _getAll[T](s.Db, s.Bucket)
}
