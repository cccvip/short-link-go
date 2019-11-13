function DisplayQR(forWho, content) {
    $("#" + forWho).qrcode({
        text: content,
        width: 128,
        height: 128
    });
}

function Short(longURLID) {
    var longURL = document.getElementById(longURLID).value;
    longURL = longURL.trim();
    if (longURL === "") {
        alert("Input text is empty. :-)");
        return
    }
    var shortURL = "";
    $.ajax({
        url: "/api/short",
        type: "POST",
        data: "link=" +  longURL,
        success: function (data) {
            shortURL = data
        },
        error: function (err) {
            console.log(err);
        }
    }).always(function () {
        $( "#shortURLBlankLine" ).html(shortURL)
        // add shortened qr code
        //document.getElementById("shortenedQR").innerHTML = "";
        //DisplayQR("shortenedQR", shortURL);
    })
}

function Expand(shortURLID) {
    var shortURL = document.getElementById(shortURLID).value;
    shortURL = shortURL.trim();
    if (shortURL === "") {
        alert("Input text is empty. :-)");
        return
    }
    var longURL = "";
    $.ajax({
        url: "/api/info?link=" + shortURL,
        type: "GET",
        success: function (data) {
            longURL = data;
        },
        error: function (err) {
            console.log(err);
        }
    }).always(function () {
        if (longURL === "") {
            return
        }
        $( "#expandedURLBlankLine" ).html(longURL)
        // clear expanded qr code
        //document.getElementById("expandedQR").innerHTML = "";
        //DisplayQR("expandedQR", longURL);
    })
}
