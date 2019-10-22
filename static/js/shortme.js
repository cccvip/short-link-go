function DisplayQR(forWho, content) {
    jQuery("#" + forWho).qrcode({
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
    jQuery.ajax({
        url: "/api/short",
        type: "POST",
        data: "link=" + longURL,
        success: function (data) {
            shortURL = data
        },
        error: function (err) {
            console.log(err);
        }
    }).always(function () {
        if (shortURL === "") {
            return
        }
        var blankList = document.getElementById("shortURLBlankLine");
        if (!blankList.hasChildNodes()) {
            var lineBreak = document.createElement("br");
            document.getElementById("shortURLBlankLine").appendChild(lineBreak);
        }
        // specify the shortened url
        document.getElementById("shortenedURL").innerHTML = shortURL;

        // add shortened qr code
        document.getElementById("shortenedQR").innerHTML = "";
        DisplayQR("shortenedQR", shortURL);
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
    jQuery.ajax({
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
        // clear blank lines
        var blankList = document.getElementById("expandedURLBlankLine");
        if (!blankList.hasChildNodes()) {
            var lineBreak = document.createElement("br");
            document.getElementById("expandedURLBlankLine").appendChild(lineBreak);
        }
        // specify the expanded url
        document.getElementById("expandedURL").innerHTML = longURL;

        // clear expanded qr code
        document.getElementById("expandedQR").innerHTML = "";
        DisplayQR("expandedQR", longURL);
    })
}
