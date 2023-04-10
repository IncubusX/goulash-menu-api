-- +goose NO TRANSACTION
-- +goose Up
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
SET NAMES 'utf8mb4';

CREATE TABLE ingredient_type (
id INT NOT NULL AUTO_INCREMENT,
title VARCHAR(255) DEFAULT NULL,
code CHAR(1) DEFAULT NULL,
PRIMARY KEY (id)
)
ENGINE = INNODB,
AUTO_INCREMENT = 4,
CHARACTER SET utf8mb4,
COLLATE utf8mb4_0900_ai_ci;

CREATE TABLE ingredient (
id INT NOT NULL AUTO_INCREMENT,
type_id INT NOT NULL,
title VARCHAR(255) NOT NULL,
price DECIMAL(19, 2) NOT NULL,
PRIMARY KEY (id)
)
ENGINE = INNODB,
AUTO_INCREMENT = 10,
CHARACTER SET utf8mb4,
COLLATE utf8mb4_0900_ai_ci;

ALTER TABLE ingredient
ADD CONSTRAINT FK_ingredient_type_id FOREIGN KEY (type_id)
REFERENCES ingredient_type(id);

INSERT INTO ingredient_type VALUES
(1, 'Тесто', 'd'),
(2, 'Сыр', 'c'),
(3, 'Начинка', 'i');

INSERT INTO ingredient VALUES
(1, 1, 'Тонкое тесто', 100.00),
(2, 1, 'Пышное тесто', 110.00),
(3, 1, 'Ржаное тесто', 150.00),
(4, 2, 'Моцарелла', 50.00),
(5, 2, 'Рикотта', 70.00),
(6, 3, 'Колбаса', 30.00),
(7, 3, 'Ветчина', 35.00),
(8, 3, 'Грибы', 50.00),
(9, 3, 'Томаты', 10.00);
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS ingredient;
DROP TABLE IF EXISTS ingredient_type;
-- +goose StatementEnd
