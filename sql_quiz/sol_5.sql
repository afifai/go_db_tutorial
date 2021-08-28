SELECT 
    t3.genre_name, COUNT(t1.movie_id)
FROM 
    movie AS t1 
INNER JOIN 
    movie_genres AS t2
ON
    t1.movie_id = t2.movie_id
INNER JOIN
    genre AS t3
ON
    t2.genre_id = t3.genre_id
GROUP BY
    t3.genre_id
ORDER BY
    COUNT(t1.movie_id)
LIMIT 5;