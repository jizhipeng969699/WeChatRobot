package 设计原则

type Student struct {
}

func (s Student) Study(book Book) { book.BookName() }

type Book struct{}

func (b Book) BookName() {}
