import { Component, OnInit, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';

import { CmsService } from '../../../core/services/cms.service';
import { PublicService } from '../../../core/services/public.service';
import { Developer, ProfessionalProject, ProfessionalProjectRequest } from '../../../core/models/diosys.model';

@Component({
  selector: 'app-admin-professional-projects',
  standalone: true,
  imports: [FormsModule],
  templateUrl: './admin-professional-projects.component.html',
})
export class AdminProfessionalProjectsComponent implements OnInit {
  developers = signal<Developer[]>([]);
  selectedUserID = signal<number | null>(null);
  projects = signal<ProfessionalProject[]>([]);

  ppForm: ProfessionalProjectRequest = { title: '', company: '', summary: '', orderNo: 0 };
  ppFeatureForms: Record<number, { title: string; description: string; orderNo: number }> = {};
  ppFeatureImageForms: Record<number, { file: File | null; caption: string }> = {};

  addingProject = signal(false);
  deletingProjectId = signal<number | null>(null);
  uploadingThumbnailId = signal<number | null>(null);
  addingFeatureProjectId = signal<number | null>(null);
  deletingFeatureId = signal<number | null>(null);
  uploadingImageFeatureId = signal<number | null>(null);
  deletingImageId = signal<number | null>(null);

  constructor(private cms: CmsService, private publicSvc: PublicService) {}

  ngOnInit(): void {
    this.cms.getDevelopers().subscribe((d) => {
      this.developers.set(d);
      if (d.length) this.selectDeveloper(d[0].userID);
    });
  }

  selectDeveloper(userID: number): void {
    this.selectedUserID.set(userID);
    this.loadProjects();
  }

  loadProjects(): void {
    const userID = this.selectedUserID();
    if (!userID) return;
    this.cms.getProfessionalProjects(userID).subscribe((cards) => {
      if (!cards.length) { this.projects.set([]); return; }
      const result: ProfessionalProject[] = [];
      let remaining = cards.length;
      cards.forEach((card) => {
        this.publicSvc.getProfessionalProject(card.professionalProjectID).subscribe({
          next: (full) => {
            result.push(full);
            remaining--;
            if (remaining === 0) {
              result.sort((a, b) => a.orderNo - b.orderNo);
              this.projects.set(result);
            }
          },
          error: () => {
            remaining--;
            if (remaining === 0) {
              result.sort((a, b) => a.orderNo - b.orderNo);
              this.projects.set(result);
            }
          },
        });
      });
    });
  }

  addProject(): void {
    const userID = this.selectedUserID();
    if (!userID || !this.ppForm.title || !this.ppForm.company) return;
    this.addingProject.set(true);
    this.cms.createProfessionalProject(userID, { ...this.ppForm }).subscribe({
      next: () => {
        this.addingProject.set(false);
        this.ppForm = { title: '', company: '', summary: '', orderNo: 0 };
        this.loadProjects();
      },
      error: () => this.addingProject.set(false),
    });
  }

  deleteProject(projectID: number): void {
    const userID = this.selectedUserID();
    if (!userID || !confirm('Delete this professional project?')) return;
    this.deletingProjectId.set(projectID);
    this.cms.deleteProfessionalProject(userID, projectID).subscribe({
      next: () => { this.deletingProjectId.set(null); this.loadProjects(); },
      error: () => this.deletingProjectId.set(null),
    });
  }

  onThumbnail(projectID: number, event: Event): void {
    const userID = this.selectedUserID();
    const file = (event.target as HTMLInputElement).files?.[0];
    if (!userID || !file) return;
    this.uploadingThumbnailId.set(projectID);
    this.cms.uploadProfessionalProjectThumbnail(userID, projectID, file).subscribe({
      next: () => { this.uploadingThumbnailId.set(null); this.loadProjects(); },
      error: () => this.uploadingThumbnailId.set(null),
    });
  }

  getFeatureForm(projectID: number): { title: string; description: string; orderNo: number } {
    if (!this.ppFeatureForms[projectID]) {
      this.ppFeatureForms[projectID] = { title: '', description: '', orderNo: 0 };
    }
    return this.ppFeatureForms[projectID];
  }

  addFeature(proj: ProfessionalProject): void {
    const userID = this.selectedUserID();
    const form = this.getFeatureForm(proj.professionalProjectID);
    if (!userID || !form.title) return;
    this.addingFeatureProjectId.set(proj.professionalProjectID);
    this.cms.addProfessionalProjectFeature(userID, proj.professionalProjectID, { ...form }).subscribe({
      next: () => {
        this.addingFeatureProjectId.set(null);
        this.ppFeatureForms[proj.professionalProjectID] = { title: '', description: '', orderNo: 0 };
        this.loadProjects();
      },
      error: () => this.addingFeatureProjectId.set(null),
    });
  }

  deleteFeature(proj: ProfessionalProject, featureID: number): void {
    const userID = this.selectedUserID();
    if (!userID) return;
    this.deletingFeatureId.set(featureID);
    this.cms.deleteProfessionalProjectFeature(userID, proj.professionalProjectID, featureID).subscribe({
      next: () => { this.deletingFeatureId.set(null); this.loadProjects(); },
      error: () => this.deletingFeatureId.set(null),
    });
  }

  getImageForm(featureID: number): { file: File | null; caption: string } {
    if (!this.ppFeatureImageForms[featureID]) {
      this.ppFeatureImageForms[featureID] = { file: null, caption: '' };
    }
    return this.ppFeatureImageForms[featureID];
  }

  onImageFile(featureID: number, event: Event): void {
    const file = (event.target as HTMLInputElement).files?.[0] ?? null;
    this.getImageForm(featureID).file = file;
  }

  uploadImage(proj: ProfessionalProject, featureID: number): void {
    const userID = this.selectedUserID();
    const form = this.getImageForm(featureID);
    if (!userID || !form.file) return;
    this.uploadingImageFeatureId.set(featureID);
    this.cms.addFeatureImage(userID, proj.professionalProjectID, featureID, form.file, form.caption).subscribe({
      next: () => {
        this.uploadingImageFeatureId.set(null);
        this.ppFeatureImageForms[featureID] = { file: null, caption: '' };
        this.loadProjects();
      },
      error: () => this.uploadingImageFeatureId.set(null),
    });
  }

  deleteImage(proj: ProfessionalProject, featureID: number, imageID: number): void {
    const userID = this.selectedUserID();
    if (!userID) return;
    this.deletingImageId.set(imageID);
    this.cms.deleteFeatureImage(userID, proj.professionalProjectID, featureID, imageID).subscribe({
      next: () => { this.deletingImageId.set(null); this.loadProjects(); },
      error: () => this.deletingImageId.set(null),
    });
  }
}
