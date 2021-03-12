import { Component, OnInit } from '@angular/core';
import { EstudiantesService } from "../../services/estudiantes/estudiantes.service";
import { Curso } from "../../models/curso/curso";
import { Estudiante } from "../../models/estudiante/estudiante";
import { FormControl } from '@angular/forms';

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
  mostrarMensaje=false
  mensajeError = ''
  carnet= new FormControl('')
  nombres = new FormControl('')
  apellidos= new FormControl('')
  cui= new FormControl('')
  correo= new FormControl('')
  anio= new FormControl('')
  nota= new FormControl('')
  opcion: string;
  cursosDisponibles: string[]=[
    "IPC1",
    "Logica de Sistemas",
    "MC1",
    "LFP",
    "MC2",
    "IPC2",
    "OLC1",
    "EDD",
    "Organizacion Computacional",
    "OLC2",
    "AC1",
    "MIA",
    "AC2",
    "REDES 1",
    "BD1"
  ]
  cursosMostrar: string[]=[]

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
      this.nombres.setValue(this.estudiante.nombres)
      this.apellidos.setValue(this.estudiante.apellidos)
      this.carnet.setValue(this.estudiante.carnet)
      this.cui.setValue(this.estudiante.cui)
      this.correo.setValue(this.estudiante.correo)
      this.cursosMostrar=[]
      this.cursosMostrar=this.quitarCursosAprobados(this.estudiante.listaCursos)
      console.log(this.estudiante)
      this.mostrarMensajeError=false
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError='Error'
    })
  }

  quitarCursosAprobados(aprobados: Curso[]){
    var cursos:string[] = []
    for(let i=0;i<this.cursosDisponibles.length;i++){
      var flag=false
      for(let j=0;j<aprobados.length;j++){
        if(this.cursosDisponibles[i]==aprobados[j].curso){
          flag=true
        }
      }
      if(flag==false){
        cursos.push(this.cursosDisponibles[i]);
      }
    }
    return cursos
  }

  guardarCurso(){
    var cursoInsertar: Curso={
      curso:this.opcion,
      anio: this.anio.value,
      nota: this.nota.value
    }
    console.log(cursoInsertar);
    var data ={
      carnet:this.carnet.value,
      curso:cursoInsertar
    }
    this.estudiateService.postCursoAprobado(data).subscribe((res:Estudiante)=>{
      this.anio.setValue("")
      this.nota.setValue("")
      this.estudiante=res
      this.cursosMostrar=this.quitarCursosAprobados(this.estudiante.listaCursos)
      this.mostrarMensaje=true
    }, (err)=>{
      this.mostrarMensajeError=true
      this.mensajeError='No se pudo guardar el curso aprobado'
    })
  }

  desactivarMensaje(){
    this.mostrarMensaje=false
    this.mostrarMensajeError=false
  }

  mensaje(carnet){
    console.log("carnet "+carnet)
  }

}
