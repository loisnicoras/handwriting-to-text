-- Create database if not exists
CREATE DATABASE IF NOT EXISTS my_database;

-- Use the database
USE my_database;

-- Create the exercises table
CREATE TABLE IF NOT EXISTS exercises (
    id INT AUTO_INCREMENT PRIMARY KEY,
    exercise_name VARCHAR(255),
    audio_path VARCHAR(255),
    text VARCHAR(1000)
);

-- Insert sample data
INSERT INTO exercises (exercise_name, audio_path, text) VALUES
('Exercise 1', '/path/to/audio1.mp3', 'Text for exercise 1'),
('Exercise 2', '/path/to/audio2.mp3', 'Text for exercise 2'),
('Exercise 3', '/path/to/audio3.mp3', 'Text for exercise 3');

-- Create the users table
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255)
);

-- Insert sample data into the users table
INSERT INTO users (first_name, last_name, email, password) VALUES
('John', 'Doe', 'john@example.com', 'password1'),
('Jane', 'Smith', 'jane@example.com', 'password2'),
('Alice', 'Johnson', 'alice@example.com', 'password3');

-- Create the users_results table
CREATE TABLE IF NOT EXISTS users_results (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    exercise_id INT NOT NULL,
    photo_text VARCHAR(255) NOT NULL,
    generate_text VARCHAR(255) NOT NULL,
    result INT,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Insert sample data into the users_results table
INSERT INTO users_results (user_id, exercise_id, photo_text, generate_text, result) VALUES
(1, 1, 'photo_text1', 'generate_text1', 80),
(1, 2, 'photo_text2', 'generate_text2', 90),
(2, 1, 'photo_text3', 'generate_text3', 75),
(2, 2, 'photo_text4', 'generate_text4', 85),
(3, 1, 'photo_text5', 'generate_text5', 95),
(3, 2, 'photo_text6', 'generate_text6', 70);
