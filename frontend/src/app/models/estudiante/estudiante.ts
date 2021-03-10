import { Curso } from "../curso/curso";
export class Estudiante {
    
    carnet: number
    nombres: string
    apellidos: string
    cui: string
    correo: string
    listaCursos: Curso[]

    constructor(_carnet:number,_nombres:string, _apellidos:string, _cui:string, _correo:string, _listacursos:Curso[]){
        this.carnet=_carnet
        this.nombres=_nombres
        this.apellidos=_apellidos
        this.cui=_cui
        this.correo=_correo
        this.listaCursos=_listacursos
    }

}
