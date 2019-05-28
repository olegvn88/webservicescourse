package books

const novel = "NOVEL"
const shortStory = "SHORT STORY"

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func (book *Book) CategoryByLength() string {

	if book.Pages < 300 {
		return shortStory
	} else {
		return novel
	}
}

func (book *Book) AuthorLastName() string {
	return book.Author
}

func GetResult(num1 int, num2 int) int {

	return num1 * num2
}
