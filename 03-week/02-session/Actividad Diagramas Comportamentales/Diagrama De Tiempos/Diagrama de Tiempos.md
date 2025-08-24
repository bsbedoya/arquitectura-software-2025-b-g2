# Documentación: Diagrama de Tiempos UML

## 1. Definición

El **diagrama de tiempos** en UML es un tipo de diagrama de interacción que se utiliza para mostrar el **cambio de estado o valor de un objeto a lo largo del tiempo**. Se centra en la evolución temporal de los objetos, describiendo cómo varían sus estados, condiciones o mensajes en función de un eje temporal.

Es especialmente útil para modelar **sistemas sensibles al tiempo**, como sistemas embebidos, telecomunicaciones, procesos en tiempo real o simulaciones.

---

## 2. Objetivo

- Mostrar la **secuencia temporal de eventos** que afectan a los objetos.
- Representar el **tiempo absoluto o relativo** en la interacción de los elementos.
- Analizar **rendimiento, latencia y sincronización** en un sistema.
- Identificar **restricciones temporales** y dependencias.

---

## 3. Elementos principales

1. **Líneas de vida (Lifelines):**
   
   - Representan los participantes (objetos o actores) en el diagrama.
   - Se dibujan de forma horizontal.

2. **Eje de tiempo:**
   
   - Generalmente se coloca en la parte inferior (horizontal).
   - Puede representar tiempo absoluto (en segundos, minutos) o relativo.

3. **Estados o valores:**
   
   - Cada línea de vida puede mostrar cómo cambian los estados o valores a lo largo del tiempo.

4. **Mensajes:**
   
   - Indican la comunicación entre los participantes, sincronizada con el tiempo.

5. **Restricciones temporales:**
   
   - Se anotan para especificar tiempos de espera, retardos, límites de tiempo o secuencias críticas.

6. **Duración y latencia:**
   
   - Se utilizan para medir intervalos de tiempo entre eventos o la duración de un estado.

---

## 4. Características

- Es un **diagrama dinámico**, centrado en la dimensión **temporal**.
- Usa **notación gráfica horizontal** (tiempo en el eje X).
- Permite modelar **tiempos de respuesta, sincronización y concurrencia**.
- Complementa a los diagramas de **secuencia** y **comunicación** al añadir la variable temporal explícitamente.

---

## 5. Ventajas

- Claridad en la representación de **eventos en tiempo real**.
- Permite detectar **cuellos de botella** y problemas de sincronización.
- Útil para **simulación y validación** de sistemas sensibles al tiempo.
- Apoya en la **optimización de rendimiento**.

---

## 6. Limitaciones

- Puede volverse complejo si hay muchos elementos o valores a representar.
- Requiere una buena definición de **medidas de tiempo**.
- No sustituye a otros diagramas de interacción, sino que los complementa.

---

## 7. Ejemplo práctico

Supongamos un sistema de **transmisión de datos**:

- Un **sensor** envía datos a un **procesador** cada 2 segundos.
- El procesador tarda 0.5 segundos en analizar los datos.
- Si los datos están fuera de rango, se envía una alerta a un **sistema de monitoreo**.

Representación:

- Línea de vida del sensor con estados de "En reposo" → "Transmitiendo".
- Línea de vida del procesador mostrando "Esperando" → "Procesando".
- Línea de vida del sistema de monitoreo mostrando "Inactivo" → "Recibiendo alerta".

---

## 8. Diferencias con otros diagramas

- **Diagrama de secuencia:** muestra mensajes en orden, pero no detalla valores de tiempo.
- **Diagrama de comunicación:** enfatiza relaciones entre objetos, sin eje temporal explícito.
- **Diagrama de tiempos:** se centra en cómo los estados y valores cambian **a lo largo del tiempo**.

---
