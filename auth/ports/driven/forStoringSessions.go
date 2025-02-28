package driven

type StoringSessions[T any] interface {
  Close()
  Insert(item T, sql string) error
  Get(id, table string) (T, error)
  GetByField(column, value, table string) (T, error)
  GetAll(table string) []T
  Delete(id, table string) error
  DeleteByField(column, value, table string) error
  GetSQL(sql string, item T) (T, error)
}
