import { Component, OnInit, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';

import { CmsService } from '../../../core/services/cms.service';
import { Service, ServiceRequest } from '../../../core/models/diosys.model';

const emptyForm = (): ServiceRequest => ({ title: '', description: '', icon: '', orderNo: 0, flagActive: true });

@Component({
  selector: 'app-admin-services',
  standalone: true,
  imports: [FormsModule],
  template: `
    <div class="admin-topbar">
      <h1>Services</h1>
      <button class="btn btn-primary" (click)="startCreate()">+ New service</button>
    </div>

    @if (showForm()) {
      <div class="admin-card">
        <h3 style="margin-bottom:1rem">{{ editingId() ? 'Edit service' : 'New service' }}</h3>
        <div class="form-row">
          <div class="form-group"><label>Title</label><input class="form-control" [(ngModel)]="form.title" /></div>
          <div class="form-group"><label>Icon (globe, smartphone, cpu, layers, code)</label><input class="form-control" [(ngModel)]="form.icon" /></div>
        </div>
        <div class="form-group"><label>Description</label><textarea class="form-control" [(ngModel)]="form.description"></textarea></div>
        <div class="form-row">
          <div class="form-group"><label>Order</label><input class="form-control" type="number" [(ngModel)]="form.orderNo" /></div>
          <div class="form-group"><label>Active</label>
            <select class="form-control" [(ngModel)]="form.flagActive">
              <option [ngValue]="true">Active</option><option [ngValue]="false">Inactive</option>
            </select>
          </div>
        </div>
        <div style="display:flex;gap:0.5rem">
          <button class="btn btn-primary" (click)="save()">Save</button>
          <button class="btn btn-secondary" (click)="showForm.set(false)">Cancel</button>
        </div>
      </div>
    }

    <div class="admin-card">
      <table class="admin-table">
        <thead><tr><th>Title</th><th>Icon</th><th>Order</th><th>Active</th><th>Actions</th></tr></thead>
        <tbody>
          @for (svc of services(); track svc.serviceID) {
            <tr>
              <td><strong>{{ svc.title }}</strong><br /><span class="muted">{{ svc.description }}</span></td>
              <td>{{ svc.icon }}</td>
              <td>{{ svc.orderNo }}</td>
              <td>{{ svc.flagActive ? 'Yes' : 'No' }}</td>
              <td class="actions">
                <button class="btn btn-secondary btn-sm" (click)="startEdit(svc)">Edit</button>
                <button class="btn btn-danger btn-sm" (click)="remove(svc)">Delete</button>
              </td>
            </tr>
          }
          @if (services().length === 0) { <tr><td colspan="5" class="muted">No services.</td></tr> }
        </tbody>
      </table>
    </div>
  `,
})
export class AdminServicesComponent implements OnInit {
  services = signal<Service[]>([]);
  showForm = signal(false);
  editingId = signal<number | null>(null);
  form: ServiceRequest = emptyForm();

  constructor(private cms: CmsService) {}

  ngOnInit(): void { this.load(); }

  load(): void { this.cms.getServices().subscribe((s) => this.services.set(s)); }

  startCreate(): void { this.form = emptyForm(); this.editingId.set(null); this.showForm.set(true); }

  startEdit(svc: Service): void {
    this.form = { title: svc.title, description: svc.description, icon: svc.icon, orderNo: svc.orderNo, flagActive: svc.flagActive === 1 };
    this.editingId.set(svc.serviceID);
    this.showForm.set(true);
  }

  save(): void {
    const id = this.editingId();
    const req = id ? this.cms.updateService(id, this.form) : this.cms.createService(this.form);
    req.subscribe(() => { this.showForm.set(false); this.load(); });
  }

  remove(svc: Service): void {
    if (!confirm(`Delete service "${svc.title}"?`)) return;
    this.cms.deleteService(svc.serviceID).subscribe(() => this.load());
  }
}
