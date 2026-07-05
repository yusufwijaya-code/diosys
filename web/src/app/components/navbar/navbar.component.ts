import { Component, HostListener, OnDestroy, signal } from '@angular/core';
import { NavigationEnd, Router, RouterLink } from '@angular/router';
import { Subscription } from 'rxjs';
import { filter } from 'rxjs/operators';

@Component({
  selector: 'app-navbar',
  standalone: true,
  imports: [RouterLink],
  template: `
    <header class="nav" [class.scrolled]="scrolled()" [class.menu-open]="menuOpen()">
      <div class="container nav-inner">
        <a class="brand" [routerLink]="['/']" (click)="close()">
          <img src="assets/diosys-logo-mark.svg" alt="Diosys" class="brand-logo" />
        </a>

        <nav class="nav-links">
          <a href="/#services" [class.active]="activeSection() === 'services'" (click)="navClick('services', $event)">Services</a>
          <a href="/#work" [class.active]="activeSection() === 'work'" (click)="navClick('work', $event)">Work</a>
          <a href="/#developers" [class.active]="activeSection() === 'developers'" (click)="navClick('developers', $event)">Developers</a>
          <a href="/#pricing" [class.active]="activeSection() === 'pricing'" (click)="navClick('pricing', $event)">Pricing</a>
          <a href="/#contact" [class.active]="activeSection() === 'contact'" (click)="navClick('contact', $event)">Contact</a>
          <a class="btn btn-primary btn-sm cta" href="/#contact" (click)="navClick('contact', $event)">
            Start a project
          </a>
        </nav>

        <button class="burger" [class.active]="menuOpen()" (click)="toggle()" aria-label="Menu">
          <span></span><span></span><span></span>
        </button>
      </div>
    </header>
  `,
  styles: [`
    /* ── Base ── */
    .nav {
      position: fixed; top: 0; left: 0; right: 0;
      z-index: 100;
      transition: background 0.25s ease, border-color 0.25s ease;
      border-bottom: 1px solid transparent;
    }
    .nav.scrolled {
      background: rgba(5, 5, 8, 0.92);
      backdrop-filter: blur(16px);
      -webkit-backdrop-filter: blur(16px);
      border-bottom-color: var(--border);
    }
    .nav-inner {
      display: flex; align-items: center;
      justify-content: space-between; height: 76px;
    }
    .brand-logo { height: 56px; width: auto; display: block; }

    /* Desktop links */
    .nav-links { display: flex; align-items: center; gap: 1.75rem; }
    .nav-links a {
      color: var(--text-tertiary); font-size: 0.92rem; font-weight: 500;
      position: relative;
    }
    .nav-links a:hover { color: var(--text-primary); }
    .nav-links a.active {
      color: var(--text-primary);
      &::after {
        content: '';
        position: absolute;
        bottom: -4px;
        left: 0; right: 0;
        height: 2px;
        background: var(--accent);
        border-radius: 1px;
      }
    }
    .nav-links a.cta { color: #0b0d12; }
    .nav-links a.cta::after { display: none; }

    /* Burger */
    .burger {
      display: none; flex-direction: column; justify-content: center;
      gap: 5px; background: none; border: 0; cursor: pointer;
      padding: 6px; width: 36px; height: 36px; flex-shrink: 0;
    }
    .burger span {
      display: block; width: 22px; height: 2px;
      background: var(--text-primary); border-radius: 2px;
      transition: transform 0.25s ease, opacity 0.2s ease;
    }
    .burger.active span:nth-child(1) { transform: translateY(7px) rotate(45deg); }
    .burger.active span:nth-child(2) { opacity: 0; transform: scaleX(0); }
    .burger.active span:nth-child(3) { transform: translateY(-7px) rotate(-45deg); }

    /* ── Mobile ── */
    @media (max-width: 820px) {
      .burger { display: flex; }

      /* Header expands to full screen when open */
      .nav.menu-open {
        background: rgba(5, 5, 8, 0.97);
        backdrop-filter: blur(24px);
        -webkit-backdrop-filter: blur(24px);
        bottom: 0;
        overflow-y: auto;
        border-bottom: none;
      }

      /* When open: allow nav-links to wrap below the 76px bar */
      .nav.menu-open .nav-inner {
        flex-wrap: wrap;
        height: auto;
        align-content: flex-start;
      }

      /* Brand and burger share the top 76px row */
      .nav.menu-open .brand {
        height: 76px;
        display: flex;
        align-items: center;
      }
      .nav.menu-open .burger {
        order: 2;
        height: 76px;
        display: flex;
        align-items: center;
      }

      /* Links: hidden by default, shown when parent has menu-open */
      .nav-links {
        display: none;
        flex-direction: column;
        align-items: stretch;
        gap: 0;
        padding: 0 0 3rem;
      }
      .nav.menu-open .nav-links {
        display: flex;
        order: 3;
        flex: 0 0 100%;
        border-top: 1px solid var(--border);
        animation: mobileMenuIn 0.28s cubic-bezier(0.25, 0.46, 0.45, 0.94) both;
      }

      @keyframes mobileMenuIn {
        from { opacity: 0; transform: translateY(-8px); }
        to   { opacity: 1; transform: none; }
      }

      .nav-links a {
        display: flex;
        align-items: center;
        padding: 1.1rem 0;
        font-size: 1.05rem;
        font-weight: 500;
        color: var(--text-tertiary);
        border-bottom: 1px solid rgba(255,255,255,0.05);
        transition: color 0.15s;
      }
      .nav-links a::after { display: none !important; } /* no desktop underline on mobile */
      .nav-links a:last-of-type { border-bottom: none; }
      .nav-links a:hover,
      .nav-links a:active { color: var(--text-primary); }
      .nav-links a.active {
        color: var(--accent);
      }
      .nav-links a.cta {
        display: flex; justify-content: center;
        margin-top: 1.5rem;
        padding: 0.85rem;
        border: 1px solid transparent;
        border-radius: 10px;
        color: #0b0d12;
        min-height: 52px;
      }
    }
  `],
})
export class NavbarComponent implements OnDestroy {
  scrolled = signal(false);
  menuOpen = signal(false);
  activeSection = signal('');

