# Módulo 5: Ética de IA

## ¿Qué es la explicabilidad?

### Introducción

Un sistema de IA es **explicable** cuando personas sin formación técnica pueden entender **cómo y por qué** el sistema ha llegado a una predicción o recomendación concreta.  

La explicabilidad es como **mostrar tu trabajo en matemáticas**: permite que todo el mundo vea los pasos que se han seguido para llegar a la respuesta.  

---

### Objetivos del curso

Después de completar este módulo, debería ser capaz de:  

- Describir la explicabilidad.  
- Describir la interpretabilidad.  
- Comparar interpretabilidad y explicabilidad.  

---

### Conocer al equipo

El equipo que trabaja en la explicabilidad de IA en este escenario está compuesto por:  

- **Olivia (ella)** – Jefa de dirección de IA  
- **Clara (ellos/ellos)** – Científica de datos  
- **Luan (ella)** – Científica de datos y validadora de modelos  

El equipo gestiona un **sistema de recomendación de productos** basado en IA para una empresa de comercio en línea.  

---

### Identificar el problema

Tras lanzar la función de recomendación de productos, los clientes plantean preguntas como:  

- “¿Por qué me recomiendan este producto?”  
- “¿Su sitio web me está espiando?”  
- “¿Hay alguna manera de desactivar las recomendaciones?”  

El equipo de ciencia de datos envía un informe a Olivia, quien reconoce la importancia de **crear sistemas de IA fiables** y decide realizar una sesión de brainstorming con su equipo.  

---

### Explicar el problema

Olivia guía al equipo en la comprensión de dos conceptos clave:  

- **Interpretabilidad:** Grado en que un observador puede entender la causa de una decisión. Permite predecir el resultado de un modelo de IA.  
- **Explicabilidad:** Va un paso más allá, mostrando **cómo el sistema ha llegado a un resultado** específico.  

---

### Escenario: Recomendaciones de productos

- Un cliente busca “proteína en polvo”.  
- El sistema recomienda tres productos basándose en compras previas: P2, P4 y P5.  

**Modelos de recomendación:**  

- **Modelo A:** Árbol de decisión sencillo. Fácil de seguir y explicar.  
- **Modelo B:** Red neuronal compleja. Difícil de rastrear y comprender el razonamiento.  

> Clara: “El Modelo A es más interpretable, mientras que el Modelo B requiere técnicas explicables para entender sus decisiones”.  
> Luan: “Las explicaciones deben adaptarse según el tipo de usuario”.  

---

### Explicaciones según el usuario

1. **Usuarios finales (clientes del sitio web):**  
   
   - Quieren entender cómo se generan las recomendaciones y cómo influir en ellas en el futuro.  
   - Ejemplo: P2 se recomienda por compras anteriores sin azúcar o con sabor a vainilla; P4 y P5 se recomiendan por búsqueda de vainilla o sin azúcar.  

2. **Aprobadores y auditores:**  
   
   - Necesitan comprender los pasos generales de decisión basados en características, para cumplir normativas y políticas internas.  

3. **Científicos de datos y validadores de modelos:**  
   
   - Quieren conocer el rendimiento general del modelo, el efecto de las características en el rendimiento y otros detalles para evaluar la calidad del modelo.  
   - Ejemplo: Cómo cambiarían las recomendaciones si se elimina la característica “Sabor”.  

---

### Abordar el problema

Gracias a la explicación de Luan, Olivia y Clara comprenden la **importancia de la explicabilidad**:  

- Todos los usuarios (clientes, auditores, aprobadores, científicos de datos) deben poder **comprender cómo y por qué** se generan las predicciones o recomendaciones.  
- El equipo está listo para implementar mejoras y asegurar la **transparencia y confianza** en el sistema de IA.  

---

### Reflexión: Explicabilidad en IA

- Piense en cómo mejorar la explicabilidad de un sistema de IA.  
- Pregúntese:  
  1. ¿Quién necesita la explicación y con qué nivel de detalle?  
  2. ¿Qué técnicas podrían emplearse para que el modelo sea comprensible para distintos usuarios?  
  3. ¿Cómo equilibrar la complejidad del modelo con la necesidad de explicabilidad?  

> Escribir estas respuestas ayuda a organizar las ideas y planificar estrategias para aumentar la confianza en los sistemas de IA.
