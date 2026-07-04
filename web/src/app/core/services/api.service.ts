import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

import { environment } from '../../../enviroments/environment';
import { ApiResponse } from '../models/api-response.model';

/**
 * Thin wrapper around HttpClient that prefixes the API base URL and unwraps
 * the standard response envelope, returning only the `result` payload.
 */
@Injectable({ providedIn: 'root' })
export class ApiService {
  private readonly baseUrl = environment.apiUrl;

  constructor(private http: HttpClient) {}

  get<T>(path: string): Observable<T> {
    return this.http
      .get<ApiResponse<T>>(`${this.baseUrl}${path}`)
      .pipe(map((response) => response.result));
  }

  post<T>(path: string, body: unknown): Observable<T> {
    return this.http
      .post<ApiResponse<T>>(`${this.baseUrl}${path}`, body)
      .pipe(map((response) => response.result));
  }

  put<T>(path: string, body: unknown): Observable<T> {
    return this.http
      .put<ApiResponse<T>>(`${this.baseUrl}${path}`, body)
      .pipe(map((response) => response.result));
  }

  patch<T>(path: string, body: unknown): Observable<T> {
    return this.http
      .patch<ApiResponse<T>>(`${this.baseUrl}${path}`, body)
      .pipe(map((response) => response.result));
  }

  delete<T>(path: string): Observable<T> {
    return this.http
      .delete<ApiResponse<T>>(`${this.baseUrl}${path}`)
      .pipe(map((response) => response.result));
  }

  upload<T>(path: string, file: File): Observable<T> {
    return this.uploadWithFields<T>(path, file, {});
  }

  uploadWithFields<T>(path: string, file: File, fields: Record<string, string>): Observable<T> {
    const formData = new FormData();
    formData.append('file', file);
    Object.entries(fields).forEach(([key, value]) => formData.append(key, value));
    return this.http
      .post<ApiResponse<T>>(`${this.baseUrl}${path}`, formData)
      .pipe(map((response) => response.result));
  }
}
