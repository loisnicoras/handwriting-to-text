CREATE TABLE `users` (
  id INT AUTO_INCREMENT PRIMARY KEY,
  sub VARCHAR(255) UNIQUE NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  name VARCHAR(255),
  avatar_url VARCHAR(255)
);

CREATE TABLE `exercises` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `exercise_name` varchar(255) NOT NULL,
  `audio_path` varchar(255) NOT NULL,
  `text` varchar(255) NOT NULL
);

CREATE TABLE `users_results` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `user_id` integer NOT NULL,
  `exercise_id` integer NOT NULL,
  `photo_text` varchar(255) NOT NULL,
  `generate_text` varchar(255) NOT NULL,
  `result` integer
);

ALTER TABLE `users_results` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `users_results` ADD FOREIGN KEY (`exercise_id`) REFERENCES `exercises` (`id`);
