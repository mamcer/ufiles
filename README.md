# UFiles

## Dev HTTP server

    python -m SimpleHTTPServer

## MySQL

    docker pull mysql:5.7.30
    docker run -p 3306:3306 --name ufiles -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7.30

    docker exec -it ufiles mysql -uroot -p
    create database ufiles;

## MyCli

    mycli -h localhost -u root -D ufiles -P 3306

## SQL

    use ufiles;

    CREATE TABLE udirectory (
      id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
      name VARCHAR(255) NOT NULL,
      path VARCHAR(260) NOT NULL,
      parent_id BIGINT UNSIGNED NOT NULL,
      size BIGINT NOT NULL,
      count BIGINT NOT NULL,
      PRIMARY KEY (id)
      ) ENGINE=INNODB AUTO_INCREMENT=1540 DEFAULT CHARSET=utf8;

    ALTER TABLE udirectory 
    ADD CONSTRAINT fk_parent_directory 
    FOREIGN KEY (parent_id) 
    REFERENCES udirectory(id);

    CREATE TABLE ufile (
      id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
      name VARCHAR(255) NOT NULL,
      extension VARCHAR(50) NOT NULL,
      path VARCHAR(260) NOT NULL,
      modified DATETIME NOT NULL,
      size BIGINT NOT NULL,
      hash VARCHAR(40) NOT NULL,
      udirectory_id BIGINT UNSIGNED NOT NULL,
      PRIMARY KEY (id)
    ) ENGINE=INNODB AUTO_INCREMENT=1540 DEFAULT CHARSET=utf8;

    ALTER TABLE ufile 
    ADD CONSTRAINT fk_directory 
    FOREIGN KEY (udirectory_id) 
    REFERENCES udirectory(id);