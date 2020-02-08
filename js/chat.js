/**
 * 
 */

(function() {
  var socket = new WebSocket("ws://146.169.207.172:8080");
  
  socket.onopen = function(e) {
    console.log("Success!");
    socket.send("Hello!");
  }
  socket.onmessage = function(e) {
    console.log(e.data);

  }
  socket.onclose = function(e) {
    
  }
  socket.onerror = function(e) {
    
  }

  
})();