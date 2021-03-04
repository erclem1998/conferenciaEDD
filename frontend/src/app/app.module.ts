import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { CrearEstudianteComponent } from './components/crear-estudiante/crear-estudiante.component';
import { InicioComponent } from './components/inicio/inicio.component';
import { AgregarCursosComponent } from './components/agregar-cursos/agregar-cursos.component';

@NgModule({
  declarations: [
    AppComponent,
    CrearEstudianteComponent,
    InicioComponent,
    AgregarCursosComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
