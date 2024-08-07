CREATE TABLE tasks(
    id SERIAL PRIMARY KEY NOT NULL,
    title VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    priority VARCHAR NOT NULL,
    status VARCHAR NOT NULL,
    createdAt TIMESTAMP NOT NULL,
    dueDate TIMESTAMP NOT NULL
);

INSERT INTO tasks (title, description, priority, status, createdAt, dueDate)
VALUES 
    ('Task 1', 'Description for task 1', 'High', 'ToDo', '2024-08-07 10:00:00', '2024-08-14 10:00:00'),
    ('Task 2', 'Description for task 2', 'Medium', 'In Progress', '2024-08-07 11:00:00', '2024-08-15 11:00:00'),
    ('Task 3', 'Description for task 3', 'Low', 'Done', '2024-08-07 12:00:00', '2024-08-16 12:00:00');
