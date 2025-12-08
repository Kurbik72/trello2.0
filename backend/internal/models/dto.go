package models

// Request DTOs
type SaveBoardRequest struct {
	Title        string `json:"title" binding:"required,min=1,max=255"`
	BackgroundID string `json:"backgroundId" binding:"required,uuid"`
	UserID       string `json:"user_id" binding:"required,uuid"`
}

type SaveFavoriteBoardRequest struct {
	BoardID string `json:"boardId" binding:"required,uuid"`
}

// Response DTOs
type BoardResponse struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	IsFavorite    bool   `json:"isFavorite"`
	LinkToBoard   string `json:"linkToBoard"`
	BackgroundSrc string `json:"backgroundSrc"`
}

type DefaultBackgroundResponse struct {
	ID  string `json:"id"`
	Src string `json:"src"`
}

// Converters
func BoardToResponse(board *Board) *BoardResponse {
	return &BoardResponse{
		ID:            board.ID,
		Title:         board.Title,
		IsFavorite:    board.IsFavorite,
		LinkToBoard:   board.LinkToBoard,
		BackgroundSrc: board.BackgroundSrc,
	}
}

func BoardsToResponse(boards []Board) []BoardResponse {
	result := make([]BoardResponse, len(boards))
	for i, board := range boards {
		result[i] = *BoardToResponse(&board)
	}
	return result
}

func DefaultBackgroundToResponse(bg *DefaultBackground) *DefaultBackgroundResponse {
	return &DefaultBackgroundResponse{
		ID:  bg.ID,
		Src: bg.Src,
	}
}

func DefaultBackgroundsToResponse(backgrounds []DefaultBackground) []DefaultBackgroundResponse {
	result := make([]DefaultBackgroundResponse, len(backgrounds))
	for i, bg := range backgrounds {
		result[i] = *DefaultBackgroundToResponse(&bg)
	}
	return result
}

