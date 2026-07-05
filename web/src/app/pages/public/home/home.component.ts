import { Component, OnInit, OnDestroy, computed, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { ActivatedRoute, NavigationEnd, Router, RouterLink } from '@angular/router';
import { ViewportScroller } from '@angular/common';
import { Subscription } from 'rxjs';
import { filter } from 'rxjs/operators';

import { IconComponent } from '../../../components/icon/icon.component';
import { PhoneInputDirective } from '../../../core/directives/phone-input.directive';
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
  imports: [FormsModule, RouterLink, IconComponent, PhoneInputDirective],
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss',
})
export class HomeComponent implements OnInit, OnDestroy {
  services = signal<Service[]>([]);
  projects = signal<Project[]>([]);
  developers = signal<DeveloperCard[]>([]);
  pricingPlans = signal<PricePlan[]>([]);
  settings = signal<SettingsMap>({});

  sending = signal(false);
  sent = signal(false);
  error = signal('');

  form: MessageRequest = { clientName: '', clientEmail: '', clientPhone: '', subject: '', messageBody: '' };

  private apisLoaded = 0;
  private pendingFragment: string | null = null;
  private routerSub!: Subscription;

  constructor(
    private publicService: PublicService,
    private route: ActivatedRoute,
    private router: Router,
    private scroller: ViewportScroller,
  ) {}

  whatsappLink = computed(() => {
    const s = this.settings();
    const number = (s['whatsappNumber'] || '').replace(/[^0-9]/g, '');
    if (!number) return '';
    const text = encodeURIComponent(s['whatsappDefaultMessage'] || '');
    return `https://wa.me/${number}${text ? '?text=' + text : ''}`;
  });

  private onApiLoaded(): void {
    if (++this.apisLoaded < 5) return;
    if (this.pendingFragment) {
      const fragment = this.pendingFragment;
      this.pendingFragment = null;
      setTimeout(() => this.scroller.scrollToAnchor(fragment), 0);
    }
  }

  ngOnInit(): void {
    // Capture fragment from initial navigation (cross-page arrival)
    this.pendingFragment = this.route.snapshot.fragment;

    this.publicService.getServices().subscribe((s) => { this.services.set(s); this.onApiLoaded(); });
    this.publicService.getProjects().subscribe((p) => { this.projects.set(p.slice(0, 6)); this.onApiLoaded(); });
    this.publicService.getDevelopers().subscribe((d) => { this.developers.set(d); this.onApiLoaded(); });
    this.publicService.getPricing().subscribe((p) => { this.pricingPlans.set(p); this.onApiLoaded(); });
    this.publicService.getSettings().subscribe((s) => { this.settings.set(s); this.onApiLoaded(); });

    // Handle fragment navigation while already on the home page (anchorScrolling is disabled globally)
    this.routerSub = this.router.events.pipe(
      filter(e => e instanceof NavigationEnd)
    ).subscribe(() => {
      const fragment = this.route.snapshot.fragment;
      if (fragment) setTimeout(() => this.scroller.scrollToAnchor(fragment), 0);
    });
  }

  ngOnDestroy(): void {
    this.routerSub?.unsubscribe();
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
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(this.form.clientEmail)) {
      this.error.set('Please enter a valid email address.');
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
