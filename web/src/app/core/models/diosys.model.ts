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

export interface Language {
  languageID: number;
  name: string;
  level: string;
  icon: string;
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
  features: string[];
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
  languages: Language[];
  projects: Project[];
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

/* ---------- Pricing ---------- */

export interface PriceFeature {
  featureID: number;
  text: string;
  isIncluded: boolean;
  orderNo: number;
}

export interface PricePlan {
  planID: number;
  title: string;
  subtitle: string;
  price: number;
  currency: string;
  billingPeriod: string;
  badge: string;
  originalPrice: number | null;
  discountPercent: number | null;
  isFeatured: boolean;
  flagActive: boolean;
  orderNo: number;
  features: PriceFeature[];
}

export interface PricePlanRequest {
  title: string;
  subtitle: string;
  price: number;
  currency: string;
  billingPeriod: string;
  badge: string;
  originalPrice: number | null;
  discountPercent: number | null;
  isFeatured: boolean;
  flagActive: boolean;
  orderNo: number;
  features: { text: string; isIncluded: boolean }[];
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

export interface LanguageRequest {
  name: string;
  level: string;
  icon: string;
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
