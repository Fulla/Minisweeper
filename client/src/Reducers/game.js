import { START_BOARD, SET_SAFEPOINTS, SET_MINES, SET_ACTIVATED, SET_STATUS } from '../constants';

// const initialState = {
//   board: null,
//   files: 0,
//   columns: 0,
//   state: "off",
// }

const initialState = {
  board: [["2"," ","1","2","4"],["2"," ","1","2","4"],["2"," ","1","2","4"],["2"," ","1","2","4"],["2"," ","1","2","4"]],
  files: 5,
  columns: 5,
  status: "playing",
}

function outside_board(x, y, state) {
  return x >= state.files || y >= state.columns || x < 0 || y < 0
}

export default function authors(state = initialState, action) {
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
      break
    case SET_SAFEPOINTS:
      var newBoard = [];
      for (var i = 0; i < state.board.length; i++) {
        newBoard[i] = state.board[i].slice();
      }
        
      for (let num in action.safepoints) {
        let sfpoints = action.safepoints[num]
        for (let sf of sfpoints) {
          if (outside_board(sf.x, sf.y, state)) {
            continue
          }
          newBoard[sf.x][sf.y] = num
        }
      }
      return {
        ...state,
        board: newBoard,
      }
      break
    case SET_MINES:
      var newBoard = [];
      for (var i = 0; i < state.board.length; i++) {
        newBoard[i] = state.board[i].slice();
      }
      for (let m of action.mines) {
        if (outside_board(m.x, m.y, state)) {
          continue
        }
        newBoard[m.x][m.y] = "*"
      }
      return {
        ...state,
        board: newBoard,
      }
      break
    case SET_ACTIVATED:
      let m = action.activated
      if (outside_board(m.x, m.y, state)) {
        return state
      }
      var newBoard = [];
      for (var i = 0; i < state.board.length; i++) {
        newBoard[i] = state.board[i].slice();
      }
      newBoard[m.x][m.y] = "X"
      return {
        ...state,
        board: newBoard,
      }
      break
    case SET_STATUS:
      return {
        ...state,
        status: action.status
      }
      break
    default:
      return state
  }
}