import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class PriceService {
  private apiUrl = 'http://localhost:5001/api';

  constructor(private http: HttpClient) {}

  getCasePrice(caseName: string): Observable<any> {
    return this.http.get(`${this.apiUrl}/price/${encodeURIComponent(caseName)}`);
  }

  getCaseHistory(caseName: string): Observable<any[]> {
    return this.http.get<any[]>(`${this.apiUrl}/history/${encodeURIComponent(caseName)}`);
  }

  setPurchasePrice(purchaseData: { caseName: string, purchasePrice: number, purchaseDate: string }): Observable<any> {
    return this.http.post(`${this.apiUrl}/purchase`, purchaseData);
  }

  getAllHistory(): Observable<any[]> {
    return this.http.get<any[]>(`${this.apiUrl}/history`);
  }
}