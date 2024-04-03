INSERT INTO users (id, first_name, last_name, username, phone_number, bio, birth_day, email, avatar, password, refresh_token)
VALUES
    ('c3b7e55d-f3f7-4d91-9d5d-85b605f19b91', 'John', 'Doe', 'johndoe', '123456789', 'Lorem ipsum dolor sit amet', '1990-01-01', 'john@example.com', 'http://example.com/avatar1.jpg', 'hashed_password_1', 'refresh_token_1'),
    ('9b785693-5df6-4e3d-a594-7dc4df1840c8', 'Alice', 'Smith', 'alicesmith', '987654321', 'Consectetur adipiscing elit', '1985-05-15', 'alice@example.com', 'http://example.com/avatar2.jpg', 'hashed_password_2', 'refresh_token_2'),
    ('5e231b33-26a7-46d1-bb94-84a5d0a51dc1', 'Michael', 'Johnson', 'michaeljohnson', '5551234567', 'Sed do eiusmod tempor incididunt', '1988-10-20', 'michael@example.com', 'http://example.com/avatar3.jpg', 'hashed_password_3', 'refresh_token_3'),
    ('f8e15130-0877-44b3-a8e8-832f4cf40bc3', 'Emily', 'Brown', 'emilybrown', '4449876543', 'Ut labore et dolore magna aliqua', '1992-03-12', 'emily@example.com', 'http://example.com/avatar4.jpg', 'hashed_password_4', 'refresh_token_4'),
    ('de19b982-23a4-4e86-b02b-c464c6f780bb', 'Daniel', 'Martinez', 'danielmartinez', '1112223333', 'Ut enim ad minim veniam', '1982-12-08', 'daniel@example.com', 'http://example.com/avatar5.jpg', 'hashed_password_5', 'refresh_token_5'),
    ('937ca1e2-f024-45c6-b10d-97787654ef11', 'Jessica', 'Lee', 'jessicalee', '9998887777', 'Quis nostrud exercitation ullamco', '1995-06-25', 'jessica@example.com', 'http://example.com/avatar6.jpg', 'hashed_password_6', 'refresh_token_6'),
    ('c6dc025e-0f3b-4f43-a107-8a36a4e1ab6f', 'David', 'Wilson', 'davidwilson', '3334445555', 'Duis aute irure dolor in reprehenderit', '1980-09-30', 'david@example.com', 'http://example.com/avatar7.jpg', 'hashed_password_7', 'refresh_token_7'),
    ('4781c550-0e57-46a8-9ff2-0fda899456fb', 'Sarah', 'Taylor', 'sarahtaylor', '6667778888', 'Excepteur sint occaecat cupidatat non proident', '1998-07-18', 'sarah@example.com', 'http://example.com/avatar8.jpg', 'hashed_password_8', 'refresh_token_8'),
    ('8f6a4328-47e0-4629-bb3b-67890123cdab', 'Matthew', 'Anderson', 'matthewanderson', '2223334444', 'Sunt in culpa qui officia deserunt mollit', '1987-04-05', 'matthew@example.com', 'http://example.com/avatar9.jpg', 'hashed_password_9', 'refresh_token_9'),
    ('7a97c5d2-1323-4fb0-a56d-89abcdeff001', 'Megan', 'Thomas', 'meganthomas', '9990001111', 'Sed ut perspiciatis unde omnis iste natus error sit voluptatem', '1990-11-28', 'megan@example.com', 'http://example.com/avatar10.jpg', 'hashed_password_10', 'refresh_token_10'),
    ('eb74aefc-625e-4edc-b31c-345678901234', 'Christopher', 'Harris', 'christopherharris', '7778889999', 'Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit', '1984-08-15', 'christopher@example.com', 'http://example.com/avatar11.jpg', 'hashed_password_11', 'refresh_token_11'),
    ('537bcb34-1f69-4eaa-af56-345678901235', 'Lauren', 'King', 'laurenking', '5556667777', 'Neque porro quisquam est qui dolorem ipsum quia dolor sit amet', '1983-02-20', 'lauren@example.com', 'http://example.com/avatar12.jpg', 'hashed_password_12', 'refresh_token_12'),
    ('afdc3011-96a4-4b71-860e-345678901236', 'James', 'Evans', 'jamesevans', '1112223333', 'Consectetur, adipisci velit', '1993-09-10', 'james@example.com', 'http://example.com/avatar13.jpg', 'hashed_password_13', 'refresh_token_13'),
    ('b673d3ac-9ba7-4710-8f11-345678901237', 'Olivia', 'Garcia', 'oliviagarcia', '8889990000', 'Sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem', '1996-12-03', 'olivia@example.com', 'http://example.com/avatar14.jpg', 'hashed_password_14', 'refresh_token_14'),
    ('c212a3a4-ea8b-4ae2-b35d-345678901238', 'Andrew', 'Rodriguez', 'andrewrodriguez', '1234567890', 'Magni dolores eos qui ratione voluptatem sequi nesciunt', '1986-06-01', 'andrew@example.com', 'http://example.com/avatar15.jpg', 'hashed_password_15', 'refresh_token_15');

