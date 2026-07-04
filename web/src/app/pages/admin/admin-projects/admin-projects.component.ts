import { Component, OnInit, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';

import { CmsService } from '../../../core/services/cms.service';
import { Developer, Project, ProjectRequest } from '../../../core/models/diosys.model';

interface ProjectForm {
  title: string;
  summary: string;
  body: string;
  client: string;
  projectLink: string;
  repoLink: string;
  projectStatusID: number | null;
  isFeatured: boolean;
  orderNo: number;
  features: string;      // newline separated
  technologies: string;  // comma separated
}

const emptyForm = (): ProjectForm => ({
  title: '', summary: '', body: '', client: '', projectLink: '', repoLink: '',
  projectStatusID: null, isFeatured: false, orderNo: 0, features: '', technologies: '',
});

@Component({
  selector: 'app-admin-projects',
  standalone: true,
  imports: [FormsModule],
  templateUrl: './admin-projects.component.html',
})
export class AdminProjectsComponent implements OnInit {
  developers = signal<Developer[]>([]);
  selectedUserID = signal<number | null>(null);
  projects = signal<Project[]>([]);

  showForm = signal(false);
  editingId = signal<number | null>(null);
  saving = signal(false);
  error = signal('');
  form: ProjectForm = emptyForm();

  statuses = [
    { id: 1, name: 'Completed' },
    { id: 2, name: 'Ongoing' },
    { id: 3, name: 'Maintenance' },
  ];

  constructor(private cms: CmsService) {}

  ngOnInit(): void {
    this.cms.getDevelopers().subscribe((d) => {
      this.developers.set(d);
      if (d.length) this.selectDeveloper(d[0].userID);
    });
  }

  selectDeveloper(userID: number): void {
    this.selectedUserID.set(userID);
    this.showForm.set(false);
    this.loadProjects();
  }

  loadProjects(): void {
    const id = this.selectedUserID();
    if (!id) return;
    this.cms.getProjects(id).subscribe((p) => this.projects.set(p));
  }

  private toRequest(): ProjectRequest {
    return {
      title: this.form.title, summary: this.form.summary, body: this.form.body,
      client: this.form.client, projectLink: this.form.projectLink, repoLink: this.form.repoLink,
      projectStatusID: this.form.projectStatusID ? Number(this.form.projectStatusID) : null,
      isFeatured: this.form.isFeatured, orderNo: this.form.orderNo,
      features: this.form.features.split('\n').map((v) => v.trim()).filter(Boolean),
      technologies: this.form.technologies.split(',').map((v) => v.trim()).filter(Boolean),
    };
  }

  startCreate(): void {
    this.form = emptyForm();
    this.editingId.set(null);
    this.showForm.set(true);
    this.error.set('');
  }

  startEdit(project: Project): void {
    this.form = {
      title: project.title, summary: project.summary, body: project.body, client: project.client,
      projectLink: project.projectLink, repoLink: project.repoLink,
      projectStatusID: project.projectStatusID, isFeatured: project.isFeatured,
      orderNo: project.orderNo,
      features: project.features.map(f => f.text).join('\n'),
      technologies: project.technologies.join(', '),
    };
    this.editingId.set(project.projectID);
    this.showForm.set(true);
    this.error.set('');
  }

  cancel(): void { this.showForm.set(false); }

  save(): void {
    const userID = this.selectedUserID();
    if (!userID) return;
    this.saving.set(true);
    this.error.set('');
    const id = this.editingId();
    const req = id
      ? this.cms.updateProject(userID, id, this.toRequest())
      : this.cms.createProject(userID, this.toRequest());
    req.subscribe({
      next: () => { this.saving.set(false); this.showForm.set(false); this.loadProjects(); },
      error: (err) => { this.saving.set(false); this.error.set(err?.error?.message || 'Failed to save project.'); },
    });
  }

  remove(project: Project): void {
    const userID = this.selectedUserID();
    if (!userID || !confirm(`Delete project "${project.title}"?`)) return;
    this.cms.deleteProject(userID, project.projectID).subscribe(() => this.loadProjects());
  }

  uploadThumb(project: Project, event: Event): void {
    const userID = this.selectedUserID();
    const file = (event.target as HTMLInputElement).files?.[0];
    if (!userID || !file) return;
    this.cms.uploadProjectThumbnail(userID, project.projectID, file).subscribe(() => this.loadProjects());
  }

  addImage(project: Project, event: Event): void {
    const userID = this.selectedUserID();
    const file = (event.target as HTMLInputElement).files?.[0];
    if (!userID || !file) return;
    const caption = prompt('Image caption (optional):') || '';
    this.cms.addProjectImage(userID, project.projectID, file, caption).subscribe(() => this.loadProjects());
  }

  deleteImage(project: Project, imageID: number): void {
    const userID = this.selectedUserID();
    if (!userID) return;
    this.cms.deleteProjectImage(userID, project.projectID, imageID).subscribe(() => this.loadProjects());
  }

  addFeatureImage(project: Project, featureID: number, event: Event): void {
    const userID = this.selectedUserID();
    const file = (event.target as HTMLInputElement).files?.[0];
    if (!userID || !file) return;
    const caption = prompt('Image caption (optional):') || '';
    this.cms.addProjectFeatureImage(userID, project.projectID, featureID, file, caption)
      .subscribe(() => this.loadProjects());
  }

  deleteFeatureImage(project: Project, featureID: number, imageID: number): void {
    const userID = this.selectedUserID();
    if (!userID) return;
    this.cms.deleteProjectFeatureImage(userID, project.projectID, featureID, imageID)
      .subscribe(() => this.loadProjects());
  }
}
