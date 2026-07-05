import { Component, OnInit, signal } from '@angular/core';

import { CmsService } from '../../../core/services/cms.service';
import { ClientMessage } from '../../../core/models/diosys.model';

@Component({
  selector: 'app-admin-inbox',
  standalone: true,
  template: `
    <div class="admin-topbar"><h1>Inbox</h1></div>

    <div class="admin-card">
      <table class="admin-table">
        <thead>
          <tr><th>From</th><th>Subject / Message</th><th>Received</th><th>Status</th><th>Actions</th></tr>
        </thead>
        <tbody>
          @for (msg of messages(); track msg.messageID) {
            <tr [class.unread]="!msg.isRead">
              <td>
                <strong>{{ msg.clientName }}</strong><br />
                <a [href]="'mailto:' + msg.clientEmail" class="muted">{{ msg.clientEmail }}</a>
                @if (msg.clientPhone) {
                  <br /><a [href]="waLink(msg.clientPhone)" target="_blank" rel="noopener" class="muted">📱 {{ msg.clientPhone }}</a>
                }
              </td>
              <td>
                @if (msg.subject) { <strong>{{ msg.subject }}</strong><br /> }
                <span class="msg-body">{{ msg.messageBody }}</span>
              </td>
              <td class="muted">{{ msg.createdDate }}</td>
              <td>
                @if (msg.isArchived) { <span class="chip">Archived</span> }
                @else if (msg.isRead) { <span class="chip">Read</span> }
                @else { <span class="chip chip-accent">New</span> }
              </td>
              <td class="actions">
                @if (!msg.isRead) {
                  <button class="btn btn-secondary btn-sm" (click)="markRead(msg)" [disabled]="loadingId() === msg.messageID">
                    @if (loadingId() === msg.messageID) { <span class="btn-spinner"></span> } @else { Mark read }
                  </button>
                }
                <button class="btn btn-secondary btn-sm" (click)="toggleArchive(msg)" [disabled]="loadingId() === msg.messageID">
                  @if (loadingId() === msg.messageID) { <span class="btn-spinner"></span> } @else { {{ msg.isArchived ? 'Unarchive' : 'Archive' }} }
                </button>
                <button class="btn btn-danger btn-sm" (click)="remove(msg)" [disabled]="loadingId() === msg.messageID">
                  @if (loadingId() === msg.messageID) { <span class="btn-spinner"></span> } @else { Delete }
                </button>
              </td>
            </tr>
          }
          @if (messages().length === 0) { <tr><td colspan="5" class="muted">No messages.</td></tr> }
        </tbody>
      </table>
    </div>

    <style>
      .unread td { background: rgba(165, 180, 252, 0.05); }
      .msg-body { display: block; max-width: 420px; color: var(--text-secondary); }
    </style>
  `,
})
export class AdminInboxComponent implements OnInit {
  messages = signal<ClientMessage[]>([]);
  loadingId = signal<number | null>(null);

  constructor(private cms: CmsService) {}

  ngOnInit(): void { this.load(); }

  load(): void {
    this.cms.getMessages().subscribe((m) => this.messages.set(m));
  }

  waLink(phone: string): string {
    return 'https://wa.me/' + phone.replace(/\D/g, '');
  }

  markRead(msg: ClientMessage): void {
    this.loadingId.set(msg.messageID);
    this.cms.updateMessageStatus(msg.messageID, { isRead: true }).subscribe({
      next: () => { this.loadingId.set(null); this.load(); },
      error: () => this.loadingId.set(null),
    });
  }

  toggleArchive(msg: ClientMessage): void {
    this.loadingId.set(msg.messageID);
    this.cms.updateMessageStatus(msg.messageID, { isArchived: msg.isArchived === 0, isRead: true }).subscribe({
      next: () => { this.loadingId.set(null); this.load(); },
      error: () => this.loadingId.set(null),
    });
  }

  remove(msg: ClientMessage): void {
    if (!confirm('Delete this message?')) return;
    this.loadingId.set(msg.messageID);
    this.cms.deleteMessage(msg.messageID).subscribe({
      next: () => { this.loadingId.set(null); this.load(); },
      error: () => this.loadingId.set(null),
    });
  }
}
