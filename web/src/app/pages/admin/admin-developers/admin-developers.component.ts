import { Component, OnInit, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { RouterLink } from '@angular/router';

import { CmsService } from '../../../core/services/cms.service';
import { Developer, DeveloperRequest } from '../../../core/models/diosys.model';

const emptyForm = (): DeveloperRequest => ({
  username: '', email: '', fullName: '', jobTitle: '', intro: '', bio: '',
  specialization: '', phone: '', website: '', githubUrl: '', linkedinUrl: '',
  instagramUrl: '', location: '', flagActive: true, orderNo: 0,
});

@Component({
  selector: 'app-admin-developers',
  standalone: true,
  imports: [FormsModule, RouterLink],
  templateUrl: './admin-developers.component.html',
})
export class AdminDevelopersComponent implements OnInit {
  developers = signal<Developer[]>([]);
  editingId = signal<number | null>(null);
  showForm = signal(false);
  saving = signal(false);
  deletingId = signal<number | null>(null);
  uploadingPhoto = signal(false);
  uploadingCV = signal(false);
  error = signal('');
  form: DeveloperRequest = emptyForm();

  constructor(private cms: CmsService) {}

  ngOnInit(): void {
    this.load();
  }

  load(): void {
    this.cms.getDevelopers().subscribe((d) => this.developers.set(d));
  }

  startCreate(): void {
    this.form = emptyForm();
    this.editingId.set(null);
    this.showForm.set(true);
    this.error.set('');
  }

  startEdit(dev: Developer): void {
    this.form = {
      username: dev.username, email: dev.email, fullName: dev.fullName,
      jobTitle: dev.jobTitle, intro: dev.intro, bio: dev.bio,
      specialization: dev.specialization, phone: dev.phone, website: dev.website,
      githubUrl: dev.githubUrl || '', linkedinUrl: dev.linkedinUrl || '',
      instagramUrl: dev.instagramUrl || '', location: dev.location,
      flagActive: dev.flagActive === 1, orderNo: dev.orderNo,
    };
    this.editingId.set(dev.userID);
    this.showForm.set(true);
    this.error.set('');
  }

  cancel(): void {
    this.showForm.set(false);
  }

  save(): void {
    this.saving.set(true);
    this.error.set('');
    const id = this.editingId();
    const request = id
      ? this.cms.updateDeveloper(id, this.form)
      : this.cms.createDeveloper(this.form);
    request.subscribe({
      next: () => {
        this.saving.set(false);
        this.showForm.set(false);
        this.load();
      },
      error: (err) => {
        this.saving.set(false);
        this.error.set(err?.error?.message || 'Failed to save developer.');
      },
    });
  }

  remove(dev: Developer): void {
    if (!confirm(`Delete developer "${dev.fullName}"? This removes their portfolio too.`)) return;
    this.deletingId.set(dev.userID);
    this.cms.deleteDeveloper(dev.userID).subscribe({
      next: () => { this.deletingId.set(null); this.load(); },
      error: (err) => { this.deletingId.set(null); alert(err?.error?.message || 'Failed to delete.'); },
    });
  }

  uploadPhoto(event: Event): void {
    const id = this.editingId();
    const file = (event.target as HTMLInputElement).files?.[0];
    if (!id || !file) return;
    this.uploadingPhoto.set(true);
    this.cms.uploadDeveloperPhoto(id, file).subscribe({
      next: () => { this.uploadingPhoto.set(false); this.load(); },
      error: () => { this.uploadingPhoto.set(false); alert('Failed to upload photo.'); },
    });
  }

  uploadCV(event: Event): void {
    const id = this.editingId();
    const file = (event.target as HTMLInputElement).files?.[0];
    if (!id || !file) return;
    this.uploadingCV.set(true);
    this.cms.uploadCV(id, file).subscribe({
      next: () => { this.uploadingCV.set(false); this.load(); },
      error: () => { this.uploadingCV.set(false); alert('Failed to upload CV.'); },
    });
  }
}
