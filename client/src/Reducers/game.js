import { START_BOARD, SET_SAFEPOINTS } from '../constants';

const initialState = {
  board: null,
  files: 0,
  columns: 0,
  state: "off",
}

export default function authors(state = initialState, action) {
  switch (action.type) {
    case START_BOARD:
      let b = [...Array(action.files)].map(item => Array(action.columns).fill(" ")) 
      return {
        ...state,
        board: b,
        files: action.files,
        columns: action.columns,
        state: action.state,
      }
    case SET_SAFEPOINTS:
      let board = state.board
      for (let num in action.safepoints) {
        let sfpoints = action.safepoints[num]
        for (let sf of sfpoints) {
          if (sf.x >= state.files || sf.y >= state.columns || sf.x < 0 || sf.y < 0) {
            continue
          }
          board[sf.x][sf.y] = num
        }
      }
      return {
        ...state,
        board: b,
      }
    default:
      return state
  }
}