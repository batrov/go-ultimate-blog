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
});
