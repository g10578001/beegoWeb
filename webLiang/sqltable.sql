CREATE TABLE user (
	id INT PRIMARY KEY ASC, 
	username VARCHAR(256) NOT NULL UNIQUE, 
	email VARCHAR(256) NOT NULL UNIQUE,
	password VARCHAR(256) NOT NULL UNIQUE
);