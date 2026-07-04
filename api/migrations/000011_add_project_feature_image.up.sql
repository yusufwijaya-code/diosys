ALTER TABLE ms_project_feature ADD COLUMN description TEXT NULL AFTER text;

CREATE TABLE IF NOT EXISTS `ms_project_feature_image` (
    `projectFeatureImageID` INT AUTO_INCREMENT PRIMARY KEY,
    `projectFeatureID`      INT NOT NULL,
    `gdriveID`              VARCHAR(255) NOT NULL,
    `fileName`              VARCHAR(255) NOT NULL,
    `caption`               VARCHAR(500) NULL,
    `orderNo`               INT NOT NULL DEFAULT 0,
    CONSTRAINT `fkProjFeatImg` FOREIGN KEY (`projectFeatureID`) REFERENCES `ms_project_feature` (`projectFeatureID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
