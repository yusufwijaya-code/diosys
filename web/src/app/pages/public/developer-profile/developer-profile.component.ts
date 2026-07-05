import { Component, OnInit, OnDestroy, signal } from '@angular/core';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { SlicePipe } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Title } from '@angular/platform-browser';

import { IconComponent } from '../../../components/icon/icon.component';
import { SpinnerComponent } from '../../../components/spinner/spinner.component';
import { PhoneFieldComponent } from '../../../components/phone-field/phone-field.component';
import { PublicService } from '../../../core/services/public.service';
import { DeveloperProfile, MessageRequest, Skill } from '../../../core/models/diosys.model';

@Component({
  selector: 'app-developer-profile',
  standalone: true,
  imports: [RouterLink, IconComponent, SpinnerComponent, PhoneFieldComponent, SlicePipe, FormsModule],
  templateUrl: './developer-profile.component.html',
  styleUrl: './developer-profile.component.scss',
})
export class DeveloperProfileComponent implements OnInit, OnDestroy {
  profile = signal<DeveloperProfile | null>(null);
  loading = signal(true);
  notFound = signal(false);

  sending = signal(false);
  sent = signal(false);
  contactError = signal('');
  contactForm: MessageRequest = { clientName: '', clientEmail: '', clientPhone: '', subject: '', messageBody: '' };

  constructor(
    private route: ActivatedRoute,
    private publicService: PublicService,
    private titleService: Title,
  ) {}

  ngOnInit(): void {
    this.route.paramMap.subscribe((params) => {
      const username = params.get('username') ?? '';
      this.load(username);
    });
  }

  private load(username: string): void {
    this.loading.set(true);
    this.notFound.set(false);
    const fragment = this.route.snapshot.fragment;
    this.publicService.getDeveloperProfile(username).subscribe({
      next: (profile) => {
        this.profile.set(profile);
        this.loading.set(false);
        this.titleService.setTitle(`${profile.developer.fullName} — Diosys`);
        if (fragment) setTimeout(() => this.scrollTo(fragment), 100);
      },
      error: () => {
        this.loading.set(false);
        this.notFound.set(true);
      },
    });
  }

  private readonly skillIcons: Record<string, string> = {
    'angular': 'angular', 'react': 'react', 'react native': 'react',
    'next.js': 'nextdotjs', 'typescript': 'typescript', 'javascript': 'javascript',
    'scss / css': 'sass', 'sass': 'sass', 'css': 'css3',
    'tailwind css': 'tailwindcss', 'tailwind': 'tailwindcss',
    'go (golang)': 'go', 'go': 'go', 'golang': 'go',
    'node.js': 'nodedotjs', 'nodejs': 'nodedotjs',
    'php / laravel': 'laravel', 'php': 'php', 'laravel': 'laravel',
    'graphql': 'graphql', 'rest api': 'postman',
    'mysql': 'mysql', 'postgresql': 'postgresql', 'mongodb': 'mongodb',
    'redis': 'redis', 'sqlite': 'sqlite',
    'flutter': 'flutter', 'dart': 'dart',
    'docker': 'docker', 'kubernetes': 'kubernetes',
    'git / github': 'github', 'git': 'git', 'github': 'github',
    'google cloud': 'googlecloud', 'aws': 'amazonaws', 'azure': 'microsoftazure',
    'figma': 'figma', 'linux': 'linux',
    'openai api': 'openai', 'openai': 'openai',
    'langchain': 'langchain', 'python': 'python',
    'claude api': 'anthropic', 'anthropic': 'anthropic',
    'vue.js': 'vuedotjs', 'vue': 'vuedotjs',
    'svelte': 'svelte', 'astro': 'astro', 'remix': 'remix', 'nuxt': 'nuxt',
    'nestjs': 'nestjs', 'express.js': 'express', 'express': 'express',
    'django': 'django', 'fastapi': 'fastapi', 'flask': 'flask',
    'firebase': 'firebase', 'supabase': 'supabase', 'stripe': 'stripe',
    'vite': 'vite', 'webpack': 'webpack', 'jest': 'jest',
    'bootstrap': 'bootstrap', 'redux': 'redux',
    'vercel': 'vercel', 'netlify': 'netlify', 'nginx': 'nginx',
    'swift': 'swift', 'kotlin': 'kotlin', 'java': 'java', 'rust': 'rust',
  };

