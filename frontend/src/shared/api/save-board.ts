import { httpService } from '../service/http-service'

interface SaveBoardRequest {
  title: string
  backgroundId: string
}

export const saveBoard = (body: SaveBoardRequest) =>
  httpService('/api/save-board', {
    method: 'POST',
    body: body,
  })
