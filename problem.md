# Problem
Today you have been asked to create an API using the framework Gin (https://pkg.go.dev/github.com/gin-gonic/gin). This project will test your skills in creating functional backends using Golang, will also test your ability to make development code and test your understanding of Golang.

# Requirements
- Create an API to catalog books, using the models of `shelf`, `book` , `author`.
- Shelf contains books in a certain books starting with specific letters section, e.g. A-G, H-J, K-T, T-Z

Book -> {
    title
    author -> String
}

Shelf -> {
    books[] 
    section (A-G, H-J, K-S, T-Z)
}

