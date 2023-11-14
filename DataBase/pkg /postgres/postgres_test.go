package postgres

import (
	"os"
	"testing"
)

var s *Storage

func TestMain(m *testing.M) {
	pwd := os.Getenv("dbpass")
	if pwd == "" {
		m.Run()
	}
}
func TestStorage_Tasks(t *testing.T) {
	data, err := s.Tasks(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
	data, err = s.Tasks(1, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}
func TestStorage_NewTask(t *testing.T) {
	task := Task{
		Title:   "Unit Test Task",
		Content: "Task Content",
	}
	id, err := s.NewTask(task)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Создана задача с id:", id)
	tasks, err := s.Tasks(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tasks)
}
