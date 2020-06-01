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
        } else {
            $("#advanced").val("0")
            $('#btn-advanced').removeClass("btn-info")
            $('#btn-advanced').addClass("btn-warning")
            $('#btn-advanced').text("Advanced Mode: False")
        }
    });

    $('.menu-wrapper #btn-exit').click(function() {
        console.log("exit")
        window.location.href = '/';
    });

    function doCalculate() {
        let age, lifespan, income, expenses, inflation

        age = $("#age").val()
        lifespan = $("#lifespan").val()
        income = $("#income").val()
        expenses = $("#expenses").val()
        inflation = $("#inflation").val()
        currency = $("#currency").val()
        raise = $("#raise").val()

        workingYear = lifespan - age
        yearlyExpenses = expenses * 12
        raisedIncome = income
        lifetimeExpenses = 0
        lifetimeIncome = 0
        workingMonth = 0

        for (i = 0; i < workingYear; i++) {
            yearlyExpenses = yearlyExpenses + (inflation / 100 * yearlyExpenses)
            lifetimeExpenses += yearlyExpenses
        }

        for (i = 0; i < workingYear * 12; i++) {
            workingMonth++
            raisedIncome = raisedIncome * (100 + (raise / 12)) / 100
            lifetimeIncome += raisedIncome

            if (lifetimeIncome >= lifetimeExpenses) {
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

            savings = lifetimeIncome - lifetimeExpenses

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