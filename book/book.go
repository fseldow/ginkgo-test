package book

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Book decribe a book
type Book struct {
	Title  string
	Author string
	Pages  int
}

// CategoryByLength returns category by pages
func (b *Book) CategoryByLength() string {
	if b.Pages > 100 {
		return "NOVEL"
	}
	return "SHORT STORY"
}

var _ = Describe("Book", func() {
	var (
		longBook  Book
		shortBook Book
	)

	BeforeEach(func() {
		longBook = Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  1488,
		}

		shortBook = Book{
			Title:  "Fox In Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}
	})

	Describe("Categorizing book length", func() {
		Context("With more than 300 pages", func() {
			It("should be a novel", func() {
				Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
			})
		})

		Context("With fewer than 300 pages", func() {
			It("should be a short story", func() {
				Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
			})
		})
	})
})
