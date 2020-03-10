import React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import * as GameActions from '../Actions/gameActions';
import '../App.css';
import Board from '../Components/Board';
import Button from '../Components/Button';

class Game extends React.Component {
    state = {}

    render() {
			const { board } = this.props.game
			const { startGame, resumeGame } = this.props.gameActions

			console.log(board);
			

			return (
				<div className="App">
					<h2>Minesweeper</h2>
					<Board board={board} />
					<div className="Buttons">
						<Button title="New game" action={startGame}/>
						<Button title="Resume game" action={resumeGame}/>
					</div>
				</div>
			)
		}

}

export default connect(
  state => ({
    game: state.game,
  }),
  dispatch => ({
    gameActions: bindActionCreators(GameActions, dispatch),
  })
)(Game)