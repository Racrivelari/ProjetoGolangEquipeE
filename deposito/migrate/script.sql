-- SCRIPT PRA MY SQL

drop database if exists deposito;
CREATE database IF NOT EXISTS deposito DEFAULT CHARACTER SET utf8MB4 ;
USE deposito;
#SET SQL_SAFE_UPDATES = 0;

CREATE TABLE IF NOT EXISTS Product (
  id integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  name VARCHAR(45) NOT NULL,
  price float, 
  code VARCHAR(45) NOT NULL,
  created_at datetime NULL DEFAULT NOW(),
  Unique(code));
  
CREATE TABLE IF NOT EXISTS Logs  (
	id_log integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
	updated_at datetime DEFAULT NOW(),
    id INT NOT NULL,
	constraint fk_product_log foreign key (id) references Product(id)
);

INSERT INTO Product (name, price, code) values ("Iphone", 5000, "rcb");
INSERT INTO Product (name, price, code) values ("Galaxy S22", 4000, "cel2");
INSERT INTO Product (name, price, code) values ("Pocophone", 1000, "cel3");
INSERT INTO Product (name, price, code) values ("Windows Phone", 1000, "cel4");

select * from Product;
select * from Logs;


DELIMITER $$

CREATE TRIGGER product_update
AFTER UPDATE
ON Product FOR EACH ROW
BEGIN
        INSERT INTO Logs(id,updated_at)
        VALUES(old.id, current_timestamp());
END$$

DELIMITER ;







  

  
