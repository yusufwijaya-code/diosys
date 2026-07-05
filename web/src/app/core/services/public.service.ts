import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { ApiService } from './api.service';
import {
  ClientMessage,
  DeveloperCard,
  DeveloperProfile,
  MessageRequest,
  ProfessionalProject,
  Project,
  Service,
  SettingsMap,
  Testimonial,
} from '../models/diosys.model';

/** Read-only public data access for the Diosys website. */
@Injectable({ providedIn: 'root' })
export class PublicService {
  constructor(private api: ApiService) {}

  getServices(): Observable<Service[]> {
    return this.api.get<Service[]>('/public/services');
  }

  getDevelopers(): Observable<DeveloperCard[]> {
    return this.api.get<DeveloperCard[]>('/public/developers');
  }

  getDeveloperProfile(username: string): Observable<DeveloperProfile> {
    return this.api.get<DeveloperProfile>(`/public/developers/${username}`);
  }

  getProjects(): Observable<Project[]> {
    return this.api.get<Project[]>('/public/projects');
  }

  getProject(id: number): Observable<Project> {
    return this.api.get<Project>(`/public/projects/${id}`);
  }

  getProfessionalProject(id: number): Observable<ProfessionalProject> {
    return this.api.get<ProfessionalProject>(`/public/professional-projects/${id}`);
  }

  getSettings(): Observable<SettingsMap> {
    return this.api.get<SettingsMap>('/public/settings');
  }

  getTestimonials(): Observable<Testimonial[]> {
    return this.api.get<Testimonial[]>('/public/testimonials');
  }

  sendContact(body: MessageRequest): Observable<ClientMessage> {
    return this.api.post<ClientMessage>('/public/contact', body);
  }
}
