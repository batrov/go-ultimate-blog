$(document).ready(function() {
    $('.menu-wrapper #btn-submit').click(function() {
        console.log("Do Calculate")
    });

    $('.menu-wrapper #btn-quit').click(function() {
        console.log("quit")
        window.location.href = '/';
    });
});
