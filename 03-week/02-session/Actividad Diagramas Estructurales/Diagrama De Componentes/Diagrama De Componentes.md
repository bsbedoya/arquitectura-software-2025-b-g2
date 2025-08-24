# Diagrama de Componentes en UML

## ¿Qué es?

Un **diagrama de componentes** es un tipo de diagrama estructural de UML (Unified Modeling Language) que se utiliza para **representar la arquitectura física** de un sistema. Muestra cómo los distintos componentes de software (módulos, subsistemas, bibliotecas, servicios, etc.) están organizados y cómo interactúan entre sí.

Se centra en el **aspecto estático** de la implementación, es decir, qué piezas de software existen y cómo se relacionan, más que en los flujos dinámicos de ejecución.

---

## Objetivo

- Visualizar la **arquitectura lógica y física** de un sistema.
- Mostrar la **modularidad** del software.
- Especificar las **dependencias** entre los componentes.
- Servir de guía para desarrolladores y arquitectos de software en la fase de **implementación y despliegue**.

---

## Elementos principales

1. **Componente (`[Componente]`)**  
   
   - Representa una pieza de software reutilizable o autónoma (ejemplo: un módulo, un servicio, un repositorio).  
   - En UML se suele dibujar como un rectángulo con dos pequeños rectángulos a la izquierda (icono de componente). En PlantUML, simplemente se define con `[NombreComponente]`.

2. **Interfaz (`() Interfaz`)**  
   
   - Define un contrato o un punto de acceso que un componente ofrece o requiere.  
   - Puede representarse con una notación tipo "lollipop" (círculo) o un "socket".

3. **Paquete (`package`)**  
   
   - Agrupa un conjunto de componentes relacionados.  
   - Se utiliza para organizar la arquitectura en capas o subsistemas (ejemplo: `Frontend`, `Backend`, `Base de Datos`).

4. **Relaciones / Dependencias (`-->`)**  
   
   - Indican que un componente depende de otro para funcionar.  
   - Ejemplo: un `ServicioCarrito` depende de un `RepositorioProductos`.

---

## Características

- Muestra **la implementación física** del sistema, no solo el diseño lógico.  
- Es útil en la fase de **desarrollo y despliegue** del software.  
- Permite identificar **acoplamientos y dependencias** entre los módulos.  
- Favorece la **modularidad y mantenibilidad** del sistema.  

---

## Ventajas

- Claridad en la **arquitectura del sistema**.  
- Facilita la **comunicación** entre equipos técnicos y no técnicos.  
- Ayuda en la **planificación del desarrollo** y la asignación de responsabilidades.  
- Mejora la **documentación del sistema**.  

---

## Ejemplo de uso típico

Un sistema web puede representarse con un diagrama de componentes dividiendo en:

- **Frontend**: WebApp, controladores de interfaz.  
- **Backend**: Servicios, controladores, repositorios.  
- **Base de datos**: Tablas o una única base central.  
