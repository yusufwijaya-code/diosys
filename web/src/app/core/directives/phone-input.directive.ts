import { Directive, HostListener } from '@angular/core';
import { NgControl } from '@angular/forms';

@Directive({
  selector: '[appPhoneInput]',
  standalone: true,
})
export class PhoneInputDirective {
  constructor(private ngControl: NgControl) {}

  @HostListener('keydown', ['$event'])
  onKeyDown(e: KeyboardEvent): void {
    const isNav = ['Backspace', 'Delete', 'ArrowLeft', 'ArrowRight', 'ArrowUp', 'ArrowDown', 'Tab', 'Home', 'End'].includes(e.key);
    if (isNav || e.ctrlKey || e.metaKey) return;
    if (!/^[0-9+]$/.test(e.key)) e.preventDefault();
  }

  @HostListener('paste', ['$event'])
  onPaste(e: ClipboardEvent): void {
    e.preventDefault();
    const raw = e.clipboardData?.getData('text') ?? '';
    const cleaned = raw.replace(/[^0-9+]/g, '');
    const ctrl = this.ngControl.control;
    if (!ctrl) return;
    const el = e.target as HTMLInputElement;
    const val = el.value;
    const start = el.selectionStart ?? val.length;
    const end = el.selectionEnd ?? val.length;
    ctrl.setValue(val.slice(0, start) + cleaned + val.slice(end));
  }
}
