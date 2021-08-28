SELECT 
    t1.title, t3.person_name
FROM 
    movie AS t1 
INNER JOIN 
    movie_cast AS t2
ON
    t1.movie_id = t2.movie_id
INNER JOIN
    person AS t3
ON
    t2.person_id = t3.person_id
WHERE
    t1.title like 'brazil';