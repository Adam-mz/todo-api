CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    task VARCHAR(255) NOT NULL,
    is_done BOOLEAN DEFAULT FALSE
    
    
);