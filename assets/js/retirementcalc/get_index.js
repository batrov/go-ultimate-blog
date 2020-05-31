$(document).ready(function() {
    doCalculate()

    $('.form-group').keyup(function() {
        doCalculate()
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

        workingYear = lifespan - age
        yearlyExpenses = expenses * 12
        lifetimeExpenses = 0

        for (i = 0; i < workingYear; i++) {
            yearlyExpenses = yearlyExpenses + (inflation / 100 * yearlyExpenses)
            lifetimeExpenses += yearlyExpenses
        }

        if (income != 0) {
            estimatedWorkPeriod = Math.ceil(lifetimeExpenses / income)
        } else {
            estimatedWorkPeriod = workingYear * 12
        }

        yearOfWork = Math.floor(estimatedWorkPeriod / 12) + " year"
        monthOfWork = estimatedWorkPeriod % 12 + " month"

        result = `You have to work for ${yearOfWork + " " + monthOfWork} to cover your ${lifetimeExpenses.toLocaleString()} ${currency} lifetime expenses.`
        $('#result').text(result)

        resultRetire = ""
        resultRetireSummary = ""

        if (parseInt(age) + parseInt(yearOfWork) < parseInt(lifespan)) {
            resultRetire = `Retired at age ${parseInt(age) + parseInt(yearOfWork)}`
            $('#result-retire').removeClass("text-danger")
            $('#result-retire').addClass("text-success")

            lifetimeIncome = estimatedWorkPeriod * income

            savings = lifetimeIncome - lifetimeExpenses

            resultRetireSummary = `You will have savings of ${savings.toLocaleString()} ${currency} when you die at age ${lifespan}`

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