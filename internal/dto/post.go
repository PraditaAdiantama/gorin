package dto

type CreatePostDTO struct {
	Title   string `json:"title" binding:"required,min=3,max=100"`
	Content string `json:"content" binding:"required,min=5"`
}

type UpdatePostDTO struct {
    Title string `json:"title" binding:"omitempty,min=3,max=100"`
    Content string `json:"content" binding:"omitempty,min=5"`
}
