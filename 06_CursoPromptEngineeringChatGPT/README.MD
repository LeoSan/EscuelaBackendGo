# Curso de Prompt Engineering con ChatGPT

## Clase 1. ChatGPT para prompt engineering

- Herramienta con múltiples usos en múltiples sectores de la economía digital
- Marketing, negocios, software, etc.
- Pueden ayudar con múltiples usos y funciones.
- Aprenderemos diseño de prompts para comunicarnos mejor con esta IA.


## CLASE 2: ¿Para qué usar ChatGPT y prompt engineering? 

**Prompt que es?**
> Son conjunto de instrucciones que nos permite comunicarnos con el servicio de IA, 
para que pueda ejecutar una tarea, ya que se usa LPN lenguaje de programación natural, 
el servicio interpreta esta orden y nos devuelve un resultado aproximado a lo descripto. 

- PD es importante señalar que mientras mas preciso sea tu orden o comunicación esta sera mas acertada. 


**Diferentes usos**

- Resumen de información
- Redacción de textos
- Asistencia personal virtual
- Redactar código de programación
- Toma de decisiones
- Chatbot de atención al cliente
- Todo lo que requiera texto o comprensión de lenguaje se puede optimizar con los distintos modelos de PLN

**Aprenderemos a diseñar preguntas o instrucciones para comunicarnos de manera efectiva con estas IA**
- Prompt: Instrucción a la IA
- Tienen cabida en todos los sectores de la industria
	- Marketing
	- Negocios
	- Software


**ADVERTENCIA**
- Nunca pongas información sensible o confidencial en ChatGPT

**Prompt engineering** 
En si no se considera una ingenieria ya que no es una disciplina, es mas una técnica que Abarca una amplia gama de habilidades y técnicas que son útiles para interactuar con LLM (Large Language Models)"

## Clase 3: Tu primera conversación con ChatGPT

Hola me puedes ayudar en describir los pasos para instalar laravel y usarlo para un servicio web api
¡Claro que sí! Aquí te describo los pasos para instalar Laravel y crear un servicio web API:

Requisitos previos: Antes de instalar Laravel, asegúrate de tener instalado PHP en tu computadora. Puedes descargar la última versión de PHP en el sitio web oficial de PHP. También necesitarás tener instalado Composer, un administrador de paquetes para PHP. Puedes descargar Composer en getcomposer.org.

Instalar Laravel: Una vez que tienes PHP y Composer instalados, puedes instalar Laravel usando Composer. Abre una terminal o línea de comandos y escribe el siguiente comando:

javascript
Copy code
composer global require laravel/installer
Este comando descargará e instalará la última versión de Laravel.

Crear un nuevo proyecto Laravel: Para crear un nuevo proyecto de Laravel, abre una terminal o línea de comandos y navega hasta la carpeta en la que deseas crear el proyecto. Luego, escribe el siguiente comando:

arduino
Copy code
laravel new nombre-de-tu-proyecto
Este comando creará un nuevo proyecto de Laravel con el nombre que hayas elegido.

Crear una nueva ruta API: Para crear una nueva ruta API, abre el archivo routes/api.php en tu editor de texto favorito y agrega la siguiente línea:

javascript
Copy code
Route::get('/ruta', function () {
    return '¡Hola, mundo!';
});
Esta ruta responderá a las solicitudes GET a /ruta y devolverá el texto "¡Hola, mundo!".

Ejecutar el servidor de desarrollo: Para ejecutar el servidor de desarrollo de Laravel, abre una terminal o línea de comandos, navega hasta la carpeta raíz de tu proyecto de Laravel y escribe el siguiente comando:

Copy code
php artisan serve
Este comando ejecutará el servidor de desarrollo de Laravel y podrás acceder a tu aplicación en http://localhost:8000.

Probar la ruta API: Abre un navegador web y visita la URL http://localhost:8000/api/ruta. Deberías ver el texto "¡Hola, mundo!" en tu navegador.

¡Eso es todo! Ahora tienes un servicio web API básico en Laravel. Puedes seguir agregando más rutas y funcionalidades para construir una aplicación web completa.

## Clase 4: Cómo hablarle a ChatGPT: estructura de un prompt
