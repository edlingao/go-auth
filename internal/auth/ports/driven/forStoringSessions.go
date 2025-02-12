package driven

type StoringSessions[T any] interface {
  Close()
  Insert(item T, sql string) error
  Get(id, table string) (T, error)
  GetAll(table string) []T
  Delete(id, table string) error
  GetByField(field, value, table string) (T, error)
  GetSQL(sql string) (T, error)
}
