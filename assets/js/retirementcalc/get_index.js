$(document).ready(function() {
    $('.form-group').change(function() {
        doCalculate()
    });

    $('.menu-wrapper #btn-quit').click(function() {
        console.log("quit")
        window.location.href = '/';
    });

    function doCalculate() {
        let age, lifespan, income, expenses, inflation

        age = $("#age").val()
        lifespan = $("#lifespan").val()
        income = $("#income").val()
        expenses = $("#expenses").val()
        inflation = $("#inflation").val()


        workingYear = lifespan - age
        yearlyExpenses = expenses * 12
        lifetimeExpenses = 0

        for(i = 0; i < workingYear; i++){
            yearlyExpenses = yearlyExpenses + (inflation / 100 * yearlyExpenses)
            lifetimeExpenses += yearlyExpenses
        }

        estimatedWorkPeriod = Math.ceil(lifetimeExpenses / income)

        yearOfWork = Math.floor(estimatedWorkPeriod / 12) + " year"
        monthOfWork = estimatedWorkPeriod%12 + " month"

        console.log(yearOfWork)
        console.log(monthOfWork)

        result = `Result: You have to work for ${yearOfWork + " " + monthOfWork} to cover your ${lifetimeExpenses.toLocaleString()} IDR lifetime expenses.`

        
        if (parseInt(age) + parseInt(yearOfWork) < parseInt(lifespan)) {
            result += `\n\nRetired at age ${parseInt(age) + parseInt(yearOfWork)}`
        }
        $('#result').text(result)
    }
});