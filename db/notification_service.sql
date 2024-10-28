BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    order_type VARCHAR(150) NOT NULL,
    session_id VARCHAR(255) NOT NULL,
    card VARCHAR(30) NOT NULL,
    event_date DATETIME NOT NULL,
    website_url TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


INSERT INTO events(order_type, session_id, card, event_date, website_url)
VALUES(
    'Purchase',
    '29827525-06c9-4b1e-9d9b-7c4584e82f56',
    '4433**1409',
    '2023-01-04 13:44:52.835626 +00:00',
    'https://amazon.com'),
(
    'CardVerify',
    '500cf308-e666-4639-aa9f-f6376015d1b4',
    '4433**1409',
    'eventDate": "2023-04-07 05:29:54.362216 +00:00',
    'https://adidas.com'),
(
    'SendOtp',
    '500cf308-e666-4639-aa9f-f6376015d1b4',
    '4433**1409',
    '2023-04-06 22:52:34.930150 +00:00',
    'https://somon.tj'
);

COMMIT;