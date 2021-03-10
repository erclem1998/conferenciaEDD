import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { EstudiantesService } from "../../services/estudiantes/estudiantes.service";
import { Estudiante } from "../../models/estudiante/estudiante";

@Component({
  selector: 'app-crear-estudiante',
  templateUrl: './crear-estudiante.component.html',
  styleUrls: ['./crear-estudiante.component.css']
})
export class CrearEstudianteComponent implements OnInit {

  carnet= new FormControl('')
  nombres = new FormControl('')
  apellidos= new FormControl('')
  cui= new FormControl('')
  correo= new FormControl('')

  constructor(private estudianteService: EstudiantesService) { }

  ngOnInit(): void {
  }

  crearEstudiante(){
    const estudiante: Estudiante={
      carnet:Number(this.carnet.value),
      nombres:this.nombres.value,
      apellidos:this.apellidos.value,
      cui:String(this.cui.value),
      correo:this.correo.value
    }
    console.log(this.cui.value)
    this.estudianteService.postEstudiante(estudiante).subscribe((res:any)=>{
      console.log("Estudiante Creado")
    })

  }

  vamosaver(){
    console.log(this.nombres.value, " ", this.carnet.value)
  }

}
