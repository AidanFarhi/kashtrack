const ctx = document.getElementById('pie-chart')

const data = {
    labels: [
        'Red',
        'Blue',
        'Yellow'
    ],
    datasets: [
        {
            data: [300, 50, 100],
            backgroundColor: [
                'rgb(255, 99, 132)',
                'rgb(54, 162, 235)',
                'rgb(255, 205, 86)'
            ],
            hoverOffset: 4
        }
    ]
}

const config = {
    type: 'pie',
    data: data,
}

new Chart(ctx, config)