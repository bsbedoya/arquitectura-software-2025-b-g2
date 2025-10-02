## Módulo 3: NLP convierte señales en significado  

### Introducción
Este módulo muestra cómo las aplicaciones de **gestión de información no estructurada (UIMA)** permiten a un sistema:
- Descomponer preguntas formuladas por humanos.  
- Generar múltiples respuestas posibles.  
- Identificar la respuesta con mayor probabilidad de ser correcta.  

### Objetivos del módulo
Al completar este módulo, el estudiante será capaz de:
- Describir cómo un **chatbot entiende, razona, aprende e interactúa** con los usuarios.  
- Distinguir entre **intenciones, entidades y diálogos**.  
- Identificar los usos adecuados de los chatbots.  
- Reconocer los usos reales del **procesamiento del lenguaje natural (NLP)**.  

---

### Estructura de un chatbot
- Los chatbots responden eficazmente a preguntas claras relacionadas con el propósito del sitio web.  
- Para preguntas ambiguas o irrelevantes, devuelven mensajes de no comprensión.  
- Son útiles en campos como **comercio minorista y medicina de atención inmediata**.  
- Operan con **datos limitados**, pero de manera eficiente, cubriendo un gran número de consultas con pocas respuestas.  

#### Frontend y Backend
- **Frontend:** canal de comunicación con el usuario (lectura y presentación de mensajes).  
- **Backend:** lógica del chatbot, memoria de conversaciones y procesamiento de consultas.  

#### Clasificadores
- Correlacionan múltiples formas de preguntar con un conjunto reducido de respuestas.  
- Permiten a un chatbot gestionar cientos de preguntas con pocas respuestas posibles.  
- Las preguntas no reconocidas se envían a **representantes humanos**.  

---

### Intenciones, entidades y diálogo
El backend del chatbot se compone de tres elementos clave:  

#### Intención
- Es el **propósito del usuario** al interactuar con el chatbot (como un verbo o acción).  
- Ejemplos: presentar una queja, preguntar una dirección, consultar horarios.  

#### Entidad
- Sustantivos que representan **personas, lugares u objetos**.  
- Ejemplo:  
  > Pregunta: “¿Cuál es el horario del restaurante de Austin?”  
  - Intención: proporcionar horario  
  - Entidad: Austin  

#### Diálogo
- Diagrama de flujo tipo **IF / THEN** que determina la respuesta del chatbot.  
- Condensa cada momento de la conversación en **nodos**, cada uno con posibles respuestas.  
- Permite manejar entradas humanas ambiguas o mal redactadas gracias a NLP.  

---

### Ejemplo de análisis NLP
- Los ordenadores convencionales necesitarían miles de líneas IF / THEN para cada pregunta posible.  
- Combinando **NLP con intenciones, entidades y diálogos**, la IA:
  - Analiza los componentes de la frase.  
  - Clasifica las posibles respuestas según puntuación de evidencia.  
  - Devuelve la respuesta con mayor confianza.  

#### Caso práctico: Staples
- El chatbot recibe consultas de clientes y el **backend ejecuta NLP** para comprender la intención.  
- Consulta servicios cognitivos y el historial de compras.  
- Permite comprar productos o realizar seguimiento de pedidos de manera eficiente.  
- Resultado: soporte 24/7, atención efectiva y liberación de recursos humanos.  
