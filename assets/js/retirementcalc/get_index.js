$(document).ready(function() {
    doCalculate()

    $('.form-group').keyup(function() {
        doCalculate()
    });

    $('.menu-wrapper #btn-advanced').click(function() {
        advanced = $("#advanced").val()
        if (advanced == "0") {
            $("#advanced").val("1")
            $('#btn-advanced').removeClass("btn-warning")
            $('#btn-advanced').addClass("btn-info")
            $('#btn-advanced').text("Advanced Mode: True")

            $('.advanced-wrapper').show()
        } else {
            $("#advanced").val("0")
            $('#btn-advanced').removeClass("btn-info")
            $('#btn-advanced').addClass("btn-warning")
            $('#btn-advanced').text("Advanced Mode: False")

            $('.advanced-wrapper').hide()
        }
        doCalculate()
    });

    $('.menu-wrapper #btn-exit').click(function() {
        window.location.href = '/';
    });

    $(".ot-expenses-wrapper").on('click', '.btn-add-oew', function(e) {
        $('.ot-expenses-wrapper').append($(this).parent().parent()[0].outerHTML)
        $('.btn-del-oew').show()
        doCalculate()
    });

    $(".ot-expenses-wrapper").on('click', '.btn-del-oew', function(e) {
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
            $(".ot-expenses").each(function() {
                lifetimeExpenses += parseInt($(this).val())
            });
        }

        for (i = 0; i < workingYear * 12; i++) {
            workingMonth++
            raisedIncome = raisedIncome * (100 + (raise / 12)) / 100
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

        yearOfWork = Math.floor(workingMonth / 12) + " year"
        monthOfWork = workingMonth % 12 + " month"

        result = `You have to work for ${yearOfWork + " " + monthOfWork} to cover your ${lifetimeExpenses.toLocaleString()} ${currency} lifetime expenses.`
        $('#result').text(result)

        resultRetire = ""
        resultRetireSummary = ""

        if (parseInt(age) + parseInt(yearOfWork) < parseInt(lifespan)) {
            resultRetire = `Retired at age ${parseInt(age) + parseInt(yearOfWork)}`
            $('#result-retire').removeClass("text-danger")
            $('#result-retire').addClass("text-success")

            savings = lifetimeIncome - lifetimeExpenses + savings

            resultRetireSummary = `You will have total income of ${lifetimeIncome.toLocaleString()} ${currency} and save ${savings.toLocaleString()} ${currency} when you die at age ${lifespan}`

        } else {
            resultRetire = `Cannot retire in your entire lifetime`
            $('#result-retire').removeClass("text-success")
            $('#result-retire').addClass("text-danger")

            lifetimeIncome = workingYear * 12 * income

            debt = lifetimeExpenses - lifetimeIncome

            resultRetireSummary = `You will die at age ${lifespan} with debt of ${debt.toLocaleString()} ${currency}`
        }

        $('#result-retire').text(resultRetire)
        $('#result-retire-summary').text(resultRetireSummary)

    }
});

function sendStatistics() {
    let otherExpenses = 0

    $(".ot-expenses").each(function() {
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
    }

    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {}
    };

    xhttp.open("POST", "/retirement-calculator/statistics", true);
    xhttp.setRequestHeader("Content-type", "application/json");
    xhttp.send(JSON.stringify(jsonReq));
}

window.addEventListener("beforeunload", function(e) {
    sendStatistics()
}, false);