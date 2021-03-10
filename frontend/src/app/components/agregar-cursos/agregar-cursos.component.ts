import { Component, OnInit } from '@angular/core';
import { EstudiantesService } from "../../services/estudiantes/estudiantes.service";
import { Curso } from "../../models/curso/curso";

@Component({
  selector: 'app-agregar-cursos',
  templateUrl: './agregar-cursos.component.html',
  styleUrls: ['./agregar-cursos.component.css']
})
export class AgregarCursosComponent implements OnInit {

  lista_estudiantes: number[]=[]
  lista_cursos: Curso[] = []

  constructor(private estudiateService:EstudiantesService) {
    this.estudiateService.getListaCarnets().subscribe((dataList:any)=>{
      this.lista_estudiantes=dataList.listacarnets
      console.log(dataList)
    })
  }

  ngOnInit(): void {
  }

  mostrarCursos(carnet: number){
    this.estudiateService.getListaCursos(carnet).subscribe((dataList:Curso[])=>{
      this.lista_cursos=dataList
      console.log(this.lista_cursos)
    })
  }

  mensaje(carnet){
    console.log("carnet "+carnet)
  }

}
