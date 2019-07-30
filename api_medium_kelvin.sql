DROP DATABASE IF EXISTS api_medium_kelvin;
CREATE DATABASE api_medium_kelvin;
USE api_medium_kelvin;
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    age INT NOT NULL
);


DROP DATABASE IF EXISTS api_medium_kelvin_test;
CREATE TABLE users_test (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    age INT NOT NULL
);