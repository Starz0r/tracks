ALTER TABLE `tracks` ADD (
    `publisher` INT(8) UNSIGNED NOT NULL,
    FOREIGN KEY (`publisher`) REFERENCES `profiles` (`id`)
)
