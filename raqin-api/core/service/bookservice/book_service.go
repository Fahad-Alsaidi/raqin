package bookservice

// book interface the will hold all book related operations
type IBook interface {
	NewBook()  // registers new book
	Books()    // gets all books
	BookByID() // get book by id
}

// book struct that will implement IBook interface
type bookService struct{}

// NewBookController return bookController struct with IBook service
// this method is the entry to this controller
func NewBookService() bookService {
	return bookService{}
}

func (bs bookService) NewBook() {

}

func (bs bookService) Books() {

}

func (bs bookService) BookByID() {

}
