import { Component, HostListener, signal } from '@angular/core';
import { RouterLink } from '@angular/router';

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
          <a [routerLink]="['/']" fragment="services" (click)="close()">Services</a>
          <a [routerLink]="['/']" fragment="work" (click)="close()">Work</a>
          <a [routerLink]="['/']" fragment="developers" (click)="close()">Developers</a>
          <a [routerLink]="['/']" fragment="pricing" (click)="close()">Pricing</a>
          <a [routerLink]="['/']" fragment="contact" (click)="close()">Contact</a>
          <a class="btn btn-primary btn-sm cta" [routerLink]="['/']" fragment="contact" (click)="close()">
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
    .nav-links a { color: var(--text-tertiary); font-size: 0.92rem; font-weight: 500; transition: color 0.15s; }
    .nav-links a:hover { color: var(--text-primary); }
    .nav-links a.cta { color: #0b0d12; }

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
        background: #050508;
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
        padding: 0.5rem 0 2rem;
      }
      .nav.menu-open .nav-links {
        display: flex;
        order: 3;
        flex: 0 0 100%;
        border-top: 1px solid var(--border);
      }

      .nav-links a {
        display: block;
        padding: 1rem 0;
        font-size: 1.05rem;
        color: var(--text-tertiary);
        border-bottom: 1px solid var(--border);
      }
      .nav-links a:last-of-type { border-bottom: none; }
      .nav-links a:hover { color: var(--text-primary); }
      .nav-links a.cta {
        display: flex; justify-content: center;
        margin-top: 1.25rem;
        padding: 0.75rem;
        border-bottom: none;
        color: #0b0d12;
      }
    }
  `],
})
export class NavbarComponent {
  scrolled = signal(false);
  menuOpen = signal(false);

  @HostListener('window:scroll')
  onScroll(): void {
    this.scrolled.set(window.scrollY > 20);
    if (this.menuOpen()) this.menuOpen.set(false);
  }

  toggle(): void {
    this.menuOpen.update((v) => !v);
  }

  close(): void {
    this.menuOpen.set(false);
  }
}
