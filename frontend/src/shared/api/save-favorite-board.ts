import { httpService } from '../service/http-service'

interface SaveFavoriteBoardRequest {
  boardId: string
}

export const saveFavoriteBoard = (body: SaveFavoriteBoardRequest) =>
  httpService('/api/save-favorite-board', {
    method: 'POST',
    body: body,
  })
