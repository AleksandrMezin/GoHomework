package postgres

import "DataBase/pkg/postgres"

type Interface interface {
	Tasks(int, int) ([]postgres.Task, error)
	NewTask(postgres.Task) (int, error)
}
