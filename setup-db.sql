CREATE DATABASE IF NOT EXISTS notesdb;
CREATE USER IF NOT EXISTS 'notesuser'@'localhost' IDENTIFIED BY 'notespwd';
CREATE USER IF NOT EXISTS 'notesuser'@'%' IDENTIFIED BY 'notespwd';
GRANT ALL PRIVILEGES ON notesdb.* TO 'notesuser'@'localhost';
GRANT ALL PRIVILEGES ON notesdb.* TO 'notesuser'@'%';
FLUSH PRIVILEGES;
