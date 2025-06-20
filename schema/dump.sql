CREATE TABLE deck (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    total_cards INTEGER,
    next_review_at TIMESTAMP 
);
CREATE TABLE flashcard (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    deck_id INTEGER NOT NULL,
    message TEXT,
    next_review_at TIMESTAMP,
    ease_factor DOUBLE PRECISION,
    CONSTRAINT fk_deck
        FOREIGN KEY (deck_id)
        REFERENCES deck(id)
        ON DELETE CASCADE
);
