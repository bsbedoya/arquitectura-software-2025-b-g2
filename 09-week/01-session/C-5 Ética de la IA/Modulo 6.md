# Módulo 6: Ética de IA

## ¿Qué es la privacidad?

### Introducción

Los datos son el **centro de la IA**. Dado que los sistemas de IA ingieren grandes cantidades de datos para aprender, identificar patrones y hacer predicciones o recomendaciones, es **fundamental salvaguardar la privacidad** de todos los datos utilizados.  

---

### Objetivos del curso

Después de completar este módulo, debería ser capaz de:  

- Identificar información personal (PI).  
- Identificar información personal confidencial (SPI).  
- Reconocer la **anonimización de modelos**.  
- Describir la **privacidad diferencial**.  
- Explicar la **minimización de datos**.  

---

### Conocer al equipo

El equipo encargado de proteger la privacidad en IA está compuesto por:  

- **Kamal (él)** – Científico de datos  
- **Nia (ella)** – Equipo de privacidad de datos  
- **Adrian (él)** – Equipo de desarrollo de IA  

---

### Identificar el problema

- Una institución educativa quiere ampliar su alcance ofreciendo cursos en línea y crear **planes de estudio personalizados** mediante IA.  
- Se desarrolla un **modelo de IA de prueba de concepto** usando datos del portal de aprendizaje.  
- Durante la revisión, surgen preocupaciones sobre la **privacidad de los datos de los usuarios**.  

> Necesidad: **proteger la privacidad mientras se desarrollan sistemas de IA**.  

---

### Explicar el problema

1. **Tipos de datos**:  
   
   - **Información personal (PI):** Relacionada con un individuo identificable, como nombre o código postal.  
   - **Información personal confidencial (SPI):** Datos que, si se ven comprometidos, podrían ser usados para dañar a una persona, como números de cuenta o fecha de nacimiento.  

2. **Riesgos en IA:**  
   
   - Los modelos de aprendizaje automático suelen requerir información personal o confidencial para entrenamiento.  
   - Si no se aplican controles de privacidad, los modelos pueden ser vulnerables a **infracciones o ataques**, como el **ataque por inferencia de pertenencia**, donde un atacante determina si un individuo formaba parte del conjunto de datos de entrenamiento.  

3. **Controles de privacidad:**  
   
   - **Durante el entrenamiento del modelo:**  
     - **Anonimización del modelo:** Oculta información personal específica en los datos de entrenamiento.  
     - **Privacidad diferencial:** Introduce ruido controlado para proteger la identidad de los individuos en el conjunto de datos.  
   - **Después del entrenamiento del modelo:**  
     - **Minimización de datos:** Limita el almacenamiento y uso de datos solo a lo estrictamente necesario.  

> Aplicando estos controles, los modelos se refuerzan contra vulneraciones y se protege la privacidad de los usuarios.  

---

### Abordar el problema

- El equipo comprende que **información aparentemente inofensiva puede identificar a personas**.  
- Proteger tanto los **datos personales** como el **modelo de IA** es fundamental para:  
  - Mantener la confianza de los usuarios  
  - Cumplir con regulaciones de privacidad  
  - Garantizar el uso responsable de la IA  
