CREATE TABLE activitys (
    id SERIAL PRIMARY KEY,
    active_day DATE,
    active_score INTEGER DEFAULT 1,
    user_id UUID REFERENCES users(id)
);


INSERT INTO activitys (active_day, active_score, user_id) VALUES
    ('2024-01-05', 40, '8a22ae56-d927-11ee-90e4-d8bbc174b998'),
    ('2024-01-10', 50, '8a22ae56-d927-11ee-90e4-d8bbc174b998'),
    ('2024-01-15', 30, '8a22ae56-d927-11ee-90e4-d8bbc174b998'),
    ('2024-01-20', 40, '8a22ae56-d927-11ee-90e4-d8bbc174b998'),
    ('2024-01-25', 30, '8a22ae56-d927-11ee-90e4-d8bbc174b998'),
    ('2024-02-01', 50, '8a22ae56-d927-11ee-90e4-d8bbc174b998'),
    ('2024-02-05', 20, '8a22ae56-d927-11ee-90e4-d8bbc174b998'),
    ('2024-02-10', 30, '8a22ae56-d927-11ee-90e4-d8bbc174b998'),
    ('2024-02-15', 40, '8a22ae56-d927-11ee-90e4-d8bbc174b998'),
    ('2024-02-20', 50, '8a22ae56-d927-11ee-90e4-d8bbc174b998'),
    ('2024-02-25', 30, '8a22ae56-d927-11ee-90e4-d8bbc174b998'),
    ('2024-03-01', 40, '8a22ae56-d927-11ee-90e4-d8bbc174b998'),
    ('2024-03-05', 30, '8a22ae56-d927-11ee-90e4-d8bbc174b998'),
    ('2024-03-10', 50, '8a22ae56-d927-11ee-90e4-d8bbc174b998'),
    ('2024-03-15', 20, '8a22ae56-d927-11ee-90e4-d8bbc174b998'),
    ('2024-03-20', 30, '8a22ae56-d927-11ee-90e4-d8bbc174b998'),
    ('2024-03-25', 40, '8a22ae56-d927-11ee-90e4-d8bbc174b998'),
    ('2024-03-30', 50, '8a22ae56-d927-11ee-90e4-d8bbc174b998'),
    ('2024-04-05', 30, '8a22ae56-d927-11ee-90e4-d8bbc174b998'),
    ('2024-04-10', 40, '8a22ae56-d927-11ee-90e4-d8bbc174b998');


SELECT 
    id,
    active_day, 
    active_score,
FROM 
    activitys a
WHERE 
    a.user_id = '8a22ae56-d927-11ee-90e4-d8bbc174b998';