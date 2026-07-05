import { Component, OnInit, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';

import { CmsService } from '../../../core/services/cms.service';
import { Testimonial, TestimonialRequest } from '../../../core/models/diosys.model';

const emptyForm = (): TestimonialRequest => ({
  clientName: '', clientRole: '', clientCompany: '',
  testimonialText: '', rating: 5, orderNo: 0, flagActive: true,
});

@Component({
  selector: 'app-admin-testimonials',
  standalone: true,
  imports: [FormsModule],
  templateUrl: './admin-testimonials.component.html',
})
export class AdminTestimonialsComponent implements OnInit {
  items = signal<Testimonial[]>([]);
  showForm = signal(false);
  editingId = signal<number | null>(null);
  saving = signal(false);
  uploading = signal(false);
  deletingId = signal<number | null>(null);
  error = signal('');

  form: TestimonialRequest = emptyForm();
  pendingPhoto: File | null = null;
  photoPreview = '';

  readonly stars = [1, 2, 3, 4, 5];

  constructor(private cms: CmsService) {}

  ngOnInit(): void { this.load(); }

  load(): void {
    this.cms.getTestimonials().subscribe((t) => this.items.set(t));
  }

  startCreate(): void {
    this.form = emptyForm();
    this.pendingPhoto = null;
    this.photoPreview = '';
    this.editingId.set(null);
    this.showForm.set(true);
    this.error.set('');
  }

  startEdit(t: Testimonial): void {
    this.form = {
      clientName: t.clientName, clientRole: t.clientRole,
      clientCompany: t.clientCompany, testimonialText: t.testimonialText,
      rating: t.rating, orderNo: t.orderNo, flagActive: t.flagActive === 1,
    };
    this.pendingPhoto = null;
    this.photoPreview = t.photoUrl || '';
    this.editingId.set(t.testimonialID);
    this.showForm.set(true);
    this.error.set('');
  }

  cancel(): void { this.showForm.set(false); }

  onFileChange(event: Event): void {
    const file = (event.target as HTMLInputElement).files?.[0];
    if (!file) return;
    this.pendingPhoto = file;
    const reader = new FileReader();
    reader.onload = (e) => this.photoPreview = e.target?.result as string;
    reader.readAsDataURL(file);
  }

  setRating(star: number): void { this.form.rating = star; }

  save(): void {
    if (!this.form.clientName.trim() || !this.form.testimonialText.trim()) {
      this.error.set('Client name and testimonial text are required.');
      return;
    }
    this.saving.set(true);
    this.error.set('');
    const id = this.editingId();
    const req = id
      ? this.cms.updateTestimonial(id, this.form)
      : this.cms.createTestimonial(this.form);

    req.subscribe({
      next: (saved) => {
        if (this.pendingPhoto) {
          this.uploading.set(true);
          this.cms.uploadTestimonialPhoto(saved.testimonialID, this.pendingPhoto).subscribe({
            next: () => { this.saving.set(false); this.uploading.set(false); this.showForm.set(false); this.load(); },
            error: () => { this.saving.set(false); this.uploading.set(false); this.showForm.set(false); this.load(); },
          });
        } else {
          this.saving.set(false);
          this.showForm.set(false);
          this.load();
        }
      },
      error: (err) => {
        this.saving.set(false);
        this.error.set(err?.error?.message || 'Failed to save.');
      },
    });
  }

  remove(t: Testimonial): void {
    if (!confirm(`Delete testimonial from "${t.clientName}"?`)) return;
    this.deletingId.set(t.testimonialID);
    this.cms.deleteTestimonial(t.testimonialID).subscribe({
      next: () => { this.deletingId.set(null); this.load(); },
      error: () => this.deletingId.set(null),
    });
  }
}
