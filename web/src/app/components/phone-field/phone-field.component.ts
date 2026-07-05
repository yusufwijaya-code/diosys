import {
  Component, Input, Output, EventEmitter, OnInit,
  HostListener, ElementRef, signal, computed, ViewChild,
} from '@angular/core';
import { FormsModule } from '@angular/forms';
import { COUNTRIES, Country } from '../../core/data/countries';

@Component({
  selector: 'app-phone-field',
  standalone: true,
  imports: [FormsModule],
  template: `
    <div class="pf-wrap" [class.pf-open]="isOpen()">

      <!-- Trigger button -->
      <button type="button" class="pf-trigger" (click)="toggle()" [attr.aria-expanded]="isOpen()">
        <span class="pf-flag">{{ selected().flag }}</span>
        <span class="pf-code">{{ selected().dialCode }}</span>
        <svg class="pf-chevron" width="12" height="12" viewBox="0 0 12 12" fill="none">
          <path d="M2 4l4 4 4-4" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
        </svg>
      </button>

      <div class="pf-sep"></div>

      <!-- Dropdown panel -->
      @if (isOpen()) {
        <div class="pf-panel">
          <div class="pf-search-row">
            <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" class="pf-search-ic">
              <circle cx="11" cy="11" r="8"/><line x1="21" y1="21" x2="16.65" y2="16.65"/>
            </svg>
            <input
              #searchRef
              class="pf-search"
              type="text"
              [(ngModel)]="query"
              placeholder="Search country…"
              name="countrySearch"
              autocomplete="off"
            />
          </div>
          <ul class="pf-list" role="listbox">
            @for (c of filtered; track c.iso) {
              <li
                class="pf-opt"
                [class.pf-opt--on]="c.iso === selectedIso()"
                (click)="pick(c)"
                role="option"
              >
                <span class="pf-opt-flag">{{ c.flag }}</span>
                <span class="pf-opt-name">{{ c.name }}</span>
                <span class="pf-opt-dial">{{ c.dialCode }}</span>
              </li>
            }
            @empty {
              <li class="pf-empty">No results</li>
            }
          </ul>
        </div>
      }

      <!-- Number input -->
      <input
        class="pf-num"
        type="tel"
        [(ngModel)]="rawNumber"
        (ngModelChange)="emit()"
        (keydown)="onKeyDown($event)"
        (paste)="onPaste($event)"
        [placeholder]="placeholder"
        name="phoneNumber"
        autocomplete="tel-national"
      />
    </div>
  `,
  styles: [`
    :host { display: block; }

    .pf-wrap {
      display: flex;
      align-items: stretch;
      position: relative;
      border: 1px solid var(--border-strong);
      border-radius: var(--radius-sm);
      background: var(--bg-elevated);
      transition: border-color 0.15s, box-shadow 0.15s;
    }
    .pf-wrap:focus-within,
    .pf-open {
      border-color: var(--accent);
      box-shadow: 0 0 0 3px rgba(165, 180, 252, 0.12);
    }

    /* Trigger */
    .pf-trigger {
      display: flex;
      align-items: center;
      gap: 0.3rem;
      padding: 0 0.75rem;
      background: transparent;
      border: none;
      cursor: pointer;
      color: var(--text-primary);
      flex-shrink: 0;
      border-radius: var(--radius-sm) 0 0 var(--radius-sm);
      transition: background 0.12s;
      white-space: nowrap;
      min-width: 80px;
      &:hover { background: rgba(255, 255, 255, 0.04); }
    }
    .pf-flag { font-size: 1.1rem; line-height: 1; }
    .pf-code { font-size: 0.87rem; font-weight: 500; }
    .pf-chevron {
      color: var(--text-secondary);
      margin-left: 0.1rem;
      transition: transform 0.2s;
      .pf-open & { transform: rotate(180deg); }
    }

    /* Divider */
    .pf-sep {
      width: 1px;
      background: var(--border);
      align-self: stretch;
      flex-shrink: 0;
    }

    /* Dropdown panel */
    .pf-panel {
      position: absolute;
      top: calc(100% + 6px);
      left: 0;
      width: 280px;
      background: var(--bg-surface, #1a1d24);
      border: 1px solid var(--border-strong);
      border-radius: var(--radius);
      box-shadow:
        0 12px 40px rgba(0, 0, 0, 0.6),
        0 2px 8px rgba(0, 0, 0, 0.3),
        0 0 0 1px rgba(165, 180, 252, 0.04);
      z-index: 500;
      overflow: hidden;
      animation: pf-in 0.14s ease;
    }
    @keyframes pf-in {
      from { opacity: 0; transform: translateY(-5px) scale(0.98); }
      to   { opacity: 1; transform: translateY(0)  scale(1); }
    }

    /* Search row */
    .pf-search-row {
      display: flex;
      align-items: center;
      gap: 0.5rem;
      padding: 0.6rem 0.75rem;
      border-bottom: 1px solid var(--border);
      background: rgba(255, 255, 255, 0.02);
    }
    .pf-search-ic { color: var(--text-secondary); flex-shrink: 0; }
    .pf-search {
      flex: 1;
      border: none;
      outline: none;
      background: transparent;
      color: var(--text-primary);
      font-size: 0.875rem;
      font-family: inherit;
      &::placeholder { color: var(--text-tertiary); }
    }

    /* List */
    .pf-list {
      max-height: 236px;
      overflow-y: auto;
      padding: 0.25rem 0;
      scrollbar-width: thin;
      scrollbar-color: var(--border-strong) transparent;
      &::-webkit-scrollbar { width: 4px; }
      &::-webkit-scrollbar-track { background: transparent; }
      &::-webkit-scrollbar-thumb {
        background: var(--border-strong);
        border-radius: 4px;
      }
    }

    .pf-opt {
      display: flex;
      align-items: center;
      gap: 0.65rem;
      padding: 0.5rem 0.875rem;
      cursor: pointer;
      transition: background 0.1s;
      &:hover { background: rgba(165, 180, 252, 0.07); }
    }
    .pf-opt--on {
      background: rgba(165, 180, 252, 0.1);
      .pf-opt-name { color: var(--accent); }
      .pf-opt-dial { color: var(--accent); opacity: 0.75; }
    }
    .pf-opt-flag { font-size: 1.1rem; flex-shrink: 0; line-height: 1; }
    .pf-opt-name {
      flex: 1;
      font-size: 0.85rem;
      color: var(--text-primary);
      white-space: nowrap;
      overflow: hidden;
      text-overflow: ellipsis;
    }
    .pf-opt-dial {
      font-size: 0.78rem;
      color: var(--text-secondary);
      font-variant-numeric: tabular-nums;
      flex-shrink: 0;
    }
    .pf-empty {
      padding: 1rem;
      text-align: center;
      font-size: 0.85rem;
      color: var(--text-secondary);
    }

    /* Number input */
    .pf-num {
      flex: 1;
      border: none;
      outline: none;
      background: transparent;
      color: var(--text-primary);
      font-size: 0.92rem;
      font-family: inherit;
      padding: 0.65rem 0.875rem;
      min-width: 0;
      &::placeholder { color: var(--text-tertiary); }
    }
  `],
})
export class PhoneFieldComponent implements OnInit {
  @Input() value = '';
  @Output() valueChange = new EventEmitter<string>();
  @ViewChild('searchRef') searchRef?: ElementRef<HTMLInputElement>;

