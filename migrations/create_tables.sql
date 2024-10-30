-- Create the ENUM type first
DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'task_type_enum') THEN
        CREATE TYPE task_type_enum AS ENUM ('academia', 'work', 'extracurricular', 'other');
    END IF;
END$$;

-- Create the table using the ENUM type
CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    description TEXT,
    completed BOOLEAN,
    task_date DATE,
    task_type task_type_enum
);

-- Insert dummy data
INSERT INTO tasks (title, description, completed, task_date, task_type)
VALUES ('Dummy Task', 'This is a Dummy task', false, '2024-10-15', 'work');


