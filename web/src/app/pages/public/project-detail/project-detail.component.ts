import { Component, OnInit, signal } from '@angular/core';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { DecimalPipe } from '@angular/common';
import { Title } from '@angular/platform-browser';

import { IconComponent } from '../../../components/icon/icon.component';
import { PublicService } from '../../../core/services/public.service';
import { Project } from '../../../core/models/diosys.model';

@Component({
  selector: 'app-project-detail',
  standalone: true,
  imports: [RouterLink, IconComponent, DecimalPipe],
  templateUrl: './project-detail.component.html',
  styleUrl: './project-detail.component.scss',
})
export class ProjectDetailComponent implements OnInit {
  project = signal<Project | null>(null);
  username = '';
  loading = signal(true);
  notFound = signal(false);

  constructor(
    private route: ActivatedRoute,
    private publicService: PublicService,
    private titleService: Title,
  ) {}

  ngOnInit(): void {
    this.username = this.route.snapshot.paramMap.get('username') ?? '';
    const projectID = Number(this.route.snapshot.paramMap.get('projectID'));
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
