/** Shared data contracts mirroring the Diosys backend DTOs. */

export interface DeveloperCard {
  userID: number;
  username: string;
  fullName: string;
  jobTitle: string;
  intro: string;
  photoUrl: string;
}

export interface Developer {
  userID: number;
  username: string;
  email: string;
  fullName: string;
  jobTitle: string;
  intro: string;
  bio: string;
  specialization: string;
  phone: string;
  website: string;
  githubUrl: string;
  linkedinUrl: string;
  instagramUrl: string;
  cvUrl: string;
  location: string;
  photoUrl: string;
  isAdmin: number;
  flagActive: number;
  orderNo: number;
}

export interface SummaryStat {
  number: string;
  label: string;
}

export interface SummaryFact {
  icon: string;
  text: string;
}

export interface Summary {
  summaryID: number;
  userID: number;
  content: string;
  stats: SummaryStat[];
  facts: SummaryFact[];
}

export interface ProfessionalProjectFeatureImage {
  featureImageID: number;
  url: string;
  caption: string;
  orderNo: number;
}

export interface ProfessionalProjectFeature {
  featureID: number;
  title: string;
  description: string;
  images: ProfessionalProjectFeatureImage[];
  orderNo: number;
}

export interface ProfessionalProjectCard {
  professionalProjectID: number;
  title: string;
  company: string;
  summary: string;
  thumbnailUrl: string;
  orderNo: number;
}

export interface ProfessionalProject {
  professionalProjectID: number;
  userID: number;
  title: string;
  company: string;
  summary: string;
  thumbnailUrl: string;
  features: ProfessionalProjectFeature[];
  orderNo: number;
}

export interface Experience {
  experienceID: number;
  position: string;
  company: string;
  period: string;
  orderNo: number;
  technologies: string[];
  responsibilities: string[];
}

export interface Education {
  educationID: number;
  degree: string;
  institution: string;
  year: string;
  type: string;
  orderNo: number;
  achievements: string[];
}

export interface Certificate {
  certificateID: number;
  name: string;
  issuer: string;
  period: string;
  link: string | null;
  orderNo: number;
}

export interface Skill {
  skillID: number;
  name: string;
  level: string;
  category: string;
  orderNo: number;
}


export interface ProjectImage {
  projectImageID: number;
  fileName: string;
  gdriveID: string;
  url: string;
  caption: string;
  displayOrder: number;
}

export interface ProjectFeatureImage {
  projectFeatureImageID: number;
  url: string;
  caption: string;
  orderNo: number;
}

export interface ProjectFeatureItem {
  projectFeatureID: number;
  text: string;
  description: string;
  images: ProjectFeatureImage[];
  orderNo: number;
}

export interface Project {
  projectID: number;
  userID: number;
  ownerUsername: string;
  ownerFullName: string;
  title: string;
  summary: string;
  body: string;
  client: string;
  projectLink: string;
  repoLink: string;
  projectStatusID: number | null;
  isFeatured: boolean;
  thumbnailFileName: string;
  thumbnailGdriveID: string;
  thumbnailUrl: string;
  orderNo: number;
  features: ProjectFeatureItem[];
  technologies: string[];
  images: ProjectImage[];
}

export interface DeveloperProfile {
  developer: Developer;
  summary: Summary;
  experiences: Experience[];
  educations: Education[];
  certificates: Certificate[];
  skills: Skill[];
  projects: Project[];
  professionalProjects: ProfessionalProjectCard[];
}

export interface Service {
  serviceID: number;
  title: string;
  description: string;
  icon: string;
  orderNo: number;
  flagActive: number;
}

export interface ClientMessage {
  messageID: number;
  clientName: string;
  clientEmail: string;
  clientPhone: string | null;
  subject: string;
  messageBody: string;
  isRead: number;
  isArchived: number;
  createdDate: string;
}

export type SettingsMap = Record<string, string>;

/* ---------- Testimonial ---------- */

export interface Testimonial {
  testimonialID: number;
  clientName: string;
  clientRole: string;
  clientCompany: string;
  testimonialText: string;
  rating: number;
  photoUrl: string;
  flagActive: number;
  orderNo: number;
}

export interface TestimonialRequest {
  clientName: string;
  clientRole: string;
  clientCompany: string;
  testimonialText: string;
  rating: number;
  orderNo: number;
  flagActive: boolean;
}

/* ---------- Request payloads ---------- */

export interface DeveloperRequest {
  username: string;
  email: string;
  fullName: string;
  jobTitle: string;
  intro: string;
  bio: string;
  specialization: string;
  phone: string;
  website: string;
  githubUrl: string;
  linkedinUrl: string;
  instagramUrl: string;
  location: string;
  flagActive: boolean;
  orderNo: number;
}

export interface SummaryRequest {
  content: string;
  stats: SummaryStat[];
  facts: SummaryFact[];
}

export interface ExperienceRequest {
  position: string;
  company: string;
  period: string;
  orderNo: number;
  technologies: string[];
  responsibilities: string[];
}

export interface ProfessionalProjectRequest {
  title: string;
  company: string;
  summary: string;
  orderNo: number;
}

export interface ProjectFeatureRequest {
  title: string;
  description: string;
  orderNo: number;
}

export interface EducationRequest {
  degree: string;
  institution: string;
  year: string;
  type: string;
  orderNo: number;
  achievements: string[];
}

export interface CertificateRequest {
  name: string;
  issuer: string;
  period: string;
  link: string;
  orderNo: number;
}

export interface SkillRequest {
  name: string;
  level: string;
  category: string;
  orderNo: number;
}


export interface ProjectRequest {
  title: string;
  summary: string;
  body: string;
  client: string;
  projectLink: string;
  repoLink: string;
  projectStatusID: number | null;
  isFeatured: boolean;
  orderNo: number;
  features: string[];
  technologies: string[];
}

export interface ServiceRequest {
  title: string;
  description: string;
  icon: string;
  orderNo: number;
  flagActive: boolean;
}

export interface MessageRequest {
  clientName: string;
  clientEmail: string;
  clientPhone: string;
  subject: string;
  messageBody: string;
}

export interface MessageStatusRequest {
  isRead?: boolean;
  isArchived?: boolean;
}
