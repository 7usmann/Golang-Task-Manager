CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    description TEXT,
    completed BOOLEAN,
    task_date DATE,
    task_type ENUM('academia', 'work', 'extracurricular', 'other')
);

