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
    this.cms.createProfessionalProject(userID, { ...this.ppForm }).subscribe(() => {
      this.ppForm = { title: '', company: '', summary: '', orderNo: 0 };
      this.loadProjects();
    });
  }

  deleteProject(projectID: number): void {
    const userID = this.selectedUserID();
    if (!userID || !confirm('Delete this professional project?')) return;
    this.cms.deleteProfessionalProject(userID, projectID).subscribe(() => this.loadProjects());
  }

  onThumbnail(projectID: number, event: Event): void {
    const userID = this.selectedUserID();
    const file = (event.target as HTMLInputElement).files?.[0];
    if (!userID || !file) return;
    this.cms.uploadProfessionalProjectThumbnail(userID, projectID, file).subscribe(() => this.loadProjects());
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
    this.cms.addProfessionalProjectFeature(userID, proj.professionalProjectID, { ...form }).subscribe(() => {
      this.ppFeatureForms[proj.professionalProjectID] = { title: '', description: '', orderNo: 0 };
      this.loadProjects();
    });
  }

  deleteFeature(proj: ProfessionalProject, featureID: number): void {
    const userID = this.selectedUserID();
    if (!userID) return;
    this.cms.deleteProfessionalProjectFeature(userID, proj.professionalProjectID, featureID)
      .subscribe(() => this.loadProjects());
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
    this.cms.addFeatureImage(userID, proj.professionalProjectID, featureID, form.file, form.caption)
      .subscribe(() => {
        this.ppFeatureImageForms[featureID] = { file: null, caption: '' };
        this.loadProjects();
      });
  }

  deleteImage(proj: ProfessionalProject, featureID: number, imageID: number): void {
    const userID = this.selectedUserID();
    if (!userID) return;
    this.cms.deleteFeatureImage(userID, proj.professionalProjectID, featureID, imageID)
      .subscribe(() => this.loadProjects());
  }
}
