package storage

type Storage interface {
	CreateConnection() error
	Close() error
	Migrate() error
	Drop() error
}
