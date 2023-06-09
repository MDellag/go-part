#Ag GOLANG - PRACTICE (JRs)

## Descripcion

El proyecto solo cuenta con un archivo `main.go` en el cual contiene ya implementado un menu de consola
donde el usuario final (el admin) debe dar de alta otros usuarios.

para eso el menu tiene opciones:
 * 1 - Lista los usuarios cargados
 * 2 - Agrega un nuevo usuario
 * 3 - Modifica un usuario
 * 4 - Remueve un usuario
 * 0 - Muestra el menu de opciones
 * 9 - Sale del programa

### POR CADA PUNTO, SE DEBERA CREAR UN NUEVO BRANCH (RAMA) DE GIT.


## 1 - Listar Usuarios
Esta funcionalidad, aun no se encuentra implementada, por lo que debemos implementarla.

Crear un package `users` y dentro un archivo `users.service.go` que se encargue de la logica para mostrar los usuarios.

Asi mismo, dentro del package `users` tambien crear un archivo `users.dto.go` que se encargue de tener el Struct del tipo User
el cual tendra los siguientes atributos:
- ID
- NOMBRE
- DNI

Cabe recalcar, que si el ID y/o DNI ingresado ya existe por otro usuario, deberia retornar un mensaje de error `el ID/DNI ya existe!` y NO ser agregado a la lista.

nota: harcodear almenos 2 usuarios para probar la funcionalidad del programa cuando en el menu, se elija la opcion `1`

## 2 - Agregar Usuarios
Cuando se ingrese la opcion `2` en la consola. El programa debera solicitar los datos para agregar un nuevo Usuario.

Una vez cargado el Usuario, al ingresar la opcion `1` debera listar los usuarios incluyendo los cargados.


## 3 - Modificar Usuario
Al ingresar `3` el programa solicitara el ID del usuario a modificar, y SI existe, entonces solicitara los datos a modificar por consola.

## 4 - Baja de Usuario

Al ingresar `4` el programa solicitara un ID de usuario, y, si existe, lo borrara del sistema. Si no existe debera mostrar un mensaje de error
Comentando que el usuario no existe.


### POR ULTIMO, UNA VEZ IMPLEMENTADOS TODOS LOS PUNTOS, PROBAR EL PROGRAMA AL 100% TODAS SUS FUNCIONALIDADES JUNTAS.

