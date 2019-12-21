ALTER TABLE `tracks` ADD (
    `publisher` VARCHAR(255) NOT NULL,
    FOREIGN KEY (`publisher`) REFERENCES `profiles` (`uuid`)
)
