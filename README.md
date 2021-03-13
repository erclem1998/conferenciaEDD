# **CONFERENCIA EDD 2021: DESARROLLO DE MODERNAS APLICACIONES WEB CON ANGULAR**

## Objetivo del ejemplo
El objetivo principal de este ejemplo, es introducir, a los estudiantes, el uso del framework 
Angular, que hoy en día es uno de los mas utilizados en el desarrollo de Aplicaciones Web.

## ¿Que encontrarás?
1. Servidor web desarollado con Golang.
2. Aplicación Web desarrollada con Angular.

## Lo que tienes que saber antes de hacer uso de este ejemplo
* Node v14.15.4
* Angular v11.2.2
* Go v1.16

## Acerca del Servidor Web
Este ejemplo se basa en el uso de un arbol binario de búsqueda, donde se almacenan los datos de
estudiantes y los cursos que ha aprobado.

Cabe destacar que el servidor web hace uso de los siguientes paquetes:

**1. github.com/gorilla/mux**

Utilizado para generar el enrutador de solicitudes HTTP.
Se debe instalar de la siguiente manera.

**go get github.com/gorilla/mux**

**2. github.com/gorilla/handlers**

Utilizado para manejar los CORS del servidor, permitiendo de esta manera
el acceso a las aplicaciones web.

**go get github.com/gorilla/handlers**