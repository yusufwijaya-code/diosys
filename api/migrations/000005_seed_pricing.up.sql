-- Sample "All-In" pricing plan ($195 flat, all features included).
INSERT INTO `ms_price_plan`
    (`title`, `subtitle`, `price`, `currency`, `billingPeriod`, `badge`, `isFeatured`, `flagActive`, `orderNo`)
VALUES
    ('All-In', 'One flat price. Every feature. No hidden costs.',
     195, 'USD', '/ project', 'Best Value', 1, 1, 1);

SET @planID = LAST_INSERT_ID();

INSERT INTO `ms_price_feature` (`planID`, `text`, `isIncluded`, `orderNo`) VALUES
    (@planID, 'Custom Website Development',   1, 1),
    (@planID, 'Mobile Application (iOS & Android)', 1, 2),
    (@planID, 'AI Integration & Automation',  1, 3),
    (@planID, 'Admin Dashboard / CMS',         1, 4),
    (@planID, 'REST API Backend',              1, 5),
    (@planID, 'Google Drive / Cloud Storage',  1, 6),
    (@planID, 'Google OAuth Login',            1, 7),
    (@planID, '3 Revision Rounds',             1, 8),
    (@planID, 'Deployment & Hosting Setup',    1, 9),
    (@planID, 'Post-launch Support (30 days)', 1, 10);
