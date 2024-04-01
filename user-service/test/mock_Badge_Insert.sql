CREATE TYPE badgeme AS ENUM ('month', 'extra');
CREATE TABLE badges(
    id SERIAL PRIMARY KEY,
    badge_name VARCHAR(100),
    badge_date DATE,
    badge_type badgeme,
    picture TEXT
);

CREATE TABLE user_badge(
    user_id UUID REFERENCES users(id),
    badge_id INT REFERENCES badges(id)
);


INSERT INTO badges (badge_name, badge_date, badge_type, picture) VALUES
    ('Star Performer', '2024-02-15', 'month', 'https://example.com/star-performer-badge.png'),
    ('Community Contributor', '2024-01-20', 'extra', 'https://example.com/community-contributor-badge.png'),
    ('Innovator of the Month', '2024-03-05', 'month', 'https://example.com/innovator-badge.png'),
    ('Super User', '2024-02-01', 'extra', 'https://example.com/super-user-badge.png'),
    ('Top Supporter', '2024-01-10', 'month', 'https://example.com/top-supporter-badge.png'),
    ('Tech Guru', '2024-03-12', 'extra', 'https://example.com/tech-guru-badge.png'),
    ('Creative Mind', '2024-02-25', 'month', 'https://example.com/creative-mind-badge.png'),
    ('Team Player', '2024-01-05', 'extra', 'https://example.com/team-player-badge.png'),
    ('Coding Wizard', '2024-03-20', 'month', 'https://example.com/coding-wizard-badge.png'),
    ('Social Butterfly', '2024-02-10', 'extra', 'https://example.com/social-butterfly-badge.png');


INSERT INTO user_badge (user_id, badge_id) VALUES
    ('8a22ae56-d927-11ee-90e4-d8bbc174b998', 1),
    ('8a22ae56-d927-11ee-90e4-d8bbc174b998', 3),
    ('8a22ae56-d927-11ee-90e4-d8bbc174b998', 5),
    ('8a22ae56-d927-11ee-90e4-d8bbc174b998', 7),
    ('8a22ae56-d927-11ee-90e4-d8bbc174b998', 9),
    ('678e9012-e89b-12d3-a456-426614174006', 2),
    ('678e9012-e89b-12d3-a456-426614174006', 4),
    ('678e9012-e89b-12d3-a456-426614174006', 6),
    ('678e9012-e89b-12d3-a456-426614174006', 8),
    ('678e9012-e89b-12d3-a456-426614174006', 1),
    ('678e9012-e89b-12d3-a456-426614174006', 3),
    ('678e9012-e89b-12d3-a456-426614174006', 10);


SELECT 
    id,
    badge_name, 
    badge_date,
    badge_type,
    picture
FROM 
    badges b
JOIN 
    user_badge ub ON b.id = ub.badge_id
WHERE 
    ub.user_id = '8a22ae56-d927-11ee-90e4-d8bbc174b998';

SELECT 
    id,
    badge_name, 
    badge_date,
    badge_type,
    picture
FROM 
    badges b
JOIN 
    user_badge ub ON b.id = ub.badge_id
WHERE 
    ub.user_id = '678e9012-e89b-12d3-a456-426614174006';
