package main

import (
	"eruaro/pulp/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"sort"
)

// CRUD books in a shelf
// GET /shelf/:section

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello!",
		})
	})

	v1 := r.Group("/shelf/v1");
	{
		v1.GET("book", getBook);
		v1.POST("book", addBook);
		v1.PUT("book", updateBook);
		v1.DELETE("book", deleteBook);
		v1.GET("shelf", getShelfBooks);
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// for easier testing, I'm adding in some "books"
func getShelfs() []models.Shelf {
	shelfs := []models.Shelf {
		models.Shelf{
			Books: []models.Book{
				models.Book{
					Title: "ABCD Song",
					Author: "ABC",
				},
				models.Book{
					Title: "ABCD Song Again",
					Author: "ABCD",
				},
			},
			Section: "A-G",
		},

		models.Shelf{
			Books: []models.Book{
				models.Book{
					Title: "Harry Poter and the Stone",
					Author: "Rowling Junior",
				},
				models.Book{
					Title: "Harry Poter and the Chamber of Secrets",
					Author: "JK Rowling",
				},
			},
			Section: "H-J",
		},

		models.Shelf{
			Books: []models.Book{
				models.Book{
					Title: "Superman",
					Author: "Bruce Wayne",
				},
				models.Book{
					Title: "Killmonger",
					Author: "T'Challa",
				},
			},
			Section: "K-S",
		},

		models.Shelf{
			Books: []models.Book{
				models.Book{
					Title: "The Last of Us",
					Author: "Zoom",
				},
				models.Book{
					Title: "The End",
					Author: "Start",
				},
			},
			Section: "T-Z",
		},
	};

	return shelfs;
}

func getBook(c *gin.Context) {
	// search book by title, or title & author 
	title := c.Query("title");
	author := c.Query("author");

	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "title is required",
		});

		return;
	}

	firstChar := title[0];

	shelfs := getShelfs();

	if firstChar >= 'A' && firstChar <= 'G' {
		shelf := shelfs[0];

		books := []models.Book{};

		if author == "" {
			for _, book := range shelf.Books {
				if book.Title == title {
					books = append(books, book);
				}
			}
		} else {
			for _, book := range shelf.Books {
				if book.Title == title && book.Author == author {
					books = append(books, book);
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"books": books,
		});

		return;
	} else if firstChar >= 'H' && firstChar <= 'J' {
		shelf := shelfs[1]

		books := []models.Book{};

		if author == "" {
			for _, book := range shelf.Books {
				if book.Title == title {
					books = append(books, book);
				}
			}
		} else {
			for _, book := range shelf.Books {
				if book.Title == title && book.Author == author {
					books = append(books, book);
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"books": books,
		});

		return;
	} else if firstChar >= 'K' && firstChar <= 'S' {
		shelf := shelfs[2]

		books := []models.Book{};

		if author == "" {
			for _, book := range shelf.Books {
				if book.Title == title {
					books = append(books, book);
				}
			}
		} else {
			for _, book := range shelf.Books {
				if book.Title == title && book.Author == author {
					books = append(books, book);
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"books": books,
		});

		return;
	} else if firstChar >= 'T' && firstChar <= 'Z' {
		shelf := shelfs[3]

		books := []models.Book{};

		if author == "" {
			for _, book := range shelf.Books {
				if book.Title == title {
					books = append(books, book);
				}
			}
		} else {
			for _, book := range shelf.Books {
				if book.Title == title && book.Author == author {
					books = append(books, book);
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"books": books,
		});

		return;
	} 

}

func addBook(c *gin.Context) {
	title := c.Query("title");
	author := c.Query("author");

	if title == "" || author == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "title and author are required",
		});

		return;
	}

	book := models.Book{
		Title: title,
		Author: author,
	}

	firstChar := title[0];

	shelfs := getShelfs();

	if firstChar >= 'A' && firstChar <= 'G' {
		shelf := shelfs[0];

		shelf.Books = append(shelf.Books, book);
	} else if firstChar >= 'H' && firstChar <= 'J' {
		shelf := shelfs[1];

		shelf.Books = append(shelf.Books, book);
	} else if firstChar >= 'K' && firstChar <= 'S' {
		shelf := shelfs[2];

		shelf.Books = append(shelf.Books, book);
	} else if firstChar >= 'T' && firstChar <= 'Z' {
		shelf := shelfs[3];

		shelf.Books = append(shelf.Books, book);
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book added!",
	});

	return;
}

