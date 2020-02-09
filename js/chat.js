"use strict";

var uid;
var first = true;
//var socket = new WebSocket("ws://146.169.207.172:8080/talk/");

//let socket = new WebSocket("ws://146.169.207.172:8080/talk/")
var url = document.URL.split("/").pop()
var convId = url.slice(0, url.length - 1)
let socket = new WebSocket("ws://146.169.207.172:8080/ws/" + convId)

socket.onopen = function(e) {
    console.log("Success!");
}
socket.onmessage = function(e) {
  if (first) {
    uid = parseInt(e.data);
    console.log(e);
    first = false;
    //Stop waiting for users
    document.getElementById("loadingBox").style.visibility = "hidden";
  } else {
    var obj = JSON.parse(e.data);
    if (uid == obj.UID) {
      addMyMessage(obj.Msg);
    } else {
      addOtherMessage(obj.Msg);
    }
    console.log(JSON.parse(e.data));
  }
}
socket.onclose = function(e) {
<<<<<<< HEAD
  console.log("closing");
  window.location.replace("http://146.169.207.172:8080");
=======
    console.log("closing")
    window.location.replace("http://146.169.207.172:8080")
>>>>>>> branch 'master' of https://github.com/ICHACK2020/ichack2020.git
}
socket.onerror = function(e) {
  console.log(e);
}

class incomingMsg {
  constructor(Msg) {
    this.UID = uid;
    this.ConvID = convId;
    this.Msg = Msg;
  }
}

function sendMessage() {
  let chatArea = document.getElementById("chatArea");
  if (chatArea.value == "") {
      alert("Enter a message");
      return;
  }
  socket.send(JSON.stringify(new incomingMsg(chatArea.value)));
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

handleResize();
