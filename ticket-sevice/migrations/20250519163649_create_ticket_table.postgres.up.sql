CREATE TABLE user_tickets (
    ticket_id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255),
    event_id VARCHAR(255),
    created_at DATE,
    updated_at DATE,
    seat_number INTEGER
);
