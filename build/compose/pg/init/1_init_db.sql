CREATE DATABASE IF NOT EXIST zadanie;
CREATE USER backend_user WITH PASSWORD 'secret_password';
GRANT ALL PRIVILEGES ON DATABASE zadanie TO backend_user;

