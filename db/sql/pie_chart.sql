SELECT category, ROUND(SUM(amount), 2)
FROM expense
WHERE SUBSTRING(timestamp, 1, 7) = STRFTIME('%Y-%m', DATE('now'))
GROUP BY category