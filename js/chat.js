/**
 * 
 */

(function() {
  var socket = new WebSocket("ws://146.169.207.172:8080");
  
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

  
})();