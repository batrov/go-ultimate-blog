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
    if (audioSec > 166) {
        clearAnimation();
        if (!$('.ucapan-final').is( ":visible" )) {
            $('.ucapan-final').show();
        }
        $('.ucapan').text("Happy Birthday Cutekk!!!")

        isFinished = true;
    } else if (audioSec > 160) {
        clearAnimation();
        $('.ucapan-img-13').show();
        $('.ucapan-img-13').attr("display", "flex");
        $('.ucapan-img-12').hide();

    } else if (audioSec > 155) {
        clearAnimation();
        $('.ucapan-img-12').show();
        $('.ucapan').hide();

    } else if (audioSec > 150) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-bottom");
        $('.ucapan').text("Semoga kedepannya kita tetap menjalani kehidupan bersama yaa")
    } else if (audioSec > 147) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-top");
        $('.ucapan').text("Dan makin sayang orang tua :)")
    } else if (audioSec > 144) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-left");
        $('.ucapan').text("Rajin menabung")
    } else if (audioSec > 141) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-bottom");
        $('.ucapan').text("Makin baik")
    } else if (audioSec > 138) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-right");
        $('.ucapan').text("Makin sabar")
    } else if (audioSec > 135) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-top");
        $('.ucapan').text("Cutek makin dewasa")
    } else if (audioSec > 130) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-right");
        $('.ucapan').text("Semoga seiring bertambahnya usia ...")
    } else if (audioSec > 125) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-top");
        $('.ucapan').text("Hari ini hari yang spesial buat cutek ...")
    } else if (audioSec > 119) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-bottom");
        $('.ucapan').text("Aku tetap bersyukur kita bisa melewatinya bersama :)")
    } else if (audioSec > 114) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-left");
        if (!$('.ucapan').is( ":visible" )) {
            $('.ucapan').show();
        }
        $('.ucapan-img-11').hide();

        $('.ucapan').text("Meski kadang kita salah paham ...")
    } else if (audioSec > 109) {
        clearAnimation();
        $('.ucapan-img-11').show();
        $('.ucapan-img-10').hide();

    } else if (audioSec > 107) {
        clearAnimation();
        $('.ucapan-img-10').show();
        $('.ucapan-img-9').hide();

    } else if (audioSec > 105) {
        clearAnimation();
        $('.ucapan-img-9').show();
        $('.ucapan').hide();

    } else if (audioSec > 100) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-bottom");
        if (!$('.ucapan').is( ":visible" )) {
            $('.ucapan').show();
        }
        $('.ucapan-img-8').hide();
        $('.ucapan').text("Kita juga habisin banyak waktu bersama :)")
    } else if (audioSec > 94) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-top");
        $('.ucapan-img-7').hide();
        $('.ucapan-img-8').show();
    } else if (audioSec > 92) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-top");
        $('.ucapan-img-6').hide();
        $('.ucapan-img-7').show();
    } else if (audioSec > 90) {
        clearAnimation();
        $('.ucapan').hide();
        $('.ucapan').addClass("w3-animate-top");
        $('.ucapan-img-6').show();
    } else if (audioSec > 85) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-top");
        if (!$('.ucapan').is( ":visible" )) {
            $('.ucapan').show();
        }
        $('.ucapan').text("Juga ada beberapa momen bersejarah :D")
        $('.ucapan-img-5').hide();

    } else if (audioSec > 81) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-top");
        $('.ucapan-img-4').hide();
        $('.ucapan-img-5').show();

    } else if (audioSec > 79) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-top");
        $('.ucapan-img-3').hide();
        $('.ucapan-img-4').show();

    } else if (audioSec > 77) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-top");
        $('.ucapan-img-2').hide();
        $('.ucapan-img-3').show();

    } else if (audioSec > 75) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-top");
        $('.ucapan').hide();
        $('.ucapan-img-2').show();
    } else if (audioSec > 70) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-left");
        $('.ucapan').text("Sampe patungan langganan Netflix dong")
    } else if (audioSec > 65) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-bottom");
        $('.ucapan').text("Bahas latar belakang pendidikan, keluarga, pekerjaan, kuliah ...")
    } else if (audioSec > 60) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-top");
        if (!$('.ucapan').is( ":visible" )) {
            $('.ucapan').show();
        }
        $('.ucapan').text("Kita bahas apapun yang kita suka ...")
        $('.ucapan-img-1').hide();
    } else if (audioSec > 50) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-top");
        $('.ucapan').hide();
        $('.ucapan-img-1').show();
    } else if (audioSec > 45) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-left");
        $('.ucapan').text("Ini chat pertama kita ...")
    } else if (audioSec > 40) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-right");
        $('.ucapan').text("Waktu itu kita saling match di Bumble :)")
    } else if (audioSec > 35) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-top");
        $('.ucapan').text("Yuk sekarang kita flashback ke awal kita ketemu ...")
        $('.ucapan-img').hide();
    } else if (audioSec > 30) {
        clearAnimation();      
        $('.ucapan').addClass("w3-animate-left");
        if (!$('.ucapan').is( ":visible" )) {
            $('.ucapan').show();
        }
        $('.ucapan').text("Senyum dulu dong biar makin cantik :)")
        $('.ucapan-img').hide();
    } else if (audioSec > 27) {
        clearAnimation();
        $('.ucapan').hide();
        $('.ucapan-img').show();
    } else if (audioSec > 25) {
        clearAnimation();
        $('.ucapan').addClass("w3-animate-bottom");
        $('.ucapan').text("Nih foto Molly :D")
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
        $('.audio').hide();
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
