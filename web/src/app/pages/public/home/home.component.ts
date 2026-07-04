import { Component, OnInit, computed, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { RouterLink } from '@angular/router';

import { IconComponent } from '../../../components/icon/icon.component';
import { PublicService } from '../../../core/services/public.service';
import {
  DeveloperCard,
  MessageRequest,
  PricePlan,
  Project,
  Service,
  SettingsMap,
} from '../../../core/models/diosys.model';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [FormsModule, RouterLink, IconComponent],
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss',
})
export class HomeComponent implements OnInit {
  services = signal<Service[]>([]);
  projects = signal<Project[]>([]);
  developers = signal<DeveloperCard[]>([]);
  pricingPlans = signal<PricePlan[]>([]);
  settings = signal<SettingsMap>({});

  sending = signal(false);
  sent = signal(false);
  error = signal('');

  form: MessageRequest = { clientName: '', clientEmail: '', clientPhone: '', subject: '', messageBody: '' };

  constructor(private publicService: PublicService) {}

  whatsappLink = computed(() => {
    const s = this.settings();
    const number = (s['whatsappNumber'] || '').replace(/[^0-9]/g, '');
    if (!number) return '';
    const text = encodeURIComponent(s['whatsappDefaultMessage'] || '');
    return `https://wa.me/${number}${text ? '?text=' + text : ''}`;
  });

  ngOnInit(): void {
    this.publicService.getServices().subscribe((s) => this.services.set(s));
    this.publicService.getProjects().subscribe((p) => this.projects.set(p.slice(0, 6)));
    this.publicService.getDevelopers().subscribe((d) => this.developers.set(d));
    this.publicService.getPricing().subscribe((p) => this.pricingPlans.set(p));
    this.publicService.getSettings().subscribe((s) => this.settings.set(s));
  }

  formatPrice(price: number, currency: string): string {
    if (currency === 'IDR') {
      return 'Rp ' + price.toLocaleString('id-ID');
    }
    return currency + ' ' + price.toLocaleString();
  }

  submit(): void {
    if (!this.form.clientName || !this.form.clientEmail || !this.form.messageBody) {
      this.error.set('Please fill in your name, email, and message.');
      return;
    }
    this.error.set('');
    this.sending.set(true);
    this.publicService.sendContact(this.form).subscribe({
      next: () => {
        this.sending.set(false);
        this.sent.set(true);
        this.form = { clientName: '', clientEmail: '', clientPhone: '', subject: '', messageBody: '' };
      },
      error: () => {
        this.sending.set(false);
        this.error.set('Something went wrong. Please try again.');
      },
    });
  }
}
