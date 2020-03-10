import { BASE_URL } from './constants';

function makeRequest(url, request) {
  request.headers = new Headers({
    'Content-Type': 'application/json'
  })
  return fetch(`${BASE_URL}/${url}`, request)
  .then(
      (response) => response.json()
  )
}

export function fetchGameState() {
  let request = {
    method: 'GET'
  }
  return makeRequest('resume', request)
  .then(
    (res) => res.data
  )
}


export function fetchStartGame(files, columns, mines) {
	let request = {
		method: 'POST',
		body: JSON.stringify({
			files,
			columns,
			mines
    })
	}
	return makeRequest('new', request)
	.then(
		(res) => res.data
	)
}

export function fetchDiscoverTile(file, column) {
	let request = {
		method: 'POST',
		body: JSON.stringify({
			file,
			column,
    })
	}
	return makeRequest('discover', request)
	.then(
		(res) => res.data
	)
}