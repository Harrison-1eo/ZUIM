CREATE DATABASE imdsb;

CREATE USER 'imdsbMaster'@'localhost' IDENTIFIED BY 'imdsbPassword';

GRANT ALL PRIVILEGES ON imdsb.* TO 'imdsbMaster'@'localhost';

FLUSH PRIVILEGES;