  private readonly skillDescs: Record<string, string> = {
    'angular': 'Google\'s framework for scalable SPAs',
    'react': 'UI library for component-based interfaces',
    'react native': 'Build native mobile apps with React',
    'next.js': 'React framework with SSR & SSG support',
    'typescript': 'Type-safe JavaScript at any scale',
    'scss / css': 'Styling with variables and nesting',
    'tailwind css': 'Utility-first CSS framework',
    'go (golang)': 'Fast, statically typed backend language',
    'node.js': 'JavaScript runtime for server-side apps',
    'php / laravel': 'PHP framework for elegant web apps',
    'graphql': 'Flexible API query language',
    'rest api': 'Standard HTTP API architecture',
    'mysql': 'Relational database management system',
    'postgresql': 'Advanced open-source SQL database',
    'mongodb': 'Flexible NoSQL document database',
    'redis': 'In-memory data store and cache',
    'flutter': 'Google\'s cross-platform UI toolkit',
    'docker': 'Containerized application deployment',
    'git / github': 'Version control & code collaboration',
    'google cloud': 'Google\'s cloud computing platform',
    'aws': 'Amazon\'s cloud services ecosystem',
    'figma': 'Collaborative UI design tool',
    'openai api': 'Integrate GPT models into products',
    'langchain': 'Framework for LLM-powered applications',
    'python': 'Versatile language for AI & scripting',
    'prompt eng.': 'Crafting effective AI model prompts',
  };

  getSkillIcon(name: string): string {
    const slug = this.skillIcons[name.toLowerCase().trim()];
    return slug ? `https://cdn.simpleicons.org/${slug}` : '';
  }

  getSkillDesc(name: string): string {
    return this.skillDescs[name.toLowerCase().trim()] ?? '';
  }

  marqueeItems(items: Skill[]): Skill[] {
    const minItems = 8;
    const copies = Math.max(2, Math.ceil(minItems / items.length));
    const base = Array.from({ length: copies }, () => items).flat();
    return [...base, ...base];
  }

  skillGroups(skills: Skill[]): { category: string; items: Skill[] }[] {
    const map = new Map<string, Skill[]>();
    for (const skill of skills) {
      const key = skill.category || 'Other';
      if (!map.has(key)) map.set(key, []);
      map.get(key)!.push(skill);
    }
    return Array.from(map, ([category, items]) => ({ category, items }));
  }

  ngOnDestroy(): void {
    this.titleService.setTitle('Diosys — Premium Technology Solutions');
  }

  submitContact(): void {
    if (!this.contactForm.clientName || !this.contactForm.clientEmail || !this.contactForm.clientPhone || !this.contactForm.messageBody) {
      this.contactError.set('Please fill in all required fields.');
      return;
    }
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(this.contactForm.clientEmail)) {
      this.contactError.set('Please enter a valid email address.');
      return;
    }
    if (this.contactForm.clientPhone.replace(/[^0-9]/g, '').length < 5) {
      this.contactError.set('Please enter a valid WhatsApp number.');
      return;
    }
    this.contactError.set('');
    this.sending.set(true);
    this.publicService.sendContact(this.contactForm).subscribe({
      next: () => {
        this.sending.set(false);
        this.sent.set(true);
        this.contactForm = { clientName: '', clientEmail: '', clientPhone: '', subject: '', messageBody: '' };
      },
      error: () => {
        this.sending.set(false);
        this.contactError.set('Something went wrong. Please try again.');
      },
    });
  }

  scrollTo(id: string): void {
    const el = document.getElementById(id);
    if (!el) return;
    const y = el.getBoundingClientRect().top + window.scrollY - 88;
    window.scrollTo({ top: y, behavior: 'smooth' });
  }
}
