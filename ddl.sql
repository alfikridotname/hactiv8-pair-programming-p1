-- Create Database Comic
CREATE DATABASE IF NOT EXISTS comic;

-- Use DB
USE comic;

-- Create Table Comic
CREATE TABLE IF NOT EXISTS comic (
    id INT(11) NOT NULL AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;