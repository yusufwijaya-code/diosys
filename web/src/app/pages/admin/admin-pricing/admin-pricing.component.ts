import { Component, OnInit, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';

import { CmsService } from '../../../core/services/cms.service';
import { PricePlan, PricePlanRequest } from '../../../core/models/diosys.model';

interface FeatureRow {
  text: string;
  isIncluded: boolean;
}

const emptyForm = (): PricePlanRequest => ({
  title: '', subtitle: '', price: 0, currency: 'IDR', billingPeriod: '/ project',
  badge: '', originalPrice: null, discountPercent: null,
  isFeatured: false, flagActive: true, orderNo: 0, features: [],
});

@Component({
  selector: 'app-admin-pricing',
  standalone: true,
  imports: [FormsModule],
  templateUrl: './admin-pricing.component.html',
})
export class AdminPricingComponent implements OnInit {
  plans = signal<PricePlan[]>([]);
  showForm = signal(false);
  editingId = signal<number | null>(null);
  saving = signal(false);
  error = signal('');

  form: PricePlanRequest = emptyForm();
  features: FeatureRow[] = [];

  // Whether to show the original price / promo section
  hasOriginalPrice = false;

  constructor(private cms: CmsService) {}

  ngOnInit(): void {
    this.load();
  }

  load(): void {
    this.cms.getPricingPlans().subscribe((p) => this.plans.set(p));
  }

  /** Auto-calculates discount % from original and current price. */
  get computedDiscount(): number | null {
    const orig = this.form.originalPrice;
    const price = this.form.price;
    if (!this.hasOriginalPrice || !orig || orig <= price) return null;
    return Math.round(((orig - price) / orig) * 100);
  }

  /** Preview text shown next to the original price field. */
  get discountPreview(): string {
    const d = this.computedDiscount;
    if (!d) return '';
    return `→ Save ${d}% off`;
  }

  startCreate(): void {
    this.form = emptyForm();
    this.features = [];
    this.hasOriginalPrice = false;
    this.editingId.set(null);
    this.showForm.set(true);
    this.error.set('');
  }

  startEdit(plan: PricePlan): void {
    this.hasOriginalPrice = plan.originalPrice !== null;
    this.form = {
      title: plan.title, subtitle: plan.subtitle, price: plan.price,
      currency: plan.currency, billingPeriod: plan.billingPeriod, badge: plan.badge,
      originalPrice: plan.originalPrice, discountPercent: plan.discountPercent,
      isFeatured: plan.isFeatured, flagActive: plan.flagActive,
      orderNo: plan.orderNo, features: [],
    };
    this.features = plan.features.map((f) => ({ text: f.text, isIncluded: f.isIncluded }));
    this.editingId.set(plan.planID);
    this.showForm.set(true);
    this.error.set('');
  }

  cancel(): void {
    this.showForm.set(false);
  }

  addFeature(): void {
    this.features.push({ text: '', isIncluded: true });
  }

  removeFeature(i: number): void {
    this.features.splice(i, 1);
  }

  save(): void {
    this.saving.set(true);
    this.error.set('');

    // Discount % always derived from prices — never manually entered.
    const discount = this.computedDiscount;
    const payload: PricePlanRequest = {
      ...this.form,
      originalPrice: this.hasOriginalPrice ? this.form.originalPrice : null,
      discountPercent: this.hasOriginalPrice ? discount : null,
      features: this.features.filter((f) => f.text.trim()),
    };

    const id = this.editingId();
    const req = id ? this.cms.updatePricingPlan(id, payload) : this.cms.createPricingPlan(payload);

    req.subscribe({
      next: () => {
        this.saving.set(false);
        this.showForm.set(false);
        this.load();
      },
      error: (err) => {
        this.saving.set(false);
        this.error.set(err?.error?.message || 'Failed to save plan.');
      },
    });
  }

  remove(plan: PricePlan): void {
    if (!confirm(`Delete plan "${plan.title}"?`)) return;
    this.cms.deletePricingPlan(plan.planID).subscribe(() => this.load());
  }

  formatPrice(price: number, currency: string): string {
    if (currency === 'IDR') return 'Rp ' + price.toLocaleString('id-ID');
    return currency + ' ' + price.toLocaleString();
  }
}
