"use strict"

const ENTER_KEYKODE = 13;

const searchButton = document.getElementById("searchBtn");
const nounDisplay = document.getElementById("nounDisplay");
const nounTextField = document.getElementById("noun");
const genderArticle = document.getElementById("gender");
const BACKEND_ENDPOINT_URL = "http://localhost:8080/article";
const BACKEND_ENDPOINT_METHOD = "POST";
const xmlHTTP = new XMLHttpRequest();

const umlautAeButton = document.getElementById("ae");
const umlautOeButton = document.getElementById("oe");
const umlautIuButton = document.getElementById("iu");

document.body.onkeydown = function(event) {
	let keyCode = event.keyCode;

	if(keyCode === ENTER_KEYKODE) {
		searchButton.click();
	}
}

umlautAeButton.onclick = () => nounTextField.value += "ä";
umlautOeButton.onclick = () => nounTextField.value += "ö";
umlautIuButton.onclick = () => nounTextField.value += "ü";

searchButton.onclick = function () {
	let nounUserText = nounTextField.value;
	nounUserText = nounUserText.trim();
	const response = JSON.parse(sendBackendRequest(nounUserText));
	if(!wasRequestSuccessful()) {
		alert("Unknown noun");
	} else {	
		updateNounDisplay(nounUserText);	
		updateGenderDisplay(response.Article);
	}
}

function sendBackendRequest(noun) {
	xmlHTTP.open(BACKEND_ENDPOINT_METHOD, BACKEND_ENDPOINT_URL, false);
	xmlHTTP.send(JSON.stringify({"RawText": noun}));
	return xmlHTTP.responseText;
}

function wasRequestSuccessful() {
	return xmlHTTP.status === 200;
}

function updateNounDisplay(noun) {
	nounDisplay.innerHTML = `Noun ${noun}`;
}

function updateGenderDisplay(article) {
	genderArticle.innerHTML = `Gender ${article.toUpperCase()}`
} 