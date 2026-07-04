ALTER TABLE ms_user
  ADD COLUMN githubUrl VARCHAR(255) NULL AFTER website,
  ADD COLUMN linkedinUrl VARCHAR(255) NULL AFTER githubUrl,
  ADD COLUMN instagramUrl VARCHAR(255) NULL AFTER linkedinUrl;
