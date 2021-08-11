$(document).ready(function() {
    $('.ucapan').click(function() {
    });
    
    var myAudio = document.getElementById('myAudio');
    var isFinished = false;

    setInterval(function(){ 
        if (myAudio.duration > 0 && !myAudio.paused) {
            isFinished = runningUcapan(myAudio.currentTime, isFinished);
        } else if (!isFinished) {
            $('.ucapan').text("Biar ga sepi, play musik di atas dulu yuk :)")
        }
    }, 1000);
});

function runningUcapan(audioSec) {
    var isFinished = false;
    if (audioSec > 26) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-top");
        $('.ucapan').text("Nih foto-foto Molly :D")
        $('.ucapan-img').show();
    } if (audioSec > 25) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-top");
        $('.ucapan').text("Nih foto-foto Molly :D")
    } else if (audioSec > 20) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-top");
        $('.ucapan').text("Supaya gak bad mood ...")
    } else if (audioSec > 15) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-right");
        $('.ucapan').text("Semoga baik-baik aja yak :)")
    } else if (audioSec > 10) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-left");
        $('.ucapan').text("Apa kabar kamu hari ini?")
    } else if (audioSec > 5) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-bottom");
        $('.ucapan').text("Hai Cutekk!")
    } else if (audioSec > 0) {
        clearAnimation();
        // $('.audio').hide();
        $('.ucapan').addClass("w3-animate-top");
        $('body').addClass("pantone");
        $('.ucapan').text("Nah gitu dong di play musiknya :D")
    }

    return isFinished;
}

function clearAnimation() {
    $('.ucapan').removeClass("w3-animate-bottom");
    $('.ucapan').removeClass("w3-animate-top");
    $('.ucapan').removeClass("w3-animate-left");
    $('.ucapan').removeClass("w3-animate-right");

    $('.ucapan-img').removeClass("w3-animate-bottom");
    $('.ucapan-img').removeClass("w3-animate-top");
    $('.ucapan-img').removeClass("w3-animate-left");
    $('.ucapan-img').removeClass("w3-animate-right");

}
