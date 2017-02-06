var updateMaster = function(moves) {
    console.log(moves);
    
    var masterBoard = document.getElementsByClassName("nonogram")[0];
    for (var i in moves) {
        console.log(moves[i]);
        ID = "cell_"+moves[i].X+"_"+moves[i].Y;
        if (moves[i].Mark == 1) {
            document.getElementById(ID).setAttribute("class", "filled");
        } else if (moves[i].Mark == 2) {
            document.getElementById(ID).setAttribute("class", "crossed");
        }
    }

}

var requestMoves = function() {
    var request = new XMLHttpRequest();
    request.onreadystatechange = function() {
        if (request.readyState == XMLHttpRequest.DONE) {
            if (request.status >= 200 && request.status < 400) {
                console.log("Request moves successful.");
                updateMaster(JSON.parse(request.responseText));
            } else {
                console.log("Request moves failure.");
            }
        }
    }

    request.open("GET", "http://localhost:8080/moves");
    request.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    console.log("Requesting moves...");
    request.send();
};

/*
for (i = 1; i <= 5; i++) {
    setTimeout(requestMoves, i*3000);
}
*/
setInterval(requestMoves, 3000);