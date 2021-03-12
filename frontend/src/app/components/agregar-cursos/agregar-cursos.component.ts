import { Component, OnInit } from '@angular/core';
import { EstudiantesService } from "../../services/estudiantes/estudiantes.service";
import { Curso } from "../../models/curso/curso";
import { Estudiante } from "../../models/estudiante/estudiante";

@Component({
  selector: 'app-agregar-cursos',
  templateUrl: './agregar-cursos.component.html',
  styleUrls: ['./agregar-cursos.component.css']
})
export class AgregarCursosComponent implements OnInit {

  lista_estudiantes: number[]=[]
  lista_cursos: Curso[] = []
  estudiante: Estudiante
  mostrarMensajeError=false
  mensajeError = ''

  constructor(private estudiateService:EstudiantesService) {
    this.estudiateService.getListaCarnets().subscribe((dataList:any)=>{
      this.lista_estudiantes=dataList.listacarnets
      //console.log(dataList)
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError='No se pudo cargar la lista de carnets'
    })
  }

  ngOnInit(): void {
  }

  mostrarCursos(carnet: number){
    this.estudiateService.getInfoEstudiante(carnet).subscribe((dataList:Estudiante)=>{
      this.estudiante=dataList
      console.log(this.estudiante)
      this.mostrarMensajeError=false
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError='Error'
    })
  }

  desactivarMensaje(){
    //this.mostrarMensaje=false
    this.mostrarMensajeError=false
  }

  mensaje(carnet){
    console.log("carnet "+carnet)
  }

}
