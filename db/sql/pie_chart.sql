SELECT category, SUM(amount) AS total
FROM expense
WHERE SUBSTRING(timestamp, 1, 7) = STRFTIME('%Y-%m', DATE('now'))
GROUP BY category
ORDER BY total DESC;