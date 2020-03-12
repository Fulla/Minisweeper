import { SET_SAFEPOINTS, START_BOARD, SET_ACTIVATED, SET_MINES, SET_STATUS } from '../constants';
import { fetchGameState, fetchStartGame, fetchDiscoverTile } from '../apiRequests';

// sync actions

function startBoard(files, columns, state) {
  return {
    type: START_BOARD,
		files,
		columns,
		state,
  }
}
  
function setSafePoints(safepoints) {
  return {
    type: SET_SAFEPOINTS,
    safepoints
  }
}

function setMines(mines) {
  return {
    type: SET_MINES,
    mines
  }
}

function setActivatedMine(activated) {
  return {
    type: SET_ACTIVATED,
    activated
  }
}

function setStatus(state) {
  return {
    type: SET_STATUS,
    state
  }
}

// thunk actions


export function startGame(f, c, m) {

  return (dispatch) => {
    
    return fetchStartGame(f, c, m)
    .then(
      (gameState) => {
				console.log(gameState)
        dispatch(startBoard(gameState.files, gameState.columns, gameState.state))
      }
    )
  }
}


export function resumeGame() {

  return (dispatch) => {
    
    return fetchGameState()
    .then(
      (gameState) => {
				console.log(gameState)
        dispatch(startBoard(gameState.files, gameState.columns, gameState.state))
				dispatch(setSafePoints(gameState.safePoints))
				dispatch(setMines(gameState.mines))
				dispatch(setActivatedMine(gameState.activatedMine))
				dispatch(setStatus(gameState.state))
      }
    )
  }
}


export function discoverTile(x, y) {

  return (dispatch) => {
    
    return fetchDiscoverTile(x, y)
    .then(
      (gameState) => {
				console.log(gameState)
				dispatch(setSafePoints(gameState.safePoints))
				dispatch(setMines(gameState.mines))
				dispatch(setActivatedMine(gameState.activatedMine))
				dispatch(setStatus(gameState.state))
      }
    )
  }
}