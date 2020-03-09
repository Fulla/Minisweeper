import { SET_SAFEPOINTS, START_BOARD } from '../constants';
import { fetchGameState, fetchStartGame } from '../apiRequests';

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

// thunk actions


export function getGameState() {

  return (dispatch, getState) => {
    
    return fetchGameState()
    .then(
      (gameState) => {
        dispatch(startBoard(gameState.files, gameState.columns, gameState.state))
        dispatch(setSafePoints(gameState.safepoints))
      }
    )
  }
}


export function startGame() {

  return (dispatch, getState) => {
    
    return fetchStartGame()
    .then(
      (gameState) => {
        dispatch(startBoard(gameState.files, gameState.columns, gameState.state))
        dispatch(setSafePoints(gameState.safepoints))
      }
    )
  }
}
