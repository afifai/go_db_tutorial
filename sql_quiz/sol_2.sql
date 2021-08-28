SELECT 
    t2.country_id, t3.country_name, COUNT(t1.movie_id)
FROM 
    movie AS t1 
INNER JOIN 
    production_country AS t2
ON
    t1.movie_id = t2.movie_id
INNER JOIN
    country AS t3
ON
    t2.country_id = t3.country_id
GROUP BY
    t3.country_id;