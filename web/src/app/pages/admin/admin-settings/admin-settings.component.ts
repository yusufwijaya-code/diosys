import { Component, OnInit, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';

import { CmsService } from '../../../core/services/cms.service';
import { PublicService } from '../../../core/services/public.service';

interface SettingRow {
  settingKey: string;
  settingValue: string;
  label: string;
}

const KNOWN: { key: string; label: string }[] = [
  { key: 'companyName', label: 'Company name' },
  { key: 'companyTagline', label: 'Company tagline' },
  { key: 'contactEmail', label: 'Contact email' },
  { key: 'whatsappNumber', label: 'WhatsApp number (e.g. 6281234567890)' },
  { key: 'whatsappDefaultMessage', label: 'WhatsApp default message' },
];

@Component({
  selector: 'app-admin-settings',
  standalone: true,
  imports: [FormsModule],
  template: `
    <div class="admin-topbar"><h1>Settings</h1></div>

    <div class="admin-card">
      @if (saved()) { <div class="alert alert-success">Settings saved.</div> }
      @for (row of rows(); track row.settingKey) {
        <div class="form-group">
          <label>{{ row.label }}</label>
          @if (row.settingKey === 'whatsappDefaultMessage' || row.settingKey === 'companyTagline') {
            <textarea class="form-control" [(ngModel)]="row.settingValue" [name]="row.settingKey"></textarea>
          } @else {
            <input class="form-control" [(ngModel)]="row.settingValue" [name]="row.settingKey" />
          }
        </div>
      }
      <button class="btn btn-primary" (click)="save()" [disabled]="saving()">{{ saving() ? 'Saving…' : 'Save settings' }}</button>
    </div>
  `,
})
export class AdminSettingsComponent implements OnInit {
  rows = signal<SettingRow[]>([]);
  saving = signal(false);
  saved = signal(false);

  constructor(private cms: CmsService, private publicService: PublicService) {}

  ngOnInit(): void {
    this.publicService.getSettings().subscribe((map) => {
      const known = KNOWN.map((k) => ({ settingKey: k.key, label: k.label, settingValue: map[k.key] ?? '' }));
      const extras = Object.keys(map)
        .filter((key) => !KNOWN.some((k) => k.key === key))
        .map((key) => ({ settingKey: key, label: key, settingValue: map[key] }));
      this.rows.set([...known, ...extras]);
    });
  }

  save(): void {
    this.saving.set(true);
    this.saved.set(false);
    const payload = this.rows().map((r) => ({ settingKey: r.settingKey, settingValue: r.settingValue }));
    this.cms.updateSettings(payload).subscribe({
      next: () => { this.saving.set(false); this.saved.set(true); setTimeout(() => this.saved.set(false), 2500); },
      error: () => this.saving.set(false),
    });
  }
}
