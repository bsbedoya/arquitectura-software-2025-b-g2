# Documentación completa: Diagrama de Estructura Compuesta (UML)

## 1) ¿Qué es un diagrama de estructura compuesta?

El **diagrama de estructura compuesta** (Composite Structure Diagram) en UML describe la **estructura interna de una clase, componente o colaboración** y muestra cómo sus partes interactúan para cumplir un comportamiento. Permite detallar **qué elementos internos conforman una entidad mayor** y **cómo se comunican** entre sí.

Se utiliza para **modelar arquitecturas internas** y la **colaboración entre objetos** dentro de un contexto definido.

---

## 2) Objetivos principales

- Mostrar la **estructura interna** de una clase o componente.
- Definir las **partes y roles** que participan en una colaboración.
- Especificar **puertos y conectores** usados para la interacción.
- Aclarar **cómo los elementos internos trabajan juntos** para proveer funcionalidades.
- Documentar el diseño detallado de componentes en arquitecturas complejas.

---

## 3) Elementos clave del diagrama

### 3.1 Clases y componentes estructurados

- Una **clase estructurada** es aquella que contiene una **estructura interna**.
- Puede representarse con un rectángulo con subdivisión para sus **partes**.

### 3.2 Partes (Parts)

- Son **instancias** o roles dentro de una clase o componente.
- Representan la composición interna (ejemplo: un `ControladorCarrito` tiene partes `GestorPedidos` y `ValidadorUsuario`).
- Notación: rectángulo dentro del contenedor, con formato `nombre:Tipo`.

### 3.3 Puertos (Ports)

- Puntos de interacción en el borde de una clase/componente.
- Definen cómo un elemento interno puede **enviar/recibir servicios o señales**.
- Se representan como **pequeños cuadrados en el borde** del rectángulo.

### 3.4 Conectores (Connectors)

- Relaciones de comunicación entre **partes** o **puertos**.
- Modelan flujos de mensajes, dependencias, invocaciones de servicios.
- Notación: línea sólida entre dos partes o puertos.

### 3.5 Interfaces / Roles

- Interfaces provistas (`lollipop`) y requeridas (`socket`) pueden asociarse a puertos.
- Indican **contratos de servicio** entre partes internas y externas.

### 3.6 Colaboraciones

- Agrupaciones de roles que trabajan juntos para realizar un objetivo.
- Permiten describir **interacciones contextuales** entre partes.

---

## 4) Notación y estereotipos útiles

- **Clase estructurada**: rectángulo con subdivisión en partes internas.
- **Parte**: rectángulo con notación `instancia:Tipo`.
- **Puerto**: cuadrado pequeño en borde, conectado a interfaces.
- **Conector**: línea sólida uniendo puertos o partes.
- **Interfaces**: lollipop (provista) o socket (requerida).
- **Estereotipos** comunes:
  - `«component»`, `«controller»`, `«entity»`, `«boundary»`.

---

## 5) Qué detallar (checklist)

- [ ] ¿Qué clase o componente tiene estructura interna?
- [ ] ¿Qué partes internas lo conforman?
- [ ] ¿Qué roles o instancias se modelan?
- [ ] ¿Qué puertos de entrada/salida se definen?
- [ ] ¿Cómo se conectan las partes y con qué protocolos/interfaces?
- [ ] ¿Se definen dependencias internas y externas?

---

## 6) Niveles de detalle (según audiencia)

- **Vista general**: partes principales y conexiones básicas.
- **Vista detallada**: puertos, interfaces provistas/requeridas, roles internos.
- **Vista técnica**: protocolos, contratos, dependencias específicas.

---

## 7) Patrones frecuentes

- **Modelo MVC interno**: un componente contiene `Modelo`, `Vista`, `Controlador`.
- **Componentes distribuidos**: un módulo con partes internas que representan microservicios conectados.
- **Dispositivos IoT**: estructura de un nodo con puertos de comunicación (WiFi, sensores, actuadores).

---

## 8) Buenas prácticas

- Usar nombres claros para partes (`instancia:Tipo`).
- Mostrar solo **puertos relevantes** en el borde.
- Asociar interfaces provistas/requeridas a los puertos.
- Evitar exceso de detalle (solo lo necesario para comprender la colaboración).
- Mantener consistencia con otros diagramas (componentes, despliegue).

---

## 9) Errores comunes

- Confundir partes con clases independientes (son roles internos).
- No definir puertos explícitos en interacciones externas.
- Exceso de detalle que complica la lectura.
- No diferenciar entre **interfaces provistas y requeridas**.

---

## 10) Proceso recomendado para construirlo

1. Identificar el **componente o clase** que será modelado internamente.
2. Definir las **partes internas** que lo conforman.
3. Determinar los **puertos** de comunicación hacia el exterior.
4. Establecer las **conexiones internas** (connectors).
5. Asociar **interfaces provistas/requeridas** a los puertos.
6. Revisar coherencia con el **diagrama de componentes**.

---

## 11) Mini-glosario

- **Parte (Part)**: rol interno de un componente.
- **Puerto (Port)**: punto de interacción entre el sistema y el exterior.
- **Conector (Connector)**: relación entre partes o puertos.
- **Colaboración**: conjunto de roles que trabajan juntos.
- **Interface provista**: servicio ofrecido.
- **Interface requerida**: servicio que necesita.
