import { Component, OnInit, signal } from '@angular/core';
import { ActivatedRoute, RouterLink } from '@angular/router';

import { IconComponent } from '../../../components/icon/icon.component';
import { PublicService } from '../../../core/services/public.service';
import { DeveloperProfile, Skill } from '../../../core/models/diosys.model';

@Component({
  selector: 'app-developer-profile',
  standalone: true,
  imports: [RouterLink, IconComponent],
  templateUrl: './developer-profile.component.html',
  styleUrl: './developer-profile.component.scss',
})
export class DeveloperProfileComponent implements OnInit {
  profile = signal<DeveloperProfile | null>(null);
  loading = signal(true);
  notFound = signal(false);

  constructor(private route: ActivatedRoute, private publicService: PublicService) {}

  ngOnInit(): void {
    this.route.paramMap.subscribe((params) => {
      const username = params.get('username') ?? '';
      this.load(username);
    });
  }

  private load(username: string): void {
    this.loading.set(true);
    this.notFound.set(false);
    this.publicService.getDeveloperProfile(username).subscribe({
      next: (profile) => {
        this.profile.set(profile);
        this.loading.set(false);
      },
      error: () => {
        this.loading.set(false);
        this.notFound.set(true);
      },
    });
  }

  /** Groups skills by their category for display. */
  skillGroups(skills: Skill[]): { category: string; items: Skill[] }[] {
    const map = new Map<string, Skill[]>();
    for (const skill of skills) {
      const key = skill.category || 'Other';
      if (!map.has(key)) map.set(key, []);
      map.get(key)!.push(skill);
    }
    return Array.from(map, ([category, items]) => ({ category, items }));
  }
}
