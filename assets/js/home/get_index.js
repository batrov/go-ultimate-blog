$(document).ready(function() {
    $('.dropdown-menu #go-quiz').click(function() {
        window.location.href = '/goquiz';
    });

    $('.nav-item #about').click(function() {
        window.location.href = '/about';
    });

    $('.dropdown-menu #retirement-calculator').click(function() {
        window.location.href = '/retirement-calculator';
    });
});