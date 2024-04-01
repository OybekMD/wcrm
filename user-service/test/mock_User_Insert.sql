-- Insert mock data into the 'users' table
INSERT INTO users (id, first_name, last_name, username, bio, birth_day, email, password_hash, avatar, experience_level, coint, score, refresh_token)
VALUES
    ('8a22ae56-d927-11ee-90e4-d8bbc174b998', 'John', 'Doe', 'john_doe', 'Passionate about technology and coding.', '1990-05-15', 'john.doe@example.com', 'hashed_password_123', 'https://example.com/john_avatar.png', 2, 100, 500, 'refresh_token_john'),
    ('678e9012-e89b-12d3-a456-426614174006', 'David', 'Miller', 'david_miller', 'Tech geek and startup enthusiast.', '1993-04-18', 'david.miller@example.com', 'hashed_password_678', 'https://example.com/david_avatar.png', 3, 180, 800, 'refresh_token_diana'),
