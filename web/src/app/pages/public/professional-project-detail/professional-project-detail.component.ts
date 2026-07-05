import { Component, OnInit, signal } from '@angular/core';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { DecimalPipe } from '@angular/common';
import { Title } from '@angular/platform-browser';

import { IconComponent } from '../../../components/icon/icon.component';
import { SpinnerComponent } from '../../../components/spinner/spinner.component';
import { PublicService } from '../../../core/services/public.service';
import { ProfessionalProject } from '../../../core/models/diosys.model';

@Component({
  selector: 'app-professional-project-detail',
  standalone: true,
  imports: [RouterLink, IconComponent, SpinnerComponent, DecimalPipe],
  templateUrl: './professional-project-detail.component.html',
  styleUrl: './professional-project-detail.component.scss',
})
export class ProfessionalProjectDetailComponent implements OnInit {
  project = signal<ProfessionalProject | null>(null);
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
    this.publicService.getProfessionalProject(projectID).subscribe({
      next: (project) => {
        this.project.set(project);
        this.loading.set(false);
        this.titleService.setTitle(`${project.title} — ${project.company} — Diosys`);
      },
      error: () => {
        this.loading.set(false);
        this.notFound.set(true);
      },
    });
  }
}
