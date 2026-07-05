import { Component, OnInit, signal } from '@angular/core';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { DecimalPipe } from '@angular/common';
import { Title } from '@angular/platform-browser';

import { IconComponent } from '../../../components/icon/icon.component';
import { SpinnerComponent } from '../../../components/spinner/spinner.component';
import { PublicService } from '../../../core/services/public.service';
import { NavigationHistoryService } from '../../../core/services/navigation-history.service';
import { Project } from '../../../core/models/diosys.model';

@Component({
  selector: 'app-project-detail',
  standalone: true,
  imports: [RouterLink, IconComponent, SpinnerComponent, DecimalPipe],
  templateUrl: './project-detail.component.html',
  styleUrl: './project-detail.component.scss',
})
export class ProjectDetailComponent implements OnInit {
  project = signal<Project | null>(null);
  username = '';
  loading = signal(true);
  notFound = signal(false);

  backRoute: string[] = [];
  backFragment = 'projects';

  constructor(
    private route: ActivatedRoute,
    private publicService: PublicService,
    private titleService: Title,
    private navHistory: NavigationHistoryService,
  ) {}

  ngOnInit(): void {
    this.username = this.route.snapshot.paramMap.get('username') ?? '';
    const projectID = Number(this.route.snapshot.paramMap.get('projectID'));

    const prev = this.navHistory.previousUrl;
    // If came from home page (prev is '/', '/#something', or empty with no history)
    if (!prev || prev === '/' || prev.startsWith('/#')) {
      this.backRoute = ['/'];
      this.backFragment = 'work';
    } else {
      // Came from developer profile
      this.backRoute = ['/', this.username];
      this.backFragment = 'projects';
    }

    this.load(projectID);
  }

  private load(projectID: number): void {
    this.publicService.getProject(projectID).subscribe({
      next: (project) => {
        this.project.set(project);
        this.loading.set(false);
        this.titleService.setTitle(`${project.title} — Diosys`);
      },
      error: () => {
        this.loading.set(false);
        this.notFound.set(true);
      },
    });
  }
}
