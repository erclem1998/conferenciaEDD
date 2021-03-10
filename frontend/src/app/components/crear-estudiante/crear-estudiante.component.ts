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
  mostrarMensaje=false
  mostrarMensajeError=false

  constructor(private estudianteService: EstudiantesService) { }

  ngOnInit(): void {
  }

  crearEstudiante(){
    const estudiante: Estudiante={
      carnet:Number(this.carnet.value),
      nombres:this.nombres.value,
      apellidos:this.apellidos.value,
      cui:String(this.cui.value),
      correo:this.correo.value,
      listaCursos:[]
    }
    console.log(this.cui.value)
    this.estudianteService.postEstudiante(estudiante).subscribe((res:any)=>{
      this.mostrarMensaje=true
      this.carnet.setValue(0)
      this.nombres.setValue("")
      this.apellidos.setValue("")
      this.cui.setValue("")
      this.correo.setValue("")
      console.log("Estudiante Creado")
    },(err)=>{
      this.mostrarMensajeError=true
    })

  }

  desactivarMensaje(){
    this.mostrarMensaje=false
    this.mostrarMensajeError=false
  }

  vamosaver(){
    console.log(this.nombres.value, " ", this.carnet.value)
  }

}
