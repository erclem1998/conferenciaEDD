import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CrearEstudianteComponent } from "./components/crear-estudiante/crear-estudiante.component";
import { InicioComponent } from "./components/inicio/inicio.component";
import { AgregarCursosComponent } from "./components/agregar-cursos/agregar-cursos.component";

const routes: Routes = [
  {
    path: 'crearEstudiante',
    component: CrearEstudianteComponent,
  },
  {
    path: '',
    component: InicioComponent,
  },
  {
    path: 'agregarCursos',
    component: AgregarCursosComponent,
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
