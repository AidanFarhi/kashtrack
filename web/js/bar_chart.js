const barChartContext = document.getElementById('bar-chart').getContext('2d')

async function fetchChartData() {
    try {
        const response = await fetch("/expense_distribution")
        const jsonData = await response.json()
        jsonData.sort((a, b) => b.amount - a.amount);
        const labels = jsonData.map(item => item.category)
        const data = jsonData.map(item => item.amount)
        const backgroundColors = labels.map(() => getRandomColor())
        const chartData = {
            labels: labels,
            datasets: [{
                label: 'Expense Amount',
                data: data,
                backgroundColor: backgroundColors,
                borderColor: backgroundColors.map(color => color.replace('0.75', '1')), // Full opacity for border
                borderWidth: 1
            }]
        }
        const config = {
            type: 'bar',
            data: chartData,
            options: {
                scales: {
                    y: {
                        beginAtZero: true,
                    }
                },
                plugins: {
                    legend: {
                        display: false
                    }
                }
            }
        }
        new Chart(barChartContext, config)
    } catch (error) {
        console.error("Error fetching or processing data:", error)
    }
}

function getRandomColor() {
    const r = Math.floor(Math.random() * 50)
    const g = Math.floor(Math.random() * 50)
    const b = Math.floor(Math.random() * 256)
    const opacity = 0.75
    return `rgba(${r}, ${g}, ${b}, ${opacity})`
}

fetchChartData()
