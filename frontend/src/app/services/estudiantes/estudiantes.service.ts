import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { baseURL } from "../../apiURL/baseURL";
import { Observable } from "rxjs";
import { Curso } from "../../models/curso/curso";

@Injectable({
  providedIn: 'root'
})
export class EstudiantesService {
  
  
  constructor(private http: HttpClient) { }
  
  getListaCarnets():Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.get<any>(baseURL + 'listaEstudiantes', httpOptions);
  }

  getListaCursos(carnet):Observable<Curso[]>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.post<Curso[]>(baseURL + 'cursosEstudiante', {carnet: carnet}, httpOptions);
  }

  postEstudiante(estudiante):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.post<Curso[]>(baseURL + 'crearEstudiante', estudiante, httpOptions);
  }

}
