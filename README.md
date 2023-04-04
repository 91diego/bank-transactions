# TRANSACTIONS

Ejecutar el comando ´cp .env .example.env´ para crear el archivo .env  con las variables de entorno necesarias.

Para correr el proyecto se deben ejecutar los comandos del archivo ´Makefile´

Dentro de este archivo existen comandos para levantar la base de datos, levantar el servicio para ejecutar las migraciones
a la base de datos.

Para ejecutar una migracion es importante crear un archivo con el script de postgresql que crea la tabla. Este archivo debe ser
creado dentro del directorio database -> migrator -> migrations.
El archivo debe ser nombrado de la siguiente de la siguiente manera:

´#_nombre_archivo_up.sql´
´#_nombre_archivo_down.sql´

El digito del nombre del archivo debe ser el consecutivo de la migracion anterior.

El servidor para ejecutar el proyecto se inicia con ´go run main .´, por default el proyecto corre en el puerto PORT 8089, mismo
que se puede modificar desde las variables de entorno.