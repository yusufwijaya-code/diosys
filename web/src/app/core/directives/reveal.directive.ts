import { Directive, ElementRef, Input, OnDestroy, OnInit } from '@angular/core';

@Directive({ selector: '[appReveal]', standalone: true })
export class RevealDirective implements OnInit, OnDestroy {
  @Input('appReveal') delay: number = 0;

  private observer: IntersectionObserver | null = null;

  constructor(private el: ElementRef<HTMLElement>) {}

  ngOnInit(): void {
    const el = this.el.nativeElement;
    el.style.setProperty('--rv-delay', `${this.delay}ms`);
    el.classList.add('rv');

    this.observer = new IntersectionObserver(
      ([entry]) => {
        if (entry.isIntersecting) {
          el.classList.add('rv-on');
          this.observer?.disconnect();
        }
      },
      { threshold: 0.08, rootMargin: '0px 0px -32px 0px' }
    );
    this.observer.observe(el);
  }

  ngOnDestroy(): void {
    this.observer?.disconnect();
  }
}
