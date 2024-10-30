CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    description TEXT,
    completed BOOLEAN,
    task_date DATE,
    task_type ENUM('academia', 'work', 'extracurricular', 'other')
);

INSERT INTO tasks (title, description, completed, task_date, task_type)
VALUES ('Dummy Task', 'This is a Dummy task', false, '2024-10-15', 'work');

