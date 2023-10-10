# go-medium-shapes
Proyecto desarrollado en Golang. El desarrollo consiste en una lambda para la creación de figuras y generación de archivo txt con listado de figuras.

En el caso de la generación de archivo txt recibe como parámetro de entrada el campo "tipo" que debe ser ELLIPSE, RECTANGLE o TRIANGLE. Con ese valor consulta a la table devShapes y filtra los resultados. Luego genera un archivo txt que sube al folder SHAPES/ del S3.

En el caso de la creación recibe como parámetros de entrada el campo "tipo" que debe ser ELLIPSE, RECTANGLE o TRIANGLE, el campo id (tipo string del "1" al "12"), y los campos a y b de tipo floar64. 


Comandos para generar el binario y comprimir en .zip:

1. cd cmd
2. GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o goMediumShapes main.go
3. zip -rm ../build/package/goMediumShapes.zip goMediumShapes

Opcionalmente se puede ejecutar los scripts "build mac" o "build windows" (dependiendo el SO) para generar el binario y comprimirlo. Estos ejecutan los archivos buildMac.bat y buildWindows.bat respectivamente. Es probable que el archivo .bat a ejecutar requiera permisos de ejecución (chmod 774)


Pasos para desplegar la lambda en AWS:

1. Crear la función en el servicio de Lambda de aws con el nombre goMediumShapes
2. Subir el archivo comprimido a la función
3. Cambiar el nombre del handler en "Runtime settings" por goMediumShapes


# Datos útiles
Tabla: devShapes
S3: bucket uala-arg-labssupport-dev, folder SHAPES/
IAM Role: GoValidationServiceRole-dev

# Consideraciones
- Logs agregados solo en algunas partes del proceso.
- Notar que el id de entrada para la creación de figuras es de tipo string.
- Datos de configuración están hardcodeados (constants.go).

# Mejoras pendientes
- Cambiar scan por query.
- El uso del file puede mejorarse creando un objeto File y asociarle métodos.
- Ubicación de logs y datos que contiene.
- Las constantes contenidas en constants.go deberían ser valores de configuración (no hardcodeadas).
- Casos de prueba aún son básicos y no tienen mockup.
- Interacción con S3 por medio de repositorio.
