export class Estudiante {
    
    carnet: number
    nombres: string
    apellidos: string
    cui: string
    correo: string

    constructor(_carnet:number,_nombres:string, _apellidos:string, _cui:string, _correo:string){
        this.carnet=_carnet
        this.nombres=_nombres
        this.apellidos=_apellidos
        this.cui=_cui
        this.correo=_correo
    }

}
