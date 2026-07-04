-- Diosys schema rollback. Drop children before parents to respect FKs.

DROP TABLE IF EXISTS `ms_system_setting`;
DROP TABLE IF EXISTS `ms_client_message`;
DROP TABLE IF EXISTS `ms_service`;

DROP TABLE IF EXISTS `ms_project_image`;
DROP TABLE IF EXISTS `ms_project_technology`;
DROP TABLE IF EXISTS `ms_project_feature`;
DROP TABLE IF EXISTS `ms_project`;

DROP TABLE IF EXISTS `ms_language`;
DROP TABLE IF EXISTS `ms_skill`;
DROP TABLE IF EXISTS `ms_certificate`;

DROP TABLE IF EXISTS `ms_education_achievement`;
DROP TABLE IF EXISTS `ms_education`;

DROP TABLE IF EXISTS `ms_experience_responsibility`;
DROP TABLE IF EXISTS `ms_experience_technology`;
DROP TABLE IF EXISTS `ms_experience`;

DROP TABLE IF EXISTS `ms_summary_fact`;
DROP TABLE IF EXISTS `ms_summary_stat`;
DROP TABLE IF EXISTS `ms_summary`;

DROP TABLE IF EXISTS `ms_user`;

DROP TABLE IF EXISTS `lk_project_status`;
DROP TABLE IF EXISTS `lk_user_role`;
