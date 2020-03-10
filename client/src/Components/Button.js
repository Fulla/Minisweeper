import React from 'react';
import '../App.css';

function Button(props) {
  return <div className="Button" onClick={props.action}>{props.title}</div>;
}

export default Button;