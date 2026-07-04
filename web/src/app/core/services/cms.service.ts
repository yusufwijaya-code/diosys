import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { ApiService } from './api.service';
import {
  Certificate,
  PricePlan,
  PricePlanRequest,
  CertificateRequest,
  ClientMessage,
  Developer,
  DeveloperRequest,
  Education,
  EducationRequest,
  Experience,
  ExperienceRequest,
  MessageStatusRequest,
  ProfessionalProject,
  ProfessionalProjectCard,
  ProfessionalProjectFeature,
  ProfessionalProjectRequest,
  ProjectFeatureRequest,
  Project,
  ProjectFeatureItem,
  ProjectRequest,
  Service,
  ServiceRequest,
  Skill,
  SkillRequest,
  Summary,
  SummaryRequest,
  SettingsMap,
} from '../models/diosys.model';

/** CMS (authenticated) data access for managing Diosys content. */
@Injectable({ providedIn: 'root' })
export class CmsService {
  constructor(private api: ApiService) {}

  // ---------- Developers ----------
  getDevelopers(): Observable<Developer[]> {
    return this.api.get<Developer[]>('/cms/developers');
  }
  getDeveloper(userID: number): Observable<Developer> {
    return this.api.get<Developer>(`/cms/developers/${userID}`);
  }
  createDeveloper(body: DeveloperRequest): Observable<Developer> {
    return this.api.post<Developer>('/cms/developers', body);
  }
  updateDeveloper(userID: number, body: DeveloperRequest): Observable<Developer> {
    return this.api.put<Developer>(`/cms/developers/${userID}`, body);
  }
  deleteDeveloper(userID: number): Observable<unknown> {
    return this.api.delete<unknown>(`/cms/developers/${userID}`);
  }
  uploadDeveloperPhoto(userID: number, file: File): Observable<Developer> {
    return this.api.upload<Developer>(`/cms/developers/${userID}/photo`, file);
  }

  // ---------- Per-developer portfolio ----------
  getSummary(userID: number): Observable<Summary> {
    return this.api.get<Summary>(`/cms/developers/${userID}/summary`);
  }
  saveSummary(userID: number, body: SummaryRequest): Observable<Summary> {
    return this.api.put<Summary>(`/cms/developers/${userID}/summary`, body);
  }

  getExperiences(userID: number): Observable<Experience[]> {
    return this.api.get<Experience[]>(`/cms/developers/${userID}/experiences`);
  }
  createExperience(userID: number, body: ExperienceRequest): Observable<Experience> {
    return this.api.post<Experience>(`/cms/developers/${userID}/experiences`, body);
  }
  updateExperience(userID: number, id: number, body: ExperienceRequest): Observable<Experience> {
    return this.api.put<Experience>(`/cms/developers/${userID}/experiences/${id}`, body);
  }
  deleteExperience(userID: number, id: number): Observable<unknown> {
    return this.api.delete<unknown>(`/cms/developers/${userID}/experiences/${id}`);
  }
  uploadCV(userID: number, file: File): Observable<Developer> {
    const fd = new FormData();
    fd.append('file', file);
    return this.api.post<Developer>(`/cms/developers/${userID}/cv`, fd);
  }

  // ---------- Professional projects ----------
  getProfessionalProjects(userID: number): Observable<ProfessionalProjectCard[]> {
    return this.api.get<ProfessionalProjectCard[]>(`/cms/developers/${userID}/professional-projects`);
  }
  createProfessionalProject(userID: number, body: ProfessionalProjectRequest): Observable<ProfessionalProject> {
    return this.api.post<ProfessionalProject>(`/cms/developers/${userID}/professional-projects`, body);
  }
  deleteProfessionalProject(userID: number, projectID: number): Observable<unknown> {
    return this.api.delete<unknown>(`/cms/developers/${userID}/professional-projects/${projectID}`);
  }
  uploadProfessionalProjectThumbnail(userID: number, projectID: number, file: File): Observable<ProfessionalProject> {
    return this.api.upload<ProfessionalProject>(`/cms/developers/${userID}/professional-projects/${projectID}/thumbnail`, file);
  }
  addProfessionalProjectFeature(userID: number, projectID: number, body: ProjectFeatureRequest): Observable<ProfessionalProject> {
    return this.api.post<ProfessionalProject>(`/cms/developers/${userID}/professional-projects/${projectID}/features`, body);
  }
  deleteProfessionalProjectFeature(userID: number, projectID: number, featureID: number): Observable<unknown> {
    return this.api.delete<unknown>(`/cms/developers/${userID}/professional-projects/${projectID}/features/${featureID}`);
  }
  addFeatureImage(userID: number, projectID: number, featureID: number, file: File, caption: string): Observable<ProfessionalProjectFeature> {
    return this.api.uploadWithFields<ProfessionalProjectFeature>(
      `/cms/developers/${userID}/professional-projects/${projectID}/features/${featureID}/images`,
      file, { caption });
  }
  deleteFeatureImage(userID: number, projectID: number, featureID: number, imageID: number): Observable<unknown> {
    return this.api.delete<unknown>(`/cms/developers/${userID}/professional-projects/${projectID}/features/${featureID}/images/${imageID}`);
  }

  getEducations(userID: number): Observable<Education[]> {
    return this.api.get<Education[]>(`/cms/developers/${userID}/educations`);
  }
  createEducation(userID: number, body: EducationRequest): Observable<Education> {
    return this.api.post<Education>(`/cms/developers/${userID}/educations`, body);
  }
  updateEducation(userID: number, id: number, body: EducationRequest): Observable<Education> {
    return this.api.put<Education>(`/cms/developers/${userID}/educations/${id}`, body);
  }
  deleteEducation(userID: number, id: number): Observable<unknown> {
    return this.api.delete<unknown>(`/cms/developers/${userID}/educations/${id}`);
  }

