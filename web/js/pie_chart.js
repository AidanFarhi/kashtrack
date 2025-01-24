const ctx = document.getElementById('pie-chart').getContext('2d')

async function fetchChartData() {
    try {
        const response = await fetch("/expense_distribution")
        const jsonData = await response.json()
        const labels = jsonData.map(item => item.category)
        const data = jsonData.map(item => item.amount)
        const backgroundColors = labels.map(() => getRandomColor())
        const chartData = {
            labels: labels,
            datasets: [{
                data: data,
                backgroundColor: backgroundColors,
                hoverOffset: 4
            }]
        }
        const config = {
            type: 'pie',
            data: chartData,
            options: {
                plugins: {
                    legend: {
                        position: 'left'
                    }
                }
            }
        }
        new Chart(ctx, config)
    } catch (error) {
        console.error("Error fetching or processing data:", error)
    }
}

function getRandomColor() {
    const gray = Math.floor(Math.random() * 256)
    return `rgb(${gray}, ${gray}, ${gray})`
}

fetchChartData()