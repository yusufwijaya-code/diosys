CREATE TABLE IF NOT EXISTS `ms_experience_work` (
    `experienceWorkID` INT AUTO_INCREMENT PRIMARY KEY,
    `experienceID`     INT NOT NULL,
    `title`            VARCHAR(200) NOT NULL,
    `description`      TEXT NULL,
    `orderNo`          INT NOT NULL DEFAULT 0,
    CONSTRAINT `fkExpWorkExperience` FOREIGN KEY (`experienceID`) REFERENCES `ms_experience` (`experienceID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `ms_experience_work_technology` (
    `expWorkTechID`    INT AUTO_INCREMENT PRIMARY KEY,
    `experienceWorkID` INT NOT NULL,
    `name`             VARCHAR(100) NOT NULL,
    `orderNo`          INT NOT NULL DEFAULT 0,
    CONSTRAINT `fkExpWorkTechWork` FOREIGN KEY (`experienceWorkID`) REFERENCES `ms_experience_work` (`experienceWorkID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
