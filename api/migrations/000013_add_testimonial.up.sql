CREATE TABLE IF NOT EXISTS `ms_testimonial` (
    `testimonialID`   INT AUTO_INCREMENT PRIMARY KEY,
    `clientName`      VARCHAR(150) NOT NULL,
    `clientRole`      VARCHAR(150) NULL,
    `clientCompany`   VARCHAR(150) NULL,
    `testimonialText` TEXT NOT NULL,
    `rating`          TINYINT UNSIGNED NOT NULL DEFAULT 5,
    `photoGdriveID`   VARCHAR(255) NULL,
    `photoFileName`   VARCHAR(255) NULL,
    `flagActive`      TINYINT NOT NULL DEFAULT 1,
    `orderNo`         INT NOT NULL DEFAULT 0,
    `createdDate`     DATETIME DEFAULT CURRENT_TIMESTAMP,
    `editedDate`      DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
