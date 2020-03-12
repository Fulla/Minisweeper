import React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import * as GameActions from '../Actions/gameActions';
import '../App.css';
import Board from '../Components/Board';
import Button from '../Components/Button';

class Game extends React.Component {
    state = {
			files: 5,
			columns: 5,
			mines: 5,
		}

		startGame = () => {
			const { files, columns, mines } = this.state
			this.props.gameActions.startGame(files, columns, mines)
		}

		filesChange = (e) => {
			let val = parseInt(e.target.value)
			if (val < 1 || val > 50) {
				return
			}
			this.setState({files: val})
		}

		columnsChange = (e) => {
			let val = parseInt(e.target.value)
			if (val < 1 || val > 50) {
				return
			}
			this.setState({columns: val})
		}

		minesChange = (e) => {
			let val = parseInt(e.target.value)
			let {files, columns} = this.state
			if (val < 1 || val > files * columns) {
				return
			}
			this.setState({mines: val})
		}

    render() {
			const { board, status } = this.props.game
			const { resumeGame, discoverTile } = this.props.gameActions

			return (
				<div className="App">
					<h2>Minesweeper</h2>
					{ status === "game over" && <h3>Game Over!!!</h3>}
					{ status === "win" && <h3>You Win!!!</h3>}
					<Board board={board} discover={discoverTile}/>
					<div className="Params">
						<label className="Label">Files:
							<input type="number" onChange={this.filesChange} value={this.state.files} className="Input"/>
						</label>
						<label className="Label">Columns:
						<input type="number" onChange={this.columnsChange} value={this.state.columns} className="Input"/>
						</label>
						<label className="Label">Mines:
							<input type="number" onChange={this.minesChange} value={this.state.mines} className="Input"/>
						</label>
					</div>
					<div className="Buttons">
						<Button title="New game" action={this.startGame}/>
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