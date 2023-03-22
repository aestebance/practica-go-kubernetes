# práctica Go & k8s

Aplicación para ser utilizada en los yml de k8s de esta práctica.

# Descripción

Este proyecto crea una imagen de docker con dos aplicaciones: cliente y servidor. Por defecto se lanza la aplicación `server`. Es necesario sobreescribir el entrypoint con `client` para que lance la aplicación cliente.

`server`: La aplicación server lanza un servidor web que escucha peticiones http en el puerto 8000 y expone dos endpoints, `/health` y `/echo`. El endpoint `/health` devuelve un json con el estado de la aplicación, mientras que el endpoint `/echo` devuelve un un json con el hostname de la máquina y el mensaje 'Esto es un gato'. El animal puede ser sobreescrito con la variable de entorno $ANIMAL.

`client`: La aplicación cliente lanza un cliente web que hace peticiones http al servidor web. El cliente se puede configurar con las siguientes variables de entorno: $SLEEP_TIME - Tiempo de espera entre peticiones, $SERVER_URL - Dirección y puerto del servidor (http://localhost:8000), $ENDPOINT - Endpoint del servidor.

# Descripción detallada del servidor

Utiliza el entrypoint `server` para iniciar el sistema como servidor web. La aplicación expone los siguientes endpoints:

| Endpoint  | Descripción                                                                                       |
|-----------|---------------------------------------------------------------------------------------------------|
| `/echo`   | Devuelve un json con los siguientes datos: `{ hostname: <hostname>, message: 'Esto es un gato' }` |
| `/health` | muestra el mensaje OK si el servidor está funcionando correctamente.                              |

Todos los endpoints devuelven la respuesta en formato json.

| Variable de entorno | Descripción                      |Tipo| Valor por defecto |
|---------------------|----------------------------------|------|-------------------|
| `ANIMAL`            | Modifica el animal de respuesta. | String | `gato`            |

# Descripción detallada del cliente

Utiliza el entrypoint `client` para iniciar el sistema como cliente web. La aplicación realiza peticiones http hacia el enlace endpoint por intervalos de tiempo. Para configurar la aplicación cliente utilice las siguientes variables de entorno:

|Variable de entorno| Descripción                                                                       |Tipo|Valor por defecto|
|-----|-----------------------------------------------------------------------------------|------|---|
|`SLEEP_TIME`| Intervalo de tiempo entre peticiones.                                             | String | `1s` |
|`SERVER_URL`| Hostname donde serán realizadas las peticiones. Ejemplo: `http://localhost:8000`. | String | `"http://localhost:8000"` |
|`ENDPOINT`| Endpoint del servidor web.                                                        | String | `""` |
