-- Reference / default data for Diosys. Idempotent via unique keys.

INSERT INTO `lk_user_role` (`code`, `name`) VALUES
    ('admin', 'Administrator'),
    ('developer', 'Developer')
ON DUPLICATE KEY UPDATE `name` = VALUES(`name`);

INSERT INTO `lk_project_status` (`code`, `name`) VALUES
    ('completed', 'Completed'),
    ('ongoing', 'Ongoing'),
    ('maintenance', 'Maintenance')
ON DUPLICATE KEY UPDATE `name` = VALUES(`name`);

INSERT INTO `ms_system_setting` (`settingKey`, `settingValue`, `description`) VALUES
    ('companyName', 'Diosys', 'Company display name'),
    ('companyTagline', 'Premium technology solutions: custom apps, websites, and AI integrations.', 'Hero tagline'),
    ('contactEmail', 'hello@diosys.com', 'Public contact email'),
    ('whatsappNumber', '6281234567890', 'WhatsApp number in international format, no plus sign'),
    ('whatsappDefaultMessage', 'Hi Diosys, I would like to discuss a project.', 'Prefilled WhatsApp message')
ON DUPLICATE KEY UPDATE `description` = VALUES(`description`);

INSERT INTO `ms_service` (`title`, `description`, `icon`, `orderNo`) VALUES
    ('Custom Web Development', 'High-performance, scalable websites and web applications tailored to your business.', 'globe', 1),
    ('Mobile Applications', 'Native and cross-platform mobile apps with seamless user experiences.', 'smartphone', 2),
    ('AI Integrations', 'Embed intelligent automation and AI capabilities into your products.', 'cpu', 3),
    ('Tech Ecosystem Solutions', 'End-to-end architecture, APIs, and cloud infrastructure for modern teams.', 'layers', 4);