  countries = COUNTRIES;
  selectedIso = signal('ID');
  rawNumber = '';
  isOpen = signal(false);
  query = '';

  selected = computed(() =>
    this.countries.find(c => c.iso === this.selectedIso()) ?? this.countries[0]
  );

  get filtered(): Country[] {
    const q = this.query.toLowerCase().trim();
    if (!q) return this.countries;
    return this.countries.filter(c =>
      c.name.toLowerCase().includes(q) ||
      c.dialCode.includes(q) ||
      c.iso.toLowerCase().includes(q)
    );
  }

  get placeholder(): string {
    return this.selectedIso() === 'ID' ? '81234567890' : '1234567890';
  }

  constructor(private el: ElementRef) {}

  ngOnInit(): void {
    if (this.value) {
      const match = this.countries.find(c => this.value.startsWith(c.dialCode));
      if (match) {
        this.selectedIso.set(match.iso);
        this.rawNumber = this.value.slice(match.dialCode.length).trimStart();
      } else {
        this.rawNumber = this.value;
      }
    }
  }

  toggle(): void {
    this.isOpen.update(v => !v);
    if (this.isOpen()) {
      this.query = '';
      setTimeout(() => this.searchRef?.nativeElement.focus(), 60);
    }
  }

  pick(c: Country): void {
    this.selectedIso.set(c.iso);
    this.isOpen.set(false);
    this.emit();
  }

  @HostListener('document:click', ['$event'])
  onDocClick(e: MouseEvent): void {
    if (!this.el.nativeElement.contains(e.target as Node)) {
      this.isOpen.set(false);
    }
  }

  emit(): void {
    const num = this.rawNumber.trim();
    this.valueChange.emit(num ? this.selected().dialCode + num : '');
  }

  onKeyDown(e: KeyboardEvent): void {
    const nav = ['Backspace', 'Delete', 'ArrowLeft', 'ArrowRight', 'ArrowUp', 'ArrowDown', 'Tab', 'Home', 'End'].includes(e.key);
    if (nav || e.ctrlKey || e.metaKey) return;
    if (!/^[0-9]$/.test(e.key)) e.preventDefault();
  }

  onPaste(e: ClipboardEvent): void {
    e.preventDefault();
    const raw = e.clipboardData?.getData('text') ?? '';
    const cleaned = raw.replace(/[^0-9]/g, '');
    const target = e.target as HTMLInputElement;
    const val = target.value;
    const s = target.selectionStart ?? val.length;
    const end = target.selectionEnd ?? val.length;
    this.rawNumber = val.slice(0, s) + cleaned + val.slice(end);
    this.emit();
  }
}
