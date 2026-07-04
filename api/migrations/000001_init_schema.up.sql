-- Diosys database schema.
-- Conventions:
--   * camelCase columns, primary keys suffixed with ID (e.g. userID, projectID).
--   * Table prefixes: ms_ (master), map_ (mapping/junction), lk_ (lookup).

-- ---------------------------------------------------------------------------
-- Lookup tables (lk_)
-- ---------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `lk_user_role` (
    `userRoleID` INT AUTO_INCREMENT PRIMARY KEY,
    `code`       VARCHAR(50)  NOT NULL,
    `name`       VARCHAR(100) NOT NULL,
    UNIQUE KEY `uqUserRoleCode` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `lk_project_status` (
    `projectStatusID` INT AUTO_INCREMENT PRIMARY KEY,
    `code`            VARCHAR(50)  NOT NULL,
    `name`            VARCHAR(100) NOT NULL,
    UNIQUE KEY `uqProjectStatusCode` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ---------------------------------------------------------------------------
-- Master: users / developers
-- ---------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `ms_user` (
    `userID`         INT AUTO_INCREMENT PRIMARY KEY,
    `username`       VARCHAR(100) NOT NULL,            -- URL slug, e.g. yusuf-wijaya
    `email`          VARCHAR(150) NOT NULL,
    `googleSub`      VARCHAR(100) NULL,                -- Google account subject id
    `fullName`       VARCHAR(150) NOT NULL,
    `jobTitle`       VARCHAR(150) NULL,                -- role shown on directory card
    `intro`          VARCHAR(500) NULL,                -- brief intro for directory card
    `bio`            TEXT         NULL,                -- long rich-text biography
    `specialization` VARCHAR(150) NULL,
    `phone`          VARCHAR(50)  NULL,
    `website`        VARCHAR(255) NULL,
    `location`       VARCHAR(150) NULL,
    `photoFileName`  VARCHAR(255) NULL,
    `photoGdriveID`  VARCHAR(255) NULL,
    `userRoleID`     INT          NULL,
    `isAdmin`        TINYINT      NOT NULL DEFAULT 0,
    `flagActive`     TINYINT      NOT NULL DEFAULT 1,
    `orderNo`        INT          NOT NULL DEFAULT 0,
    `createdDate`    DATETIME DEFAULT CURRENT_TIMESTAMP,
    `editedDate`     DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `uqUserUsername` (`username`),
    UNIQUE KEY `uqUserEmail` (`email`),
    CONSTRAINT `fkUserRole` FOREIGN KEY (`userRoleID`) REFERENCES `lk_user_role` (`userRoleID`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ---------------------------------------------------------------------------
-- Master: per-developer portfolio (summary)
-- ---------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `ms_summary` (
    `summaryID`   INT AUTO_INCREMENT PRIMARY KEY,
    `userID`      INT NOT NULL,
    `content`     TEXT NOT NULL,
    `createdDate` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `editedDate`  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `uqSummaryUser` (`userID`),
    CONSTRAINT `fkSummaryUser` FOREIGN KEY (`userID`) REFERENCES `ms_user` (`userID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `ms_summary_stat` (
    `summaryStatID` INT AUTO_INCREMENT PRIMARY KEY,
    `summaryID`     INT NOT NULL,
    `number`        VARCHAR(50)  NOT NULL,
    `label`         VARCHAR(150) NOT NULL,
    `orderNo`       INT NOT NULL DEFAULT 0,
    CONSTRAINT `fkSummaryStatSummary` FOREIGN KEY (`summaryID`) REFERENCES `ms_summary` (`summaryID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `ms_summary_fact` (
    `summaryFactID` INT AUTO_INCREMENT PRIMARY KEY,
    `summaryID`     INT NOT NULL,
    `icon`          VARCHAR(50)  NOT NULL,
    `text`          VARCHAR(255) NOT NULL,
    `orderNo`       INT NOT NULL DEFAULT 0,
    CONSTRAINT `fkSummaryFactSummary` FOREIGN KEY (`summaryID`) REFERENCES `ms_summary` (`summaryID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ---------------------------------------------------------------------------
-- Master: per-developer portfolio (experience)
-- ---------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `ms_experience` (
    `experienceID` INT AUTO_INCREMENT PRIMARY KEY,
    `userID`       INT NOT NULL,
    `position`     VARCHAR(150) NOT NULL,
    `company`      VARCHAR(150) NOT NULL,
    `period`       VARCHAR(100) NOT NULL,
    `orderNo`      INT NOT NULL DEFAULT 0,
    `createdDate`  DATETIME DEFAULT CURRENT_TIMESTAMP,
    `editedDate`   DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT `fkExperienceUser` FOREIGN KEY (`userID`) REFERENCES `ms_user` (`userID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `ms_experience_technology` (
    `experienceTechnologyID` INT AUTO_INCREMENT PRIMARY KEY,
    `experienceID`           INT NOT NULL,
    `name`                   VARCHAR(100) NOT NULL,
    `orderNo`                INT NOT NULL DEFAULT 0,
    CONSTRAINT `fkExperienceTechnologyExperience` FOREIGN KEY (`experienceID`) REFERENCES `ms_experience` (`experienceID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `ms_experience_responsibility` (
    `experienceResponsibilityID` INT AUTO_INCREMENT PRIMARY KEY,
    `experienceID`               INT NOT NULL,
    `description`                TEXT NOT NULL,
    `orderNo`                    INT NOT NULL DEFAULT 0,
    CONSTRAINT `fkExperienceResponsibilityExperience` FOREIGN KEY (`experienceID`) REFERENCES `ms_experience` (`experienceID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ---------------------------------------------------------------------------
-- Master: per-developer portfolio (education)
-- ---------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `ms_education` (
    `educationID`  INT AUTO_INCREMENT PRIMARY KEY,
    `userID`       INT NOT NULL,
    `degree`       VARCHAR(150) NOT NULL,
    `institution`  VARCHAR(200) NOT NULL,
    `year`         VARCHAR(20)  NOT NULL,
    `type`         VARCHAR(100) NOT NULL,
    `orderNo`      INT NOT NULL DEFAULT 0,
    `createdDate`  DATETIME DEFAULT CURRENT_TIMESTAMP,
    `editedDate`   DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT `fkEducationUser` FOREIGN KEY (`userID`) REFERENCES `ms_user` (`userID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `ms_education_achievement` (
    `educationAchievementID` INT AUTO_INCREMENT PRIMARY KEY,
    `educationID`            INT NOT NULL,
    `description`            TEXT NOT NULL,
    `orderNo`                INT NOT NULL DEFAULT 0,
    CONSTRAINT `fkEducationAchievementEducation` FOREIGN KEY (`educationID`) REFERENCES `ms_education` (`educationID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ---------------------------------------------------------------------------
-- Master: per-developer portfolio (certificate, skill, language)
-- ---------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `ms_certificate` (
    `certificateID` INT AUTO_INCREMENT PRIMARY KEY,
    `userID`        INT NOT NULL,
    `name`          VARCHAR(200) NOT NULL,
    `issuer`        VARCHAR(200) NOT NULL,
    `period`        VARCHAR(100) NOT NULL,
    `link`          VARCHAR(500) NULL,
    `orderNo`       INT NOT NULL DEFAULT 0,
    `createdDate`   DATETIME DEFAULT CURRENT_TIMESTAMP,
    `editedDate`    DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT `fkCertificateUser` FOREIGN KEY (`userID`) REFERENCES `ms_user` (`userID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `ms_skill` (
    `skillID`     INT AUTO_INCREMENT PRIMARY KEY,
    `userID`      INT NOT NULL,
    `name`        VARCHAR(150) NOT NULL,
    `level`       VARCHAR(50)  NOT NULL,
    `category`    VARCHAR(50)  NOT NULL,
    `orderNo`     INT NOT NULL DEFAULT 0,
    `createdDate` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `editedDate`  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT `fkSkillUser` FOREIGN KEY (`userID`) REFERENCES `ms_user` (`userID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `ms_language` (
    `languageID`  INT AUTO_INCREMENT PRIMARY KEY,
    `userID`      INT NOT NULL,
    `name`        VARCHAR(100) NOT NULL,
    `level`       VARCHAR(50)  NOT NULL,
    `icon`        VARCHAR(50)  NOT NULL,
    `orderNo`     INT NOT NULL DEFAULT 0,
    `createdDate` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `editedDate`  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT `fkLanguageUser` FOREIGN KEY (`userID`) REFERENCES `ms_user` (`userID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ---------------------------------------------------------------------------
-- Master: per-developer projects (rich detail + gallery)
-- ---------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `ms_project` (
    `projectID`         INT AUTO_INCREMENT PRIMARY KEY,
    `userID`            INT NOT NULL,
    `title`             VARCHAR(200) NOT NULL,
    `summary`           VARCHAR(500) NULL,             -- short description for cards
    `body`              TEXT         NULL,             -- rich-text detailed description
    `client`            VARCHAR(200) NULL,
    `projectLink`       VARCHAR(500) NULL,             -- live deployment URL
    `repoLink`          VARCHAR(500) NULL,             -- repository URL
    `projectStatusID`   INT          NULL,
    `isFeatured`        TINYINT      NOT NULL DEFAULT 0,
    `thumbnailFileName` VARCHAR(255) NULL,
    `thumbnailGdriveID` VARCHAR(255) NULL,
    `orderNo`           INT NOT NULL DEFAULT 0,
    `createdDate`       DATETIME DEFAULT CURRENT_TIMESTAMP,
    `editedDate`        DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT `fkProjectUser` FOREIGN KEY (`userID`) REFERENCES `ms_user` (`userID`) ON DELETE CASCADE,
    CONSTRAINT `fkProjectStatus` FOREIGN KEY (`projectStatusID`) REFERENCES `lk_project_status` (`projectStatusID`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `ms_project_feature` (
    `projectFeatureID` INT AUTO_INCREMENT PRIMARY KEY,
    `projectID`        INT NOT NULL,
    `text`             VARCHAR(500) NOT NULL,
    `orderNo`          INT NOT NULL DEFAULT 0,
    CONSTRAINT `fkProjectFeatureProject` FOREIGN KEY (`projectID`) REFERENCES `ms_project` (`projectID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `ms_project_technology` (
    `projectTechnologyID` INT AUTO_INCREMENT PRIMARY KEY,
    `projectID`           INT NOT NULL,
    `name`                VARCHAR(100) NOT NULL,
    `orderNo`             INT NOT NULL DEFAULT 0,
    CONSTRAINT `fkProjectTechnologyProject` FOREIGN KEY (`projectID`) REFERENCES `ms_project` (`projectID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `ms_project_image` (
    `projectImageID` INT AUTO_INCREMENT PRIMARY KEY,
    `projectID`      INT NOT NULL,
    `fileName`       VARCHAR(255) NULL,
    `gdriveID`       VARCHAR(255) NULL,
    `caption`        VARCHAR(255) NULL,
    `displayOrder`   INT NOT NULL DEFAULT 0,
    CONSTRAINT `fkProjectImageProject` FOREIGN KEY (`projectID`) REFERENCES `ms_project` (`projectID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ---------------------------------------------------------------------------
-- Master: agency-level content (services, client messages, system settings)
-- ---------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `ms_service` (
    `serviceID`   INT AUTO_INCREMENT PRIMARY KEY,
    `title`       VARCHAR(150) NOT NULL,
    `description` TEXT         NULL,
    `icon`        VARCHAR(100) NULL,
    `orderNo`     INT NOT NULL DEFAULT 0,
    `flagActive`  TINYINT NOT NULL DEFAULT 1,
    `createdDate` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `editedDate`  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `ms_client_message` (
    `messageID`   INT AUTO_INCREMENT PRIMARY KEY,
    `clientName`  VARCHAR(150) NOT NULL,
    `clientEmail` VARCHAR(150) NOT NULL,
    `subject`     VARCHAR(200) NULL,
    `messageBody` TEXT NOT NULL,
    `isRead`      TINYINT NOT NULL DEFAULT 0,
    `isArchived`  TINYINT NOT NULL DEFAULT 0,
    `createdDate` DATETIME DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `ms_system_setting` (
    `settingID`    INT AUTO_INCREMENT PRIMARY KEY,
    `settingKey`   VARCHAR(100) NOT NULL,
    `settingValue` TEXT NULL,
    `description`  VARCHAR(255) NULL,
    `editedDate`   DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `uqSystemSettingKey` (`settingKey`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
