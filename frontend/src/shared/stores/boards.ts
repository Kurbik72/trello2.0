import { defineStore } from 'pinia'
import { ref } from 'vue'
import { saveBoards } from '../api/save-boards'
import { getBoards } from '../api/get-boards'
import { saveFavoriteBoard } from '../api/save-favorite-board'

interface SaveBoardRequest {
  title: string
  backgroundId: string
  user_id: string
}

interface BoardInterface {
  id: string
  title: string
  isFavorite: boolean
  linkToBoard: string
  backgroundSrc: string
}
export const useBoardsStore = defineStore('boards', () => {
  const boards = ref<BoardInterface[]>([])

  const saveBoard = async (board: SaveBoardRequest) => {
    const { data, error, execute } = saveBoards(board)
    await execute()
    if (error.value) {
      throw new Error(error.value.errorMessage)
    }
    if (data.value) {
      boards.value.push(data.value)
    }
  }

  const getBoardsList = async (user_id: string) => {
    const { data, error, execute } = getBoards(user_id)
    await execute()
    if (error.value) {
      throw new Error(error.value.errorMessage)
    }
    if (data.value) {
      boards.value = [...data.value]
    }
  }

  const favoriteBoardList = async (board_id: string, user_id: string) => {
    const { error, execute, data } = saveFavoriteBoard({ boardId: board_id, user_id: user_id })
    await execute()
    if (error.value) {
      throw new Error(error.value.errorMessage)
    }
    if (data.value) {
      boards.value = boards.value.map((favoriteBoard) =>
        favoriteBoard.id === board_id ? data.value! : favoriteBoard,
      )
    }
  }
  return { boards, saveBoard, getBoardsList, favoriteBoardList }
})
