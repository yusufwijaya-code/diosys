package developer_dto

import (
	"portfolio-api/modules/certificate/certificate_model"
	"portfolio-api/modules/education/education_dto"
	"portfolio-api/modules/experience/experience_dto"
	"portfolio-api/modules/professional_project/professional_project_dto"
	"portfolio-api/modules/project/project_dto"
	"portfolio-api/modules/skill/skill_model"
	"portfolio-api/modules/summary/summary_dto"
)

// DeveloperCardResponse is the compact representation for the directory grid.
type DeveloperCardResponse struct {
	UserID   int    `json:"userID"`
	Username string `json:"username"`
	FullName string `json:"fullName"`
	JobTitle string `json:"jobTitle"`
	Intro    string `json:"intro"`
	PhotoUrl string `json:"photoUrl"`
}

// DeveloperResponse is the full developer profile header (no nested portfolio).
type DeveloperResponse struct {
	UserID         int    `json:"userID"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	FullName       string `json:"fullName"`
	JobTitle       string `json:"jobTitle"`
	Intro          string `json:"intro"`
	Bio            string `json:"bio"`
	Specialization string `json:"specialization"`
	Phone          string `json:"phone"`
	Website        string `json:"website"`
	GithubUrl      string `json:"githubUrl"`
	LinkedinUrl    string `json:"linkedinUrl"`
	InstagramUrl   string `json:"instagramUrl"`
	CvUrl          string `json:"cvUrl"`
	Location       string `json:"location"`
	PhotoUrl       string `json:"photoUrl"`
	IsAdmin        int    `json:"isAdmin"`
	FlagActive     int    `json:"flagActive"`
	OrderNo        int    `json:"orderNo"`
}

// DeveloperRequest is the CMS create/update payload for a developer profile.
type DeveloperRequest struct {
	Username       string `json:"username" binding:"required"`
	Email          string `json:"email" binding:"required"`
	FullName       string `json:"fullName" binding:"required"`
	JobTitle       string `json:"jobTitle"`
	Intro          string `json:"intro"`
	Bio            string `json:"bio"`
	Specialization string `json:"specialization"`
	Phone          string `json:"phone"`
	Website        string `json:"website"`
	GithubUrl      string `json:"githubUrl"`
	LinkedinUrl    string `json:"linkedinUrl"`
	InstagramUrl   string `json:"instagramUrl"`
	Location       string `json:"location"`
	FlagActive     *bool  `json:"flagActive"`
	OrderNo        int    `json:"orderNo"`
}

// DeveloperProfileResponse is the aggregated public profile page payload.
type DeveloperProfileResponse struct {
	Developer            DeveloperResponse                                          `json:"developer"`
	Summary              summary_dto.SummaryResponse                                `json:"summary"`
	Experiences          []experience_dto.ExperienceResponse                        `json:"experiences"`
	Educations           []education_dto.EducationResponse                          `json:"educations"`
	Certificates         []certificate_model.Certificate                            `json:"certificates"`
	Skills               []skill_model.Skill                                        `json:"skills"`
	Projects             []project_dto.ProjectResponse                              `json:"projects"`
	ProfessionalProjects []professional_project_dto.ProfessionalProjectCardResponse `json:"professionalProjects"`
}
