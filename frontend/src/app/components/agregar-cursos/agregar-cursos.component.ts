import { Component, OnInit } from '@angular/core';
import { EstudiantesService } from "../../services/estudiantes/estudiantes.service";

@Component({
  selector: 'app-agregar-cursos',
  templateUrl: './agregar-cursos.component.html',
  styleUrls: ['./agregar-cursos.component.css']
})
export class AgregarCursosComponent implements OnInit {

  lista_estudiantes: number[]=[]

  constructor(private estudiateService:EstudiantesService) {
    this.estudiateService.getListaCarnets().subscribe((dataList:any)=>{
      this.lista_estudiantes=dataList.listacarnets
      console.log(dataList)
    })
  }

  ngOnInit(): void {
  }

  mensaje(carnet){
    console.log("carnet "+carnet)
  }

}
