ALTER TABLE `tracks` ADD (
    `publisher` VARCHAR(255) UNSIGNED NOT NULL,
    FOREIGN KEY (`publisher`) REFERENCES `profiles` (`uuid`)
)
