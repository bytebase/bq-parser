WITH Coordinates AS (SELECT [1,2] AS position)
SELECT results FROM Coordinates, UNNEST(Coordinates.position) AS results