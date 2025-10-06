## Módulo 1: Introducción al Aprendizaje de Máquinas

### Objetivos del Módulo
Al completar este módulo, el estudiante será capaz de:
- Distinguir entre **inteligencia artificial (IA)**, **aprendizaje automático (machine learning)** y **aprendizaje profundo (deep learning)**.
- Describir los tres tipos principales de aprendizaje de máquinas: **supervisado**, **no supervisado** y **por refuerzo**.

---

### Conceptos Clave

#### Inteligencia Artificial, Aprendizaje Automático y Aprendizaje Profundo
- **Inteligencia Artificial (IA):** Campo amplio que busca crear sistemas capaces de realizar tareas que normalmente requieren inteligencia humana.
- **Aprendizaje Automático (Machine Learning):** Subconjunto de la IA que permite a las máquinas aprender de datos y mejorar su rendimiento sin ser programadas explícitamente.
- **Aprendizaje Profundo (Deep Learning):** Subconjunto del aprendizaje automático que utiliza redes neuronales profundas para aprender de grandes cantidades de datos, especialmente en tareas complejas como reconocimiento de imágenes y procesamiento de lenguaje natural.

---

### Formas Generales de Aprender de los Datos
Los sistemas de IA utilizan **algoritmos** para predecir y clasificar datos, y pueden aprender de tres formas generales:

1. **Aprendizaje Supervisado**
   - La máquina recibe **datos estructurados y etiquetados**.
   - Detecta **patrones** en los datos para hacer **predicciones futuras**.
   - Ejemplo: 
     - Datos de temperaturas diarias en una tabla para predecir temperaturas futuras.
     - Imágenes etiquetadas de flores para que la máquina aprenda a identificar rosas.
   - La salida incluye un **valor de confianza** que indica la probabilidad de acierto (p. ej., 85 % de certeza).
   - Más datos → mayor precisión.

2. **Aprendizaje No Supervisado**
   - La máquina recibe **datos no etiquetados** y debe **estructurarlos y encontrar patrones por sí misma**.
   - Ejemplo: 
     - Analizar textos de un libro o artículos sobre plantas para identificar relaciones y atributos sin etiquetas previas.
   - Útil para descubrir **estructuras ocultas** y hacer **clasificaciones automáticas**.
   - El sistema asigna un **valor de confianza** a sus propias conclusiones.

3. **Aprendizaje por Refuerzo**
   - La máquina aprende mediante **ensayo y error**.
   - Las acciones correctas reciben **recompensas**, y las incorrectas, **penalizaciones**.
   - Ejemplo: 
     - Identificar si una imagen muestra un perro. Las respuestas correctas ajustan positivamente los algoritmos; las incorrectas, negativamente.
   - Cada decisión se basa en la **decisión anterior**.
   - Muy útil para problemas donde se necesita **alcanzar un objetivo a largo plazo**, como juegos o estrategias secuenciales.

---

### Cómo Aprenden las Máquinas
- Se analizan dos tecnologías principales:
  1. **Aprendizaje Automático Clásico:** Algoritmos tradicionales de machine learning que trabajan con datos estructurados y menos complejos.
  2. **Ecosistema de Aprendizaje Profundo:** Tecnologías más avanzadas basadas en redes neuronales, capaces de manejar grandes cantidades de datos no estructurados.
- Analizar estas tecnologías es como **mirar bajo el capó de dos automóviles** diferentes para entender cómo funcionan sus "motores" internos.

---

### Resumen Visual del Aprendizaje de Máquinas
| Tipo de Aprendizaje | Entrada de Datos | Método de Aprendizaje | Salida / Resultado |
|--------------------|----------------|---------------------|-----------------|
| Supervisado         | Datos estructurados y etiquetados | Detecta patrones | Predicciones con valor de confianza |
| No Supervisado      | Datos no etiquetados | Encuentra estructura y relaciones | Clasificación o agrupamiento con valor de confianza |
| Por Refuerzo        | Información mínima sobre el entorno | Ensayo y error con recompensas/penalizaciones | Decisiones optimizadas para alcanzar un objetivo |

---

### Notas Adicionales
- La máquina **nunca predice con certeza absoluta**, siempre asigna un valor de confianza.
- **Más datos y entrenamiento** mejoran la precisión del sistema.
- La elección entre aprendizaje supervisado, no supervisado o por refuerzo depende del **tipo de problema y datos disponibles**.