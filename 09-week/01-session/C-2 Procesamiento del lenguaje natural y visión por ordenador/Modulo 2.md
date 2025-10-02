## Módulo 2: La IA procesa el lenguaje natural

### Introducción

Este módulo aborda cómo el **procesamiento del lenguaje natural (NLP)** permite a la IA:

- Descomponer frases en **entidades** y sus **relaciones**.  
- Distinguir entre **emociones** y **sentimientos humanos**.  
- Resolver el significado de **expresiones ambiguas o confusas**.  

### Objetivos del módulo

Al completar este módulo, el estudiante será capaz de:

- Explicar cómo la IA extrae significado de un texto usando NLP.  
- Comprender el **problema de la clasificación** en lenguaje humano y sus soluciones.  
- Explicar la **segmentación de oraciones** y la identificación de señales (tokens) en NLP.  

---

### Procesamiento de lenguaje natural y estructura del lenguaje

- Los ordenadores trabajan mejor con **datos estructurados**, mientras que el lenguaje humano es **no estructurado**.  
- La **segmentación de oraciones** permite que los sistemas NLP analicen una oración a la vez.  
- Posteriormente, los textos se dividen en **tokens** (señales) que pueden clasificarse.  

#### Entidades y relaciones

- **Entidad:** sustantivo que representa una persona, lugar o cosa.  
- **Relación:** conjunto de entidades conectadas entre sí.  
- Ejemplo del chiste de Groucho Marx:  
  
  > “One morning, I shot an elephant in my pajamas. How he got my pajamas, I don’t know.”  
- La IA clasifica entidades y relaciones para estructurar información antes de comprenderla.  

#### Conceptos implícitos

- Un **concepto** es una idea implícita en una oración.  
- NLP debe identificar relaciones y conceptos entre oraciones conectadas.  
- Ejemplo:  
  
  > “Armen broke the glass. He always breaks the glass.”  
  - “He” se relaciona con “Armen”, mostrando cómo se manejan relaciones entre frases.  

---

### Detección de emociones y análisis de opinión

- **Detección de emociones:** identifica emociones humanas específicas (ira, felicidad, miedo).  
- **Análisis de opinión:** evalúa la intensidad o polaridad de la emoción (positivo, negativo, neutro).  
- Ejemplo: comentarios en encuestas o mensajes en redes sociales.  

---

### Problema de clasificación en NLP

- El lenguaje humano es **ambiguo**, con palabras de doble sentido.  
- Ejemplo de ambigüedad:  
  
  > “Why does your nose run and your feet smell?”  
  - "Run" puede significar “correr” o “gotear”.  
  - "Smell" puede significar “oler” o “apestar”.  
- La clasificación depende del **contexto** y requiere aprendizaje basado en patrones frecuentes.  
- Los sistemas NLP asignan un **valor de confianza** a cada clasificación, pues nunca son 100% perfectos.  

---

### Ejemplo de NLP en la práctica: IBM Watson en Jeopardy!

- Watson fue entrenado para competir en **Jeopardy!**, demostrando capacidades de NLP avanzadas.  
- El proceso:
  1. Descomponer la pregunta en partes significativas.  
  2. Analizar millones de documentos para generar posibles respuestas.  
  3. Evaluar y puntuar cada respuesta usando **miles de algoritmos**.  
  4. Clasificar respuestas según la **confianza del sistema**.  
- Resultado: Watson derrotó a los dos mejores jugadores humanos del mundo, mostrando que NLP puede:
  - Comprender el lenguaje humano.  
  - Generar respuestas precisas en contextos complejos.  
  - Aplicarse más allá de concursos, ayudando a **resolver problemas reales en industrias con grandes volúmenes de datos**.  
