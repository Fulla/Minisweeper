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


export function fetchStartGame() {
	let request = {
		method: 'POST',
		body: JSON.stringify({
      files: 10,
      columns: 10
    })
	}
	return makeRequest('new', request)
	.then(
		(res) => res.data
	)
}

