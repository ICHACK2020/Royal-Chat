"use strict";
 
var socket;
var id;

function sendMessage() {
  let chatArea = document.getElementById("chatArea");
  socket.send(id + ": " + chatArea.value);
  chatArea.value = "";
}

function addMyMessage(message) {
  let messageBox = document.getElementById("messageBox");
  messageBox.insertAdjacentHTML("beforeend", '<div class="panel message myMessage"><div class="panel-body"></div></div>');
  printMessage(message);
}

function addOtherMessage(message) {
  let messageBox = document.getElementById("messageBox");
  messageBox.insertAdjacentHTML("beforeend", '<div class="panel message otherMessage"><div class="panel-body"></div></div>');
  printMessage(message);
}

function printMessage(message) {
  var node = document.createTextNode(message);
  document.getElementById("messageBox").lastChild.firstChild.appendChild(node);
  window.scrollTo(0, document.body.scrollHeight);
}

function initializeSocket() {
  var socket = new WebSocket("ws://146.169.207.172:8080/talk/");
  
  socket.onopen = function(e) {
    console.log("Success!");
    var obj = {ID: 12345678, Msg: "Test", Troll: 5.0, Relevance: 10.0};
    socket.send(JSON.stringify(obj));
  }
  socket.onmessage = function(e) {
    console.log(JSON.parse(e.data));

  }
  socket.onclose = function(e) {
    
  }
  socket.onerror = function(e) {
    
  }
  return socket;
}

document.onkeydown = function (event) {
  if (event.key === "Enter") {
    sendMessage();
    return false;
  }
};

function handleResize() {
  $(".myMessage").css("left", (window.innerWidth >= 390 ? window.innerWidth - 390 : 0) + "px");
}

window.onresize = function (event) {
  handleResize();
};

socket = initializeSocket();
handleResize();
