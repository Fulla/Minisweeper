import React from 'react';

import '../App.css';

class Board extends React.Component {
    state = {}

    renderCells = (n, row) => {
			return row.map((cell, index) => {
					return (
							<td key={index} className="Tile" onClick={() => this.props.discover(n, index)}>{cell}</td>
					)
				})
    }
		
		renderRows = (table) => {
			return table.map((row, index) => {
				return (
						<tr key={index}>
							{this.renderCells(index, row)}
						</tr>
				)
			})
		}

    render() {
			const { board } = this.props
            
			return (
				<div className="Table">
					{ board != null && (
						<table>
							<tbody>
								{this.renderRows(board)}
							</tbody>
						</table>
						)
					}
					
				</div>
			)
		}

}

export default Board;