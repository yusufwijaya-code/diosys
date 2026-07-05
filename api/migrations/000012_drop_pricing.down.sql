CREATE TABLE IF NOT EXISTS `ms_price_plan` (
    `planID`          INT AUTO_INCREMENT PRIMARY KEY,
    `title`           VARCHAR(150) NOT NULL,
    `subtitle`        VARCHAR(255) NULL,
    `price`           DECIMAL(12,2) NOT NULL DEFAULT 0,
    `currency`        VARCHAR(10)  NOT NULL DEFAULT 'IDR',
    `billingPeriod`   VARCHAR(50)  NULL,
    `badge`           VARCHAR(80)  NULL,
    `originalPrice`   DECIMAL(12,2) NULL,
    `discountPercent` TINYINT UNSIGNED NULL,
    `isFeatured`      TINYINT NOT NULL DEFAULT 0,
    `flagActive`      TINYINT NOT NULL DEFAULT 1,
    `orderNo`         INT     NOT NULL DEFAULT 0,
    `createdDate`     DATETIME DEFAULT CURRENT_TIMESTAMP,
    `editedDate`      DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `ms_price_feature` (
    `featureID`  INT AUTO_INCREMENT PRIMARY KEY,
    `planID`     INT NOT NULL,
    `text`       VARCHAR(255) NOT NULL,
    `isIncluded` TINYINT NOT NULL DEFAULT 1,
    `orderNo`    INT NOT NULL DEFAULT 0,
    CONSTRAINT `fkPriceFeaturePlan` FOREIGN KEY (`planID`) REFERENCES `ms_price_plan` (`planID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
