import { Component, OnInit, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { ActivatedRoute, RouterLink } from '@angular/router';

import { CmsService } from '../../../core/services/cms.service';
import {
  Certificate, Developer, Education, Experience, Skill, Summary,
} from '../../../core/models/diosys.model';

@Component({
  selector: 'app-admin-portfolio',
  standalone: true,
  imports: [FormsModule, RouterLink],
  templateUrl: './admin-portfolio.component.html',
})
export class AdminPortfolioComponent implements OnInit {
  userID = 0;
  developer = signal<Developer | null>(null);
  notice = signal('');

  summary: Summary = { summaryID: 0, userID: 0, content: '', stats: [], facts: [] };

  experiences = signal<Experience[]>([]);
  educations = signal<Education[]>([]);
  certificates = signal<Certificate[]>([]);
  skills = signal<Skill[]>([]);

  expForm = { position: '', company: '', period: '', orderNo: 0, technologies: '', responsibilities: '' };
  eduForm = { degree: '', institution: '', year: '', type: '', orderNo: 0, achievements: '' };
  certForm = { name: '', issuer: '', period: '', link: '', orderNo: 0 };
  skillForm = { name: '', level: '', category: '', orderNo: 0 };

  constructor(private route: ActivatedRoute, private cms: CmsService) {}

  ngOnInit(): void {
    this.userID = Number(this.route.snapshot.paramMap.get('userID'));
    this.cms.getDeveloper(this.userID).subscribe((d) => this.developer.set(d));
    this.cms.getSummary(this.userID).subscribe((s) => (this.summary = s));
    this.reloadLists();
  }

  reloadLists(): void {
    this.cms.getExperiences(this.userID).subscribe((x) => this.experiences.set(x));
    this.cms.getEducations(this.userID).subscribe((x) => this.educations.set(x));
    this.cms.getCertificates(this.userID).subscribe((x) => this.certificates.set(x));
    this.cms.getSkills(this.userID).subscribe((x) => this.skills.set(x));
  }

  private flash(message: string): void {
    this.notice.set(message);
    setTimeout(() => this.notice.set(''), 2500);
  }

  private lines(value: string): string[] {
    return value.split('\n').map((v) => v.trim()).filter(Boolean);
  }
  private csv(value: string): string[] {
    return value.split(',').map((v) => v.trim()).filter(Boolean);
  }

  addStat(): void { this.summary.stats.push({ number: '', label: '' }); }
  removeStat(i: number): void { this.summary.stats.splice(i, 1); }
  addFact(): void { this.summary.facts.push({ icon: 'check', text: '' }); }
  removeFact(i: number): void { this.summary.facts.splice(i, 1); }
  saveSummary(): void {
    this.cms.saveSummary(this.userID, {
      content: this.summary.content,
      stats: this.summary.stats,
      facts: this.summary.facts,
    }).subscribe((s) => { this.summary = s; this.flash('Summary saved.'); });
  }

  addExperience(): void {
    if (!this.expForm.position || !this.expForm.company) return;
    this.cms.createExperience(this.userID, {
      position: this.expForm.position, company: this.expForm.company, period: this.expForm.period,
      orderNo: this.expForm.orderNo,
      technologies: this.csv(this.expForm.technologies),
      responsibilities: this.lines(this.expForm.responsibilities),
    }).subscribe(() => {
      this.expForm = { position: '', company: '', period: '', orderNo: 0, technologies: '', responsibilities: '' };
      this.reloadLists();
    });
  }
  deleteExperience(id: number): void {
    this.cms.deleteExperience(this.userID, id).subscribe(() => this.reloadLists());
  }

  addEducation(): void {
    if (!this.eduForm.degree || !this.eduForm.institution) return;
    this.cms.createEducation(this.userID, {
      degree: this.eduForm.degree, institution: this.eduForm.institution, year: this.eduForm.year,
      type: this.eduForm.type, orderNo: this.eduForm.orderNo,
      achievements: this.lines(this.eduForm.achievements),
    }).subscribe(() => {
      this.eduForm = { degree: '', institution: '', year: '', type: '', orderNo: 0, achievements: '' };
      this.reloadLists();
    });
  }
  deleteEducation(id: number): void {
    this.cms.deleteEducation(this.userID, id).subscribe(() => this.reloadLists());
  }

  addCertificate(): void {
    if (!this.certForm.name || !this.certForm.issuer) return;
    this.cms.createCertificate(this.userID, { ...this.certForm }).subscribe(() => {
      this.certForm = { name: '', issuer: '', period: '', link: '', orderNo: 0 };
      this.reloadLists();
    });
  }
  deleteCertificate(id: number): void {
    this.cms.deleteCertificate(this.userID, id).subscribe(() => this.reloadLists());
  }

  addSkill(): void {
    if (!this.skillForm.name) return;
    this.cms.createSkill(this.userID, { ...this.skillForm }).subscribe(() => {
      this.skillForm = { name: '', level: '', category: '', orderNo: 0 };
      this.reloadLists();
    });
  }
  deleteSkill(id: number): void {
    this.cms.deleteSkill(this.userID, id).subscribe(() => this.reloadLists());
  }
}
