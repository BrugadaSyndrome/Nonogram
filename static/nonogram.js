var requestMoves = function() {
    var request = new XMLHttpRequest();
    request.onreadystatechange = function() {
        if (request.readyState == XMLHttpRequest.DONE) {
            if (request.status >= 200 && request.status < 400) {
                console.log("Request moves successful.");
            } else {
                console.log("Request moves failure.");
            }
        }
    }

    request.open("GET", "http://localhost:8080/moves");
    console.log("Requesting moves...");
    request.send();
};

setInterval(requestMoves, 3000);