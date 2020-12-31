$(function() {
	var cookies = document.cookie;
	var cookiesArray = cookies.split(';');

	console.log(cookies);

	for(var c of cookiesArray){
		var cArray = c.split('=');
		if( cArray[0] == 'auth'){
			document.getElementById("username").textContent = cArray[1];
		}
	}

});
