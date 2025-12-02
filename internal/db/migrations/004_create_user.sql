CREATE USER 'web'@'%' IDENTIFIED BY 'password';

GRANT SELECT, INSERT, UPDATE, DELETE
ON snippetbox.*
TO 'web'@'%';

