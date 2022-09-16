package models 

type Shelf struct {
	Books []Book;
	Section string `json:"string"`; // (A-G, H-J, K-T, T-Z)
}
