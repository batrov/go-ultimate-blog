var allowedNumberInput = ["1", "2", "3", "4", "5", "6", "7", "8", "9", "0", ".", "Backspace", "ArrowLeft", "ArrowUp", "ArrowRight", "ArrowDown", "Delete"];

$(document).ready(function () {
    doCalculate()

    $('.form-group').keyup(function () {
        doCalculate()
    });

    $('.menu-wrapper #btn-advanced').click(function () {
        advanced = $("#advanced").val()
        if (advanced == "0") {
            $("#advanced").val("1")
            $('#btn-advanced').removeClass("btn-warning")
            $('#btn-advanced').addClass("btn-info")
            $('#btn-advanced').text("Simple Mode")

            $('.advanced-wrapper').show()
        } else {
            $("#advanced").val("0")
            $('#btn-advanced').removeClass("btn-info")
            $('#btn-advanced').addClass("btn-warning")
            $('#btn-advanced').text("Advanced Mode")

            $('.advanced-wrapper').hide()
        }
        doCalculate()
    });

    $('.menu-wrapper #btn-exit').click(function () {
        window.location.href = '/';
    });

    $(".ot-expenses-wrapper").on('click', '.btn-add-oew', function (e) {
        $('.ot-expenses-wrapper').append($(this).parent().parent()[0].outerHTML)
        $('.btn-del-oew').show()
        doCalculate()
    });

    $(".ot-expenses-wrapper").on('click', '.btn-del-oew', function (e) {
        if ($('.ot-expenses-wrapper').children().length > 1) {
            $(this).parent().parent().remove()
            if ($('.ot-expenses-wrapper').children().length > 1) {
                $('.btn-del-oew').show()
            } else {
                $('.btn-del-oew').hide()
            }
        }
        doCalculate()
    });

    function doCalculate() {
        let age, lifespan, income, expenses, inflation

        age = parseInt($("#age").val()) || 0
        lifespan = parseInt($("#lifespan").val()) || 0
        income = parseFloat($("#income").val()) || 0
        expenses = parseFloat($("#expenses").val()) || 0
        inflation = parseFloat($("#inflation").val()) || 0
        currency = $("#currency").val()
        raise = parseFloat($("#raise").val()) || 0
        advancedMode = $("#advanced").val() == "1"
        investments = parseFloat($("#investments").val()) || 0
        returns = parseFloat($("#returns").val()) || 0
        savings = parseFloat($("#savings").val()) || 0
        savingsInterests = parseFloat($("#savings-interests").val()) || 0

        workingYear = lifespan - age
        yearlyExpenses = expenses * 12
        raisedIncome = income
        lifetimeExpenses = 0
        lifetimeIncome = 0
        workingMonth = 0
        routineInvestment = investments
        savingsInterestsSum = 0

        for (i = 0; i < workingYear; i++) {
            yearlyExpenses = yearlyExpenses + (inflation / 100 * yearlyExpenses)
            lifetimeExpenses += yearlyExpenses
        }

        if (advancedMode) {
            $(".ot-expenses").each(function () {
                lifetimeExpenses += parseInt($(this).val())
            });
        }

        for (i = 0; i < workingYear * 12; i++) {
            workingMonth++
            if (workingMonth % 13 == 0) { // raise only at the start of next year
                raisedIncome = raisedIncome * (100 + raise) / 100
            }
            savingsInterestsSum = savingsInterests / 100 * savings
            savings += savingsInterestsSum
            lifetimeIncome += raisedIncome

            if (advancedMode) {
                investmentsProfits = returns / 12 / 100 * investments
                investments += investmentsProfits
                investments += routineInvestment
                lifetimeIncome += investmentsProfits
            }

            if (lifetimeIncome + savings >= lifetimeExpenses) {
                break
            }
        }

        yearOfWork = Math.floor(workingMonth / 12) + " year(s)"
        monthOfWork = workingMonth % 12 + " month(s)"

        $('#result-work-time').text(yearOfWork + " " + monthOfWork)
        $('#result-lifetime-expenses').text(lifetimeExpenses.toLocaleString() + " " + currency)

        resultRetire = ""
        resultRetireSummary = ""

        if (parseInt(age) + parseInt(yearOfWork) < parseInt(lifespan)) {
            resultRetire = parseInt(age) + parseInt(yearOfWork)
            $('#result-retired-age').removeClass("text-danger")
            $('#result-retired-age').addClass("text-success")
            $('.result-table').addClass("result-table-positive")
            $('.result-table').removeClass("result-table-negative")
            $('.result-cannot-retire').hide()
            $('.result-can-retire').show()

            savings = lifetimeIncome - lifetimeExpenses + savings

            resultRetireSummary = `You will have total income of ${lifetimeIncome.toLocaleString()} ${currency} and save ${savings.toLocaleString()} ${currency} when you die at age ${lifespan}`
            $('#result-lifetime-savings').text(savings.toLocaleString() + " " + currency)
        } else {
            resultRetire = `Cannot retire in your entire lifetime`
            $('#result-retired-age').removeClass("text-success")
            $('#result-retired-age').addClass("text-danger")
            $('.result-table').removeClass("result-table-positive")
            $('.result-table').addClass("result-table-negative")
            $('.result-cannot-retire').show()
            $('.result-can-retire').hide()

            lifetimeIncome = workingYear * 12 * income

            debt = lifetimeExpenses - lifetimeIncome

            resultRetireSummary = `You will die at age ${lifespan} with debt of ${debt.toLocaleString()} ${currency}`

            $('#result-remaining-debts').text(debt.toLocaleString() + " " + currency)
        }

        $('#result-retired-age').text(resultRetire)
        $('#result-lifetime-income').text(lifetimeIncome.toLocaleString() + " " + currency)

    }

    $(".custom-form-input").keydown(function (e) {
        var isAllowed = allowedNumberInput.includes(e.key)

        if (e.key == ".") {
            isAllowed = !$(this).val().includes(e.key)
        }

        if (!isAllowed) {
            e.preventDefault();
        }
    });

    $(".custom-form-input").keyup(function (e) {
        var key = $(this).attr('id').substring(5);
        $(this).val($(this).val().replace(/,/g, ""));

        if ($(this).val().length > 1) {
            while ($(this).val()[0] == "0") {
                $(this).val($(this).val().substring(1));
            }
        }

        if ($(this).val().length < 1) {
            $(this).val("0");
        }

        var floatVal = parseFloat($(this).val())

        $(`#${key}`).val(floatVal);
        $(this).val(addCommas($(this).val()));
    });
});

