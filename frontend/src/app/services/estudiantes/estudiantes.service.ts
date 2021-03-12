import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { baseURL } from "../../apiURL/baseURL";
import { Observable } from "rxjs";
import { Curso } from "../../models/curso/curso";
import { Estudiante } from "../../models/estudiante/estudiante";

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

  getInfoEstudiante(carnet):Observable<Estudiante>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.post<Estudiante>(baseURL + 'cursosEstudiante', {carnet: carnet}, httpOptions);
  }

  postEstudiante(estudiante):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.post<Curso[]>(baseURL + 'crearEstudiante', estudiante, httpOptions);
  }

  postCursoAprobado(data):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.post<any>(baseURL + 'insertarCurso', data, httpOptions);
  }

}
