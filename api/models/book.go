package models

type Book struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Category_id string `json:"category_id"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}
type CreateBook struct {
	Name        string `json:"name"`
	Category_id string `json:"category_id"`
}

type GetAllBookResponse struct {
	BookList []Book `json:"book_list"`
	Count    int32  `json:"count"`
}
type GetBookResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}
type UpdateBook struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Category_id string `json:"category_id"`
}