INSERT INTO category_icons (name, picture) VALUES
    ('Category 1', 'https://example.com/category1.jpg'),
    ('Category 2', 'https://example.com/category2.jpg'),
    ('Category 3', 'https://example.com/category3.jpg'),
    ('Category 4', 'https://example.com/category4.jpg'),
    ('Category 5', 'https://example.com/category5.jpg'),
    ('Category 6', 'https://example.com/category6.jpg'),
    ('Category 7', 'https://example.com/category7.jpg'),
    ('Category 8', 'https://example.com/category8.jpg'),
    ('Category 9', 'https://example.com/category9.jpg'),
    ('Category 10', 'https://example.com/category10.jpg'),
    ('Category 11', 'https://example.com/category11.jpg'),
    ('Category 12', 'https://example.com/category12.jpg'),
    ('Category 13', 'https://example.com/category13.jpg'),
    ('Category 14', 'https://example.com/category14.jpg'),
    ('Category 15', 'https://example.com/category15.jpg');

INSERT INTO categorys (name, icon_id) VALUES
    ('Category 1', 1),
    ('Category 2', 2),
    ('Category 3', 3),
    ('Category 4', 4),
    ('Category 5', 5),
    ('Category 6', 6),
    ('Category 7', 7),
    ('Category 8', 8),
    ('Category 9', 9),
    ('Category 10', 10),
    ('Category 11', 11),
    ('Category 12', 12),
    ('Category 13', 13),
    ('Category 14', 14),
    ('Category 15', 15);


INSERT INTO products (title, description, price, picture, category_id) VALUES
    ('Product 1', 'Description for Product 1', 1000, 'product1.jpg', 1),
    ('Product 2', 'Description for Product 2', 1500, 'product2.jpg', 2),
    ('Product 3', 'Description for Product 3', 2000, 'product3.jpg', 3),
    ('Product 4', 'Description for Product 4', 2500, 'product4.jpg', 4),
    ('Product 5', 'Description for Product 5', 3000, 'product5.jpg', 5),
    ('Product 6', 'Description for Product 6', 3500, 'product6.jpg', 6),
    ('Product 7', 'Description for Product 7', 4000, 'product7.jpg', 7),
    ('Product 8', 'Description for Product 8', 4500, 'product8.jpg', 8),
    ('Product 9', 'Description for Product 9', 5000, 'product9.jpg', 9),
    ('Product 10', 'Description for Product 10', 5500, 'product10.jpg', 10),
    ('Product 11', 'Description for Product 11', 6000, 'product11.jpg', 11),
    ('Product 12', 'Description for Product 12', 6500, 'product12.jpg', 12),
    ('Product 13', 'Description for Product 13', 7000, 'product13.jpg', 13),
    ('Product 14', 'Description for Product 14', 7500, 'product14.jpg', 14),
    ('Product 15', 'Description for Product 15', 8000, 'product15.jpg', 15);

INSERT INTO orderproducts (user_id, product_id) VALUES
    ('67686d61-45a7-4a88-9f61-6a869a3b97aa', 1),
    ('67686d61-45a7-4a88-9f61-6a869a3b97aa', 2),
    ('67686d61-45a7-4a88-9f61-6a869a3b97aa', 3),
    ('67686d61-45a7-4a88-9f61-6a869a3b97aa', 4),
    ('67686d61-45a7-4a88-9f61-6a869a3b97aa', 5),
    ('67686d61-45a7-4a88-9f61-6a869a3b97aa', 6),
    ('67686d61-45a7-4a88-9f61-6a869a3b97aa', 7),
    ('67686d61-45a7-4a88-9f61-6a869a3b97aa', 8),
    ('67686d61-45a7-4a88-9f61-6a869a3b97aa', 9),
    ('67686d61-45a7-4a88-9f61-6a869a3b97aa', 10),
    ('67686d61-45a7-4a88-9f61-6a869a3b97aa', 11),
    ('67686d61-45a7-4a88-9f61-6a869a3b97aa', 12),
    ('67686d61-45a7-4a88-9f61-6a869a3b97aa', 13),
    ('67686d61-45a7-4a88-9f61-6a869a3b97aa', 14),
    ('67686d61-45a7-4a88-9f61-6a869a3b97aa', 15);

INSERT INTO comments (content, user_id, product_id) VALUES
    ('Great product!', '67686d61-45a7-4a88-9f61-6a869a3b97aa', 1),
    ('Nice description!', '67686d61-45a7-4a88-9f61-6a869a3b97aa', 2),
    ('Love it!', '67686d61-45a7-4a88-9f61-6a869a3b97aa', 3),
    ('Excellent quality!', '67686d61-45a7-4a88-9f61-6a869a3b97aa', 4),
    ('Highly recommended!', '67686d61-45a7-4a88-9f61-6a869a3b97aa', 5),
    ('Fast delivery!', '67686d61-45a7-4a88-9f61-6a869a3b97aa', 6),
    ('Good customer service!', '67686d61-45a7-4a88-9f61-6a869a3b97aa', 7),
    ('Awesome!', '67686d61-45a7-4a88-9f61-6a869a3b97aa', 8),
    ('Very satisfied!', '67686d61-45a7-4a88-9f61-6a869a3b97aa', 9),
    ('Impressive!', '67686d61-45a7-4a88-9f61-6a869a3b97aa', 10),
    ('Good value for money!', '67686d61-45a7-4a88-9f61-6a869a3b97aa', 11),
    ('Exactly as described!', '67686d61-45a7-4a88-9f61-6a869a3b97aa', 12),
    ('Prompt response!', '67686d61-45a7-4a88-9f61-6a869a3b97aa', 13),
    ('Great communication!', '67686d61-45a7-4a88-9f61-6a869a3b97aa', 14),
    ('Amazing experience!', '67686d61-45a7-4a88-9f61-6a869a3b97aa', 15);

INSERT INTO admins (email) VALUES ('oybekatamatov799@gmail.com');
