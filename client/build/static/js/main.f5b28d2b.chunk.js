(this.webpackJsonpclient=this.webpackJsonpclient||[]).push([[0],{17:function(t,e,n){t.exports=n(30)},28:function(t,e,n){},29:function(t,e,n){},30:function(t,e,n){"use strict";n.r(e);var r={};n.r(r),n.d(r,"getGameState",(function(){return S})),n.d(r,"startGame",(function(){return A}));var a=n(0),o=n.n(a),c=n(2),i=n.n(c),s=n(3),u=n(1),l=n(10),f=n(7),m=n(14),h={board:null,files:0,columns:0,state:"off"};var d=Object(u.c)({game:function(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:h,e=arguments.length>1?arguments[1]:void 0;switch(e.type){case"START_BOARD":var n=Object(m.a)(Array(e.files)).map((function(t){return Array(e.columns).fill(" ")}));return Object(f.a)({},t,{board:n,files:e.files,columns:e.columns,state:e.state});case"SET_SAFEPOINTS":var r=t.board;for(var a in e.safepoints){var o=e.safepoints[a],c=!0,i=!1,s=void 0;try{for(var u,l=o[Symbol.iterator]();!(c=(u=l.next()).done);c=!0){var d=u.value;d.x>=t.files||d.y>=t.columns||d.x<0||d.y<0||(r[d.x][d.y]=a)}}catch(p){i=!0,s=p}finally{try{c||null==l.return||l.return()}finally{if(i)throw s}}}return Object(f.a)({},t,{board:n});default:return t}}});Boolean("localhost"===window.location.hostname||"[::1]"===window.location.hostname||window.location.hostname.match(/^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/));n(28);var p=n(11),b=n(12),v=n(15),y=n(13),O=n(16);function g(t,e){return e.headers=new Headers({"Content-Type":"application/json"}),fetch("".concat("http://0.0.0.0:3000/api","/").concat(t),e).then((function(t){return t.json()}))}function j(t,e,n){return{type:"START_BOARD",files:t,columns:e,state:n}}function w(t){return{type:"SET_SAFEPOINTS",safepoints:t}}function S(){return function(t,e){return g("resume",{method:"GET"}).then((function(t){return t.data})).then((function(e){t(j(e.files,e.columns,e.state)),t(w(e.safepoints))}))}}function A(){return function(t,e){return g("new",{method:"POST",body:JSON.stringify({files:10,columns:10})}).then((function(t){return t.data})).then((function(e){t(j(e.files,e.columns,e.state)),t(w(e.safepoints))}))}}n(29);var T,E=function(t){function e(){var t,n;Object(p.a)(this,e);for(var r=arguments.length,a=new Array(r),o=0;o<r;o++)a[o]=arguments[o];return(n=Object(v.a)(this,(t=Object(y.a)(e)).call.apply(t,[this].concat(a)))).state={},n}return Object(O.a)(e,t),Object(b.a)(e,[{key:"render",value:function(){this.props.game.board;return o.a.createElement("div",{className:"App"},o.a.createElement("h2",null,"Minesweeper"))}}]),e}(o.a.Component),k=Object(s.b)((function(t){return{game:t.game}}),(function(t){return{gameActions:Object(u.b)(r,t)}}))(E),x=Object(u.d)(d,T,Object(u.a)(l.a));i.a.render(o.a.createElement(s.a,{store:x},o.a.createElement(k,null)),document.getElementById("root")),"serviceWorker"in navigator&&navigator.serviceWorker.ready.then((function(t){t.unregister()})).catch((function(t){console.error(t.message)}))}},[[17,1,2]]]);
//# sourceMappingURL=main.f5b28d2b.chunk.js.map