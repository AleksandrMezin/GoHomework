package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Пакет демонстрирует выполнение принципа LSP в Go.

// Интерфейс с контрактом на представление
// значения некоторого типа данных в виде
// последовательности байт.
type Serializer interface {
	Serialize() []byte
}

type String string

func (s String) Serialize() []byte {
	return []byte(s)
}

type Str struct {
	one string
	two string
}

func (s Str) Serialize() []byte {
	return []byte(fmt.Sprintf("%v", s))
}

type Bool bool

func (b Bool) Serialize() []byte {
	return []byte(fmt.Sprintf("%v", b))
}

// Функция принимает на вход объект, в который требуется
// сохранить состояние: файл, принтер и т.д.
// Также принимается любое количество объектов, поддерживающих
// сериализацию.
func WriteObject(w io.Writer, objects ...Serializer) error {
	for _, obj := range objects {
		b := obj.Serialize()
		_, err := w.Write(b)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	var s String = "ABC"
	var b Bool = true
	c := Str{
		one: "one",
		two: "two",
	}
	f, err := os.Create("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = WriteObject(f, s, b, c)
	if err != nil {
		log.Fatal(err)
	}
}