func updateBook(c *gin.Context) {
	title := c.Query("title");
	author := c.Query("author");
	newTitle := c.Query("newTitle");
	newAuthor := c.Query("newAuthor");

	if title == "" || author == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "title and author is required",
		});

		return;
	}

	if (title != "" && author != "") {
		firstChar := title[0];

		shelfs := getShelfs();

		if firstChar >= 'A' && firstChar <= 'G' {
			shelf := shelfs[0];

			for i, book := range shelf.Books {
				if book.Title == title && book.Author == author {
					shelf.Books[i].Author = newAuthor;
					shelf.Books[i].Title = newTitle;
				}
			}
		} else if firstChar >= 'H' && firstChar <= 'J' {
			shelf := shelfs[1];

			for i, book := range shelf.Books {
				if book.Title == title && book.Author == author {
					shelf.Books[i].Author = newAuthor;
					shelf.Books[i].Title = newTitle;
				}
			}
		} else if firstChar >= 'K' && firstChar <= 'S' {
			shelf := shelfs[2];

			for i, book := range shelf.Books {
				if book.Title == title && book.Author == author {
					shelf.Books[i].Author = newAuthor;
					shelf.Books[i].Title = newTitle;
				}
			}
		} else if firstChar >= 'T' && firstChar <= 'Z' {
			shelf := shelfs[3];

			for i, book := range shelf.Books {
				if book.Title == title && book.Author == author {
					shelf.Books[i].Author = newAuthor;
					shelf.Books[i].Title = newTitle;
				}
			}
		}
	} 
}

// since Golang doesn't have a built in method for deleting in an array, I created a new array without the book that we wanted to "delete"
func deleteBook(c *gin.Context) {
	title := c.Query("title");
	author := c.Query("author");

	if title == "" || author == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "title and author is required",
		});

		return;
	}

	if (title != "" && author != "") {
		firstChar := title[0];

		shelfs := getShelfs();

		if firstChar >= 'A' && firstChar <= 'G' {
			shelf := shelfs[0];
			books := []models.Book{};

			for _, book := range shelf.Books {
				if book.Title != title && book.Author != author {
					books = append(books, book);
				}
			}

			shelf.Books = books;
		} else if firstChar >= 'H' && firstChar <= 'J' {
			shelf := shelfs[1];

			books := []models.Book{};

			for _, book := range shelf.Books {
				if book.Title != title && book.Author != author {
					books = append(books, book);
				}
			}

			shelf.Books = books;
		} else if firstChar >= 'K' && firstChar <= 'S' {
			shelf := shelfs[2];

			books := []models.Book{};

			for _, book := range shelf.Books {
				if book.Title != title && book.Author != author {
					books = append(books, book);
				}
			}

			shelf.Books = books;
		} else if firstChar >= 'T' && firstChar <= 'Z' {
			shelf := shelfs[3];

			books := []models.Book{};

			for _, book := range shelf.Books {
				if book.Title != title && book.Author != author {
					books = append(books, book);
				}
			}

			shelf.Books = books;
		}
	} 

	c.JSON(http.StatusOK, gin.H{
		"message": "Book deleted!",
	});

	return;
}

func getShelfBooks(c *gin.Context) {
	shelfSection := c.Query("shelfSection"); // A-G, H-J, K-S, T-Z

	if shelfSection != "A-G" || shelfSection != "H-J" || shelfSection != "K-S" || shelfSection != "T-Z" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Section not found",
		});

		return;
	}

	shelfs := getShelfs();

	if shelfSection == "A-G" {
		shelf := shelfs[0];

		sortedBooks := shelf.Books[:];

		sort.Slice(sortedBooks, func(i, j int) bool {
			return sortedBooks[i].Title < sortedBooks[j].Title;
		});

		c.JSON(http.StatusOK, gin.H{
			"books": sortedBooks,
		});

		return;
	} else if shelfSection == "H-J" {
		shelf := shelfs[1];

		sortedBooks := shelf.Books[:];

		sort.Slice(sortedBooks, func(i, j int) bool {
			return sortedBooks[i].Title < sortedBooks[j].Title;
		});

		c.JSON(http.StatusOK, gin.H{
			"books": sortedBooks,
		});

		return;
	} else if shelfSection == "K-S" {
		shelf := shelfs[2];

		sortedBooks := shelf.Books[:];

		sort.Slice(sortedBooks, func(i, j int) bool {
			return sortedBooks[i].Title < sortedBooks[j].Title;
		});

		c.JSON(http.StatusOK, gin.H{
			"books": sortedBooks,
		});

		return;
	} else if shelfSection == "T-Z" {
		shelf := shelfs[3];
	
		sortedBooks := shelf.Books[:];

		sort.Slice(sortedBooks, func(i, j int) bool {
			return sortedBooks[i].Title < sortedBooks[j].Title;
		});

		c.JSON(http.StatusOK, gin.H{
			"books": sortedBooks,
		});

		return;
	}
}
