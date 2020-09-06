-- SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
-- SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
-- SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

CREATE SCHEMA IF NOT EXISTS `summer` DEFAULT CHARACTER SET utf8;
USE `summer`;

 --
 -- table
 --
-- -----------------------------------------------------
-- Table `summer`.`signin`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `summer`.`signUp` (
`id` VARCHAR(64) NOT NULL COMMENT 'ID',
`password` VARCHAR(64) NOT NULL COMMENT 'パスワード',
`token` varchar(64) NOT NULL COMMENT 'アクセストークン',
PRIMARY KEY (`id`))
ENGINE = InnoDB
COMMENT = 'ユーザ';