  private isHomePage = false;
  private routerSub: Subscription;
  private scrollTimer: ReturnType<typeof setTimeout> | null = null;
  private suppressScroll = false;

  private static readonly SECTION_IDS = ['services', 'work', 'developers', 'pricing', 'contact'];

  constructor(private router: Router) {
    this.routerSub = router.events.pipe(
      filter(e => e instanceof NavigationEnd)
    ).subscribe((e: NavigationEnd) => {
      const url = (e as NavigationEnd).urlAfterRedirects;
      if (url === '/' || url.startsWith('/#')) {
        this.isHomePage = true;
        // Don't override if a nav click already set the section
        if (!this.suppressScroll) this.updateActiveSection();
      } else if (!url.startsWith('/admin')) {
        this.isHomePage = false;
        if (!this.suppressScroll) this.activeSection.set('developers');
      } else {
        this.isHomePage = false;
        this.activeSection.set('');
      }
    });
  }

  @HostListener('document:keydown.escape')
  onEscapeKey(): void {
    if (this.menuOpen()) this.close();
  }

  @HostListener('window:scroll')
  onScroll(): void {
    this.scrolled.set(window.scrollY > 20);
    if (this.menuOpen()) this.menuOpen.set(false);
    if (this.isHomePage && !this.suppressScroll) {
      this.updateActiveSection();
    }
  }

  navClick(section: string, event: MouseEvent): void {
    // Allow ctrl/cmd+click and middle-click to open in new tab normally
    if (event.ctrlKey || event.metaKey || event.button === 1) return;
    event.preventDefault();

    this.menuOpen.set(false);
    this.activeSection.set(section);
    this.suppressScroll = true;
    if (this.scrollTimer) clearTimeout(this.scrollTimer);

    if (this.isHomePage) {
      const el = document.getElementById(section);
      if (el) {
        const top = el.getBoundingClientRect().top + window.scrollY - 80;
        window.scrollTo({ top, behavior: 'smooth' });
      }
      this.scrollTimer = setTimeout(() => {
        this.suppressScroll = false;
        this.updateActiveSection();
      }, 900);
    } else {
      this.router.navigate(['/'], { fragment: section });
      this.scrollTimer = setTimeout(() => { this.suppressScroll = false; }, 900);
    }
  }

  private updateActiveSection(): void {
    const threshold = 100;
    let active = '';
    for (const id of [...NavbarComponent.SECTION_IDS].reverse()) {
      const el = document.getElementById(id);
      if (!el) continue;
      if (el.getBoundingClientRect().top <= threshold) {
        active = id;
        break;
      }
    }
    this.activeSection.set(active);
  }

  ngOnDestroy(): void {
    this.routerSub.unsubscribe();
    if (this.scrollTimer) clearTimeout(this.scrollTimer);
    document.body.style.overflow = '';
  }

  toggle(): void {
    const next = !this.menuOpen();
    this.menuOpen.set(next);
    document.body.style.overflow = next ? 'hidden' : '';
  }

  close(): void {
    this.menuOpen.set(false);
    document.body.style.overflow = '';
  }
}
