$(document).ready(function() {
    $('.menu-wrapper #btn-survey').click(function() {
        window.location.href = '/goquiz/survey';
    });

    $('.menu-wrapper #btn-quiz').click(function() {
        window.location.href = '/goquiz/quiz';
    });

    $('.menu-wrapper #btn-quit').click(function() {
        console.log("quit")
        window.location.href = '/';
    });

    $('.menu-wrapper #btn-reset').click(function() {
        console.log("reset")
        var xhttp = new XMLHttpRequest();
        xhttp.onreadystatechange = function() {
            if (this.readyState == 4 && this.status == 200) {
                $('#reset-result').html(this.responseText)
            }
        };

        xhttp.open("POST", "/goquiz/reset-answer", true);
        xhttp.send();
    });

});
