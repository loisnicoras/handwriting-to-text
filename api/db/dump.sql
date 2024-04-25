-- Create database if not exists
CREATE DATABASE IF NOT EXISTS my_database;

-- Use the database
USE my_database;

-- Create the audio_exercises table
CREATE TABLE IF NOT EXISTS audio_exercises (
    id INT AUTO_INCREMENT PRIMARY KEY,
    exercise_name VARCHAR(255),
    audio_path VARCHAR(255),
    text VARCHAR(1000)
);

CREATE TABLE IF NOT EXISTS vowels_exercises (
    id INT AUTO_INCREMENT PRIMARY KEY,
    exercise_name VARCHAR(255),
    vowel VARCHAR(5),
    comparison_text VARCHAR(1000),
    text VARCHAR(10000)
);

-- Create the users table
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    sub VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255),
    avatar_url VARCHAR(255)
);

-- Create the audio_results table
CREATE TABLE IF NOT EXISTS audio_results (
    id INT AUTO_INCREMENT PRIMARY KEY,
    sub VARCHAR(255) NOT NULL,
    exercise_id INT NOT NULL,
    photo_text VARCHAR(255) NOT NULL,
    gen_text VARCHAR(255) NOT NULL,
    result INT
    -- FOREIGN KEY (sub) REFERENCES users(sub)
);

CREATE TABLE IF NOT EXISTS vowels_results (
    id INT AUTO_INCREMENT PRIMARY KEY,
    sub VARCHAR(255) NOT NULL,
    exercise_id INT NOT NULL,
    text VARCHAR(255) NOT NULL,
    result INT
    -- FOREIGN KEY (sub) REFERENCES users(sub)
);
-- Insert sample data
INSERT INTO audio_exercises (exercise_name, audio_path, text) VALUES
('Exercise 1', '/path/to/audio1.mp3', 'Text for exercise 1'),
('Exercise 2', '/path/to/audio2.mp3', 'Text for exercise 2'),
('Exercise 3', '/path/to/audio3.mp3', 'Text for exercise 3');

INSERT INTO vowels_exercises (exercise_name, text) VALUES
('Exercise 1', 'Text for exercise 1'),
('Exercise 2', 'Text for exercise 2'),
('Exercise 3', 'Text for exercise 3');