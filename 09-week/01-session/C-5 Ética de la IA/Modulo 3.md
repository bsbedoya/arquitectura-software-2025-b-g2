# Módulo 5: Ética de IA

## ¿Qué es la solidez?

### Introducción

La IA se utiliza cada vez más para tomar decisiones cruciales, por lo que es vital que sea **segura y sólida**.  

Cuando la IA es **sólida**, puede manejar con eficacia **condiciones excepcionales**, como anomalías en la entrada de datos o ataques maliciosos, sin causar daños involuntarios.  

---

### Objetivos del curso

Después de completar este módulo, debería ser capaz de:  

- Describir la solidez frente a adversarios en la IA.  
- Explicar cómo un adversario puede influir en un sistema de IA.  
- Clasificar los ataques de adversario.  

---

### Conocer al equipo

En esta historia aprenderemos sobre **solidez** de la mano de:  

- **Charlie (ella)** – Directora de tecnología (CTO)  
- **Manuel (ellos/ellas)** – Responsable de privacidad (CPO)  
- **Prashant (él)** – Responsable de datos (CDO)  

La historia muestra cómo proteger los sistemas de IA de los **ataques de adversario**, centrándose en una empresa de diagnóstico médico que desarrolla una aplicación para **detectar cáncer de pulmón**.  

---

### Identificar el problema

Se acerca la fecha de lanzamiento de la aplicación.  
Charlie revisa el plan con su equipo cuando Manuel interrumpe:  

> **Manuel:** “AY-APP se enfrenta a una demanda porque su aplicación diagnosticó erróneamente enfermedades a muchos pacientes. El problema parece ser un ataque de adversario a su modelo de IA”.  

Charlie considera inmediatamente la posibilidad de ataques de adversario y pregunta:  

> “¿Todo el mundo sabe qué es un ataque de adversario?”  

Prashant se encarga de explicar los conceptos clave.  

---

### Explicar el problema

**Ataques de adversario:**  

- Se realizan intencionadamente en sistemas de IA para lograr objetivos maliciosos aprovechando vulnerabilidades.  
- El **objetivo** es influir negativamente sobre el rendimiento del sistema, explotar los datos y corromper la lógica del modelo.  
- La persona que explota estas vulnerabilidades se denomina **adversario**.  

Ejemplo: Un adversario añade pequeñas **perturbaciones o ruido** a una radiografía.  

- La perturbación puede ser **imperceptible** para los humanos.  
- Sin embargo, la IA puede generar una **predicción incorrecta o no deseada**.  

---

### Objetivos finales del adversario

Algunos objetivos posibles del adversario incluyen:  

1. Acceder a **información personal** de los pacientes (edad, sexo, raza, historial médico, identificación, finanzas).  
2. Añadir **muestras de rayos X maliciosas** a los datos de entrenamiento para manipular el aprendizaje de la IA.  
3. Pronosticar incorrectamente la presencia de enfermedades en pacientes sanos o enfermos.  
4. Añadir **ruido intencionado** para alterar las predicciones.  
5. Conocer el modelo de IA de la empresa y replicarlo para fines propios.  
6. Enviar imágenes maliciosas para recrear datos de entrenamiento basándose en la respuesta del sistema.  

Formas de ataque:  

- Acceder a los datos de entrenamiento y conocer su distribución.  
- Modificar los datos de entrenamiento o de prueba.  
- Acceder al código y parámetros del modelo.  
- Dañar los datos de los usuarios y enviar datos modificados al sistema.  

> Los ataques pueden producirse tanto durante el **entrenamiento** como después del **despliegue** del modelo.  

---

### Tipos de ataques de adversario

#### 1. Envenenamiento

Ocurre durante la **fase de entrenamiento**:  

- Inyectar muestras maliciosas en los datos de entrenamiento.  
- Modificar características o etiquetas de los datos.  
- Alterar la arquitectura, parámetros o lógica del modelo.  

**Consecuencia:** El modelo se vuelve sensible a patrones específicos de los datos maliciosos.  

**Escenario:**  

- Adversario 1 introduce imágenes maliciosas con mayor luminosidad.  
- Adversario 2, en el hospital, modifica entradas basándose en esta información.  
- Resultado: El modelo clasifica incorrectamente pacientes sanos como enfermos.  

---

#### 2. Evasión

Ocurre después del **despliegue del modelo**:  

- Enviar muestras maliciosas al modelo desplegado.  
- Corromper los datos de prueba enviados al modelo.  

**Consecuencia:** Un adversario encuentra cambios en la entrada que hacen que el sistema genere predicciones incorrectas.  

**Escenario:**  

- El adversario ajusta las entradas para que un paciente sano sea etiquetado como enfermo por la IA.  

---

### Abordar el problema

El equipo ahora entiende cómo un modelo de IA puede corromperse por:  

- **Envenenamiento** (durante entrenamiento).  
- **Evasión** (después del despliegue).  

> 🧠 **Conclusión:**  
> Es fundamental que el modelo sea **sólido frente a ataques de adversario**.  
> Charlie y el equipo tienen los conocimientos necesarios para empezar a implementar medidas de solidez.  