  getCertificates(userID: number): Observable<Certificate[]> {
    return this.api.get<Certificate[]>(`/cms/developers/${userID}/certificates`);
  }
  createCertificate(userID: number, body: CertificateRequest): Observable<Certificate> {
    return this.api.post<Certificate>(`/cms/developers/${userID}/certificates`, body);
  }
  updateCertificate(userID: number, id: number, body: CertificateRequest): Observable<Certificate> {
    return this.api.put<Certificate>(`/cms/developers/${userID}/certificates/${id}`, body);
  }
  deleteCertificate(userID: number, id: number): Observable<unknown> {
    return this.api.delete<unknown>(`/cms/developers/${userID}/certificates/${id}`);
  }

  getSkills(userID: number): Observable<Skill[]> {
    return this.api.get<Skill[]>(`/cms/developers/${userID}/skills`);
  }
  createSkill(userID: number, body: SkillRequest): Observable<Skill> {
    return this.api.post<Skill>(`/cms/developers/${userID}/skills`, body);
  }
  updateSkill(userID: number, id: number, body: SkillRequest): Observable<Skill> {
    return this.api.put<Skill>(`/cms/developers/${userID}/skills/${id}`, body);
  }
  deleteSkill(userID: number, id: number): Observable<unknown> {
    return this.api.delete<unknown>(`/cms/developers/${userID}/skills/${id}`);
  }

  // ---------- Projects ----------
  getProjects(userID: number): Observable<Project[]> {
    return this.api.get<Project[]>(`/cms/developers/${userID}/projects`);
  }
  createProject(userID: number, body: ProjectRequest): Observable<Project> {
    return this.api.post<Project>(`/cms/developers/${userID}/projects`, body);
  }
  updateProject(userID: number, id: number, body: ProjectRequest): Observable<Project> {
    return this.api.put<Project>(`/cms/developers/${userID}/projects/${id}`, body);
  }
  deleteProject(userID: number, id: number): Observable<unknown> {
    return this.api.delete<unknown>(`/cms/developers/${userID}/projects/${id}`);
  }
  uploadProjectThumbnail(userID: number, id: number, file: File): Observable<Project> {
    return this.api.upload<Project>(`/cms/developers/${userID}/projects/${id}/thumbnail`, file);
  }
  addProjectImage(userID: number, id: number, file: File, caption: string): Observable<Project> {
    return this.api.uploadWithFields<Project>(
      `/cms/developers/${userID}/projects/${id}/images`,
      file,
      { caption }
    );
  }
  deleteProjectImage(userID: number, id: number, imageID: number): Observable<unknown> {
    return this.api.delete<unknown>(
      `/cms/developers/${userID}/projects/${id}/images/${imageID}`
    );
  }
  addProjectFeatureImage(userID: number, projectID: number, featureID: number, file: File, caption: string): Observable<Project> {
    return this.api.uploadWithFields<Project>(
      `/cms/developers/${userID}/projects/${projectID}/features/${featureID}/images`,
      file,
      { caption }
    );
  }
  deleteProjectFeatureImage(userID: number, projectID: number, featureID: number, imageID: number): Observable<unknown> {
    return this.api.delete<unknown>(
      `/cms/developers/${userID}/projects/${projectID}/features/${featureID}/images/${imageID}`
    );
  }

  // ---------- Services ----------
  getServices(): Observable<Service[]> {
    return this.api.get<Service[]>('/cms/services');
  }
  createService(body: ServiceRequest): Observable<Service> {
    return this.api.post<Service>('/cms/services', body);
  }
  updateService(id: number, body: ServiceRequest): Observable<Service> {
    return this.api.put<Service>(`/cms/services/${id}`, body);
  }
  deleteService(id: number): Observable<unknown> {
    return this.api.delete<unknown>(`/cms/services/${id}`);
  }

  // ---------- Messages (inbox) ----------
  getMessages(): Observable<ClientMessage[]> {
    return this.api.get<ClientMessage[]>('/cms/messages');
  }
  updateMessageStatus(id: number, body: MessageStatusRequest): Observable<ClientMessage> {
    return this.api.patch<ClientMessage>(`/cms/messages/${id}/status`, body);
  }
  deleteMessage(id: number): Observable<unknown> {
    return this.api.delete<unknown>(`/cms/messages/${id}`);
  }

  // ---------- Settings ----------
  // Current values are read from the public settings map; this saves changes.
  updateSettings(settings: { settingKey: string; settingValue: string }[]): Observable<SettingsMap> {
    return this.api.put<SettingsMap>('/cms/settings', { settings });
  }

  // ---------- Pricing ----------
  getPricingPlans(): Observable<PricePlan[]> {
    return this.api.get<PricePlan[]>('/cms/pricing');
  }
  createPricingPlan(body: PricePlanRequest): Observable<PricePlan> {
    return this.api.post<PricePlan>('/cms/pricing', body);
  }
  updatePricingPlan(id: number, body: PricePlanRequest): Observable<PricePlan> {
    return this.api.put<PricePlan>(`/cms/pricing/${id}`, body);
  }
  deletePricingPlan(id: number): Observable<unknown> {
    return this.api.delete<unknown>(`/cms/pricing/${id}`);
  }
}
