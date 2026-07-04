CREATE TABLE IF NOT EXISTS `ms_professional_project` (
    `professionalProjectID` INT AUTO_INCREMENT PRIMARY KEY,
    `userID`               INT NOT NULL,
    `title`                VARCHAR(255) NOT NULL,
    `company`              VARCHAR(200) NOT NULL,
    `summary`              TEXT NULL,
    `thumbnailGdriveID`    VARCHAR(255) NULL,
    `thumbnailFileName`    VARCHAR(255) NULL,
    `orderNo`              INT NOT NULL DEFAULT 0,
    `createdDate`          DATETIME DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT `fkProfProjUser` FOREIGN KEY (`userID`) REFERENCES `ms_user` (`userID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `ms_professional_project_feature` (
    `featureID`             INT AUTO_INCREMENT PRIMARY KEY,
    `professionalProjectID` INT NOT NULL,
    `title`                 VARCHAR(200) NOT NULL,
    `description`           TEXT NULL,
    `orderNo`               INT NOT NULL DEFAULT 0,
    CONSTRAINT `fkProfFeatProj` FOREIGN KEY (`professionalProjectID`) REFERENCES `ms_professional_project` (`professionalProjectID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `ms_professional_project_feature_image` (
    `featureImageID` INT AUTO_INCREMENT PRIMARY KEY,
    `featureID`      INT NOT NULL,
    `gdriveID`       VARCHAR(255) NOT NULL,
    `fileName`       VARCHAR(255) NOT NULL,
    `caption`        VARCHAR(500) NULL,
    `orderNo`        INT NOT NULL DEFAULT 0,
    CONSTRAINT `fkProfFeatImgFeat` FOREIGN KEY (`featureID`) REFERENCES `ms_professional_project_feature` (`featureID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
