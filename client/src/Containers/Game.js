import React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import * as GameActions from '../Actions/gameActions';


import '../App.css';

class Game extends React.Component {
    state = {}

    render() {
			const { board } = this.props.game

			return (
				<div className="App">
					<h2>Minesweeper</h2>
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