CREATE TABLE socials(
    id SERIAL PRIMARY KEY,
    social_name VARCHAR(65),
    social_icon TEXT
);

CREATE TABLE user_social(
    user_id UUID REFERENCES users(id),
    social_id INT REFERENCES socials(id),
    name_of_social VARCHAR(100),
    link_of_social VARCHAR(512)
);


INSERT INTO socials (social_name, social_icon) VALUES
    ('Facebook', 'https://example.com/facebook-icon.png'),
    ('Twitter', 'https://example.com/twitter-icon.png'),
    ('Instagram', 'https://example.com/instagram-icon.png'),
    ('LinkedIn', 'https://example.com/linkedin-icon.png'),
    ('TikTok', 'https://example.com/tiktok-icon.png'),
    ('Telegram', 'https://example.com/telegram-icon.png');


INSERT INTO user_social (user_id, social_id, name_of_social, link_of_social) VALUES
    ('8a22ae56-d927-11ee-90e4-d8bbc174b998', 1, 'John Doe Facebook', 'https://facebook.com/johndoe'),
    ('8a22ae56-d927-11ee-90e4-d8bbc174b998', 2, 'John Doe Twitter', 'https://twitter.com/johndoe'),
    ('8a22ae56-d927-11ee-90e4-d8bbc174b998', 3, 'John Doe Instagram', 'https://instagram.com/johndoe'),
    ('8a22ae56-d927-11ee-90e4-d8bbc174b998', 4, 'John Doe LinkedIn', 'https://linkedin.com/johndoe'),
    ('8a22ae56-d927-11ee-90e4-d8bbc174b998', 5, 'John Doe TikTok', 'https://tiktok.com/johndoe'),
    ('678e9012-e89b-12d3-a456-426614174006', 1, 'David Snapchat', 'https://facebook.com/david'),
    ('678e9012-e89b-12d3-a456-426614174006', 2, 'David Twitter', 'https://twitter.com/david'),
    ('678e9012-e89b-12d3-a456-426614174006', 3, 'David Instagram', 'https://instagram.com/david'),
    ('678e9012-e89b-12d3-a456-426614174006', 4, 'David LinkedIn', 'https://linkedin.com/david'),
    ('678e9012-e89b-12d3-a456-426614174006', 5, 'David TikTok', 'https:/tiktok.com/david'),
    ('678e9012-e89b-12d3-a456-426614174006', 6, 'David Telegram', 'https://t.me/david');
    ('678e9012-e89b-12d3-a456-426614174006', 6, 'David Telegram 2', 'https://t.me/david 2222');


SELECT 
    id, 
    social_name, 
    social_icon,
    name_of_social,
    link_of_social
FROM 
    socials s 
JOIN 
    user_social us ON s.id = us.social_id
WHERE 
    us.user_id = '8a22ae56-d927-11ee-90e4-d8bbc174b998';

SELECT 
    id, 
    social_name, 
    social_icon,
    name_of_social,
    link_of_social  
FROM 
    socials s 
JOIN 
    user_social us ON s.id = us.social_id
WHERE 
    us.user_id = '678e9012-e89b-12d3-a456-426614174006';