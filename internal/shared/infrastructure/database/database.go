package database

type Database interface {
	Open() error
	Close() error
	Ping() error
}
