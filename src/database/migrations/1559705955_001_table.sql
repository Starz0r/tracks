CREATE TABLE `music` (
    `id` INT(8) UNSIGNED NOT NULL AUTO_INCREMENT UNIQUE,
    `date_created` DATETIME NOT NULL DEFAULT NOW(),
    `date_modified` DATETIME,
    `jacket` TEXT NOT NULL,
    `genre` VARCHAR(256) NOT NULL,
    `credit` VARCHAR(512),
    `displaybpm` VARCHAR(16) NOT NULL,
    `length` INT(8) UNSIGNED NOT NULL,
    `title` VARCHAR(256) NOT NULL,
    `title_romani` VARCHAR(256),
    `artists` VARCHAR(256) NOT NULL,
    `artists_romani` VARCHAR(256),
    `subtitle` VARCHAR(256),
    `subtitle_romani` VARCHAR(256),
    PRIMARY KEY (`id`)
)
