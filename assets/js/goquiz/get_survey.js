$(document).ready(function() {
    $('.survey-menu-wrapper #btn-submit').click(function() {
        postSurvey()
    });

    $('.survey-menu-wrapper #btn-exit').click(function() {
        window.location.href = '/goquiz/';
    });
});


function postSurvey() {

    jsonReq = {
        survey_id: parseInt($('#survey-id').val()),
        answer: $('#answer').val(),
    }

    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            document.getElementById("result").innerHTML = this.responseText;
        }
    };

    xhttp.open("POST", "/goquiz/survey", true);
    xhttp.setRequestHeader("Content-type", "application/json");
    xhttp.send(JSON.stringify(jsonReq));
}