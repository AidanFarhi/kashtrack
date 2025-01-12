const fader = document.getElementById('fader')
const addExpenseButton = document.getElementById('add-expense-btn')
const expenseForm = document.getElementById('expense-form')

const toggleExpenseForm = () => {
    expenseForm.classList.toggle('closed')
    fader.classList.toggle('hidden')
}

fader.addEventListener('click', toggleExpenseForm)
addExpenseButton.addEventListener('click', toggleExpenseForm)