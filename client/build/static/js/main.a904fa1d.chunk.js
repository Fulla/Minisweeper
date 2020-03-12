(this.webpackJsonpclient=this.webpackJsonpclient||[]).push([[0],{10:function(e,t,n){},18:function(e,t,n){e.exports=n(30)},29:function(e,t,n){},30:function(e,t,n){"use strict";n.r(t);var a={};n.r(a),n.d(a,"startGame",(function(){return C})),n.d(a,"resumeGame",(function(){return N})),n.d(a,"discoverTile",(function(){return I}));var r=n(0),o=n.n(r),s=n(3),c=n.n(s),i=n(7),l=n(1),u=n(16),m=n(2),f=n(17),d={board:null,files:0,columns:0,status:"off"};function h(e,t,n){return e>=n.files||t>=n.columns||e<0||t<0}var v=Object(l.c)({game:function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:d,t=arguments.length>1?arguments[1]:void 0,n=[];switch(t.type){case"START_BOARD":var a=Object(f.a)(Array(t.files)).map((function(e){return Array(t.columns).fill("")}));return Object(m.a)({},e,{board:a,files:t.files,columns:t.columns,status:t.state});case"SET_SAFEPOINTS":n=[];for(var r=0;r<e.board.length;r++)n[r]=e.board[r].slice();for(var o in t.safepoints){var s=t.safepoints[o],c=!0,i=!1,l=void 0;try{for(var u,v=s[Symbol.iterator]();!(c=(u=v.next()).done);c=!0){var b=u.value;h(b.file,b.column,e)||(n[b.file][b.column]=o)}}catch(j){i=!0,l=j}finally{try{c||null==v.return||v.return()}finally{if(i)throw l}}}return Object(m.a)({},e,{board:n});case"SET_MINES":n=[];for(r=0;r<e.board.length;r++)n[r]=e.board[r].slice();var p=!0,g=!1,E=void 0;try{for(var y,O=t.mines[Symbol.iterator]();!(p=(y=O.next()).done);p=!0){var T=y.value;h(T.file,T.column,e)||(n[T.file][T.column]="*")}}catch(j){g=!0,E=j}finally{try{p||null==O.return||O.return()}finally{if(g)throw E}}return Object(m.a)({},e,{board:n});case"SET_ACTIVATED":var S=t.activated;if(null==S)return e;if(h(S.file,S.column,e))return e;n=[];for(r=0;r<e.board.length;r++)n[r]=e.board[r].slice();return n[S.file][S.column]="X",Object(m.a)({},e,{board:n});case"SET_STATUS":return Object(m.a)({},e,{status:t.state});default:return e}}});Boolean("localhost"===window.location.hostname||"[::1]"===window.location.hostname||window.location.hostname.match(/^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/));n(29);var b=n(4),p=n(5),g=n(8),E=n(6),y=n(9);function O(e,t){return t.headers=new Headers({"Content-Type":"application/json"}),fetch("".concat("/api","/").concat(e),t).then((function(e){return e.json()}))}function T(e,t,n){return{type:"START_BOARD",files:e,columns:t,state:n}}function S(e){return{type:"SET_SAFEPOINTS",safepoints:e}}function j(e){return{type:"SET_MINES",mines:e}}function w(e){return{type:"SET_ACTIVATED",activated:e}}function A(e){return{type:"SET_STATUS",state:e}}function C(e,t,n){return function(a){return(r=e,o=t,s=n,O("new",{method:"POST",body:JSON.stringify({files:r,columns:o,mines:s})}).then((function(e){return e.data}))).then((function(e){console.log(e),a(T(e.files,e.columns,e.state))}));var r,o,s}}function N(){return function(e){return O("resume",{method:"GET"}).then((function(e){return e.data})).then((function(t){console.log(t),e(T(t.files,t.columns,t.state)),e(S(t.safePoints)),e(j(t.mines)),e(w(t.activatedMine)),e(A(t.state))}))}}function I(e,t){return function(n){return(a=e,r=t,O("discover",{method:"POST",body:JSON.stringify({file:a,column:r})}).then((function(e){return e.data}))).then((function(e){console.log(e),n(S(e.safePoints)),n(j(e.mines)),n(w(e.activatedMine)),n(A(e.state))}));var a,r}}n(10);var k=function(e){function t(){var e,n;Object(b.a)(this,t);for(var a=arguments.length,r=new Array(a),s=0;s<a;s++)r[s]=arguments[s];return(n=Object(g.a)(this,(e=Object(E.a)(t)).call.apply(e,[this].concat(r)))).state={},n.renderCells=function(e,t){return t.map((function(t,a){return o.a.createElement("td",{key:a,className:"Tile",onClick:function(){return n.props.discover(e,a)}},t)}))},n.renderRows=function(e){return e.map((function(e,t){return o.a.createElement("tr",{key:t},n.renderCells(t,e))}))},n}return Object(y.a)(t,e),Object(p.a)(t,[{key:"render",value:function(){var e=this.props.board;return o.a.createElement("div",{className:"Container"},null!=e&&o.a.createElement("table",{className:"Table"},o.a.createElement("tbody",null,this.renderRows(e))))}}]),t}(o.a.Component);var _,G=function(e){return o.a.createElement("div",{className:"Button",onClick:e.action},e.title)},P=function(e){function t(){var e,n;Object(b.a)(this,t);for(var a=arguments.length,r=new Array(a),o=0;o<a;o++)r[o]=arguments[o];return(n=Object(g.a)(this,(e=Object(E.a)(t)).call.apply(e,[this].concat(r)))).state={files:5,columns:5,mines:5},n.startGame=function(){var e=n.state,t=e.files,a=e.columns,r=e.mines;n.props.gameActions.startGame(t,a,r)},n.filesChange=function(e){var t=parseInt(e.target.value);t<1||t>50||n.setState({files:t})},n.columnsChange=function(e){var t=parseInt(e.target.value);t<1||t>50||n.setState({columns:t})},n.minesChange=function(e){var t=parseInt(e.target.value),a=n.state,r=a.files,o=a.columns;t<1||t>r*o||n.setState({mines:t})},n}return Object(y.a)(t,e),Object(p.a)(t,[{key:"render",value:function(){var e=this.props.game,t=e.board,n=e.status,a=this.props.gameActions,r=a.resumeGame,s=a.discoverTile;return o.a.createElement("div",{className:"App"},o.a.createElement("h2",null,"Minesweeper"),"game over"===n&&o.a.createElement("h3",null,"Game Over!!!"),"win"===n&&o.a.createElement("h3",null,"You Win!!!"),o.a.createElement(k,{board:t,discover:s}),o.a.createElement("div",{className:"Params"},o.a.createElement("label",{className:"Label"},"Files:",o.a.createElement("input",{type:"number",onChange:this.filesChange,value:this.state.files,className:"Input"})),o.a.createElement("label",{className:"Label"},"Columns:",o.a.createElement("input",{type:"number",onChange:this.columnsChange,value:this.state.columns,className:"Input"})),o.a.createElement("label",{className:"Label"},"Mines:",o.a.createElement("input",{type:"number",onChange:this.minesChange,value:this.state.mines,className:"Input"}))),o.a.createElement("div",{className:"Buttons"},o.a.createElement(G,{title:"New game",action:this.startGame}),o.a.createElement(G,{title:"Resume game",action:r})))}}]),t}(o.a.Component),R=Object(i.b)((function(e){return{game:e.game}}),(function(e){return{gameActions:Object(l.b)(a,e)}}))(P),B=Object(l.d)(v,_,Object(l.a)(u.a));c.a.render(o.a.createElement(i.a,{store:B},o.a.createElement(R,null)),document.getElementById("root")),"serviceWorker"in navigator&&navigator.serviceWorker.ready.then((function(e){e.unregister()})).catch((function(e){console.error(e.message)}))}},[[18,1,2]]]);
//# sourceMappingURL=main.a904fa1d.chunk.js.map