function sendStatistics() {
    let otherExpenses = 0

    $(".ot-expenses").each(function () {
        otherExpenses += parseInt($(this).val())
    });

    jsonReq = {
        age: parseInt($("#age").val()),
        lifespan: parseInt($("#lifespan").val()),
        income: parseFloat($("#income").val()),
        expenses: parseFloat($("#expenses").val()),
        inflation: parseFloat($("#inflation").val()),
        currency: $("#currency").val(),
        raise: parseFloat($("#raise").val()),
        advanced_mode: $("#advanced").val() == "1",
        investments: parseFloat($("#investments").val()),
        returns: parseFloat($("#returns").val()),
        other_expenses: parseInt(otherExpenses),
        current_savings: parseFloat($("#savings").val()),
    }

    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function () {
        if (this.readyState == 4 && this.status == 200) { }
    };

    xhttp.open("POST", "/fire-calculator/statistics", true);
    xhttp.setRequestHeader("Content-type", "application/json");
    xhttp.send(JSON.stringify(jsonReq));
}

window.addEventListener("beforeunload", function (e) {
    sendStatistics()
}, false);

function addCommas(x) {
    var parts = x.toString().split(".");
    parts[0] = parts[0].replace(/\B(?=(\d{3})+(?!\d))/g, ",");
    return parts.join(".");
}

function formatNumber() { // TODO: fix this
    var input = document.getElementsByClassName("ot-expenses");
    var value = input.value.replace(/,/g, "");
    var formattedValue = Number(value).toLocaleString();
    input.value = formattedValue;
}