package models

type BookCategory struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type CreateBookCategory struct {
	Name string `json:"name"`
}

type GetAllBookCategoryResponse struct {
	BookCategoryList []BookCategory `json:"book_category_list"`
	Count            int32          `json:"count"`
}
type UpdateBookCategory struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type MsgModel struct {
	Msg string `json:"msg"`
}
