import { START_BOARD, SET_SAFEPOINTS, SET_MINES, SET_ACTIVATED, SET_STATUS } from '../constants';

const initialState = {
  board: null,
  files: 0,
  columns: 0,
  status: "off",
}

function outside_board(x, y, state) {
  return x >= state.files || y >= state.columns || x < 0 || y < 0
}

export default function authors(state = initialState, action) {
  let newBoard = [];
  switch (action.type) {
    case START_BOARD:
      let b = [...Array(action.files)].map(item => Array(action.columns).fill("")) 
      return {
        ...state,
        board: b,
        files: action.files,
        columns: action.columns,
        status: action.state,
      }
    case SET_SAFEPOINTS:
      newBoard = []
      for (var i = 0; i < state.board.length; i++) {
        newBoard[i] = state.board[i].slice();
      } 
      for (let num in action.safepoints) {
        let sfpoints = action.safepoints[num]
        for (let sf of sfpoints) {
          if (outside_board(sf.file, sf.column, state)) {
            continue
          }
          newBoard[sf.file][sf.column] = num
        }
      }
      return {
        ...state,
        board: newBoard,
      }
    case SET_MINES:
      newBoard = []
      for (var i = 0; i < state.board.length; i++) {
        newBoard[i] = state.board[i].slice();
      }
      for (let m of action.mines) {
        if (outside_board(m.file, m.column, state)) {
          continue
        }
        newBoard[m.file][m.column] = "*"
      }
      return {
        ...state,
        board: newBoard,
      }
    case SET_ACTIVATED:
      let activ = action.activated
      if (activ == null) {
        return state
      }
      if (outside_board(activ.file, activ.column, state)) {
        return state
      }
      newBoard = [];
      for (var i = 0; i < state.board.length; i++) {
        newBoard[i] = state.board[i].slice();
      }
      newBoard[activ.file][activ.column] = "X"
      return {
        ...state,
        board: newBoard,
      }
    case SET_STATUS:
      return {
        ...state,
        status: action.state
      }
    default:
      return state
  }
}