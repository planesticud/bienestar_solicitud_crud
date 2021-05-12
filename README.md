# bienestar_solicitud_crud
API de gestión de solicitudes de servicios de bienestar virtual

Integración con

 - `CI`
 - `Drone 1.x`
 - `bienestar_solicitud_crud master/develop`

## Requerimientos
Go version >= 1.9.

## Preparación
Para usar el API, usar el comando:

 - `go get github.com/planesticud/bienestar_solicitud_crud`

## Ejecución
Definir los valores de las siguientes variables de entorno:

 - `BIENESTAR_SOLICITUD_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `BIENESTAR_SOLICITUD_CRUD__PGUSER`: Usuario de la base de datos
 - `BIENESTAR_SOLICITUD_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `BIENESTAR_SOLICITUD_CRUD__PGURLS`: Host de conexión
 - `BIENESTAR_SOLICITUD_CRUD__PGDB`: Nombre de la base de datos
 - `BIENESTAR_SOLICITUD_CRUD__SCHEMA`: Esquema a utilizar en la base de datos

## Ejemplo
BIENESTAR_SOLICITUD_HTTP_PORT=9018 BIENESTAR_SOLICITUD_CRUD__PGUSER=user BIENESTAR_SOLICITUD_CRUD__PGPASS=password BIENESTAR_SOLICITUD_CRUD__PGURLS=localhost BIENESTAR_SOLICITUD_CRUD__PGDB=bd BIENESTAR_SOLICITUD_CRUD__SCHEMA=schema_new bee run

## Modelo BD
![image](https://github.com/planesticud/bienestar_solicitud_crud/blob/master/modelo_bienestar_solicitud_crud.png).
