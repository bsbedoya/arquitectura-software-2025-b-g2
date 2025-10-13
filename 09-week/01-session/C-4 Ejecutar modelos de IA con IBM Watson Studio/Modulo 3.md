# Módulo 3: Desarrollar su proyecto de aprendizaje automático

## Submódulo 3: Introducción

En este módulo practicará cómo crear un modelo de IA, cómo entrenarlo con datos de bancos alemanes y cómo ejecutarlo para que genere predicciones.

### Objetivos del curso

Después de completar este submódulo, debería ser capaz de:

- Crear un modelo de IA utilizando AutoAI en IBM Watson Studio
- Ejecutar un experimento de predicción para un modelo de IA
- Explicar la matriz de confusión

### Simulación: Configurar y ejecutar su proyecto

Se toma un breve descanso y charla con miembros del departamento informático del banco alemán. Hasta ahora, admiran tu trabajo. Ahora vuelve al proyecto. ¡Es hora de empezar a ejecutarlo!

Hasta ahora lo ha hecho bien, y la dirección del banco está deseando ver lo que puede conseguir. Ahora empieza la parte más importante de su proyecto: configurar cuatro algoritmos y ponerlos a competir. Le permitirá recomendar el mejor sistema a su cliente, ahorrándole potencialmente millones de euros en préstamos impagados. Sin IBM Watson Studio, esto podría suponer mucho trabajo. Pero el entorno de desarrollo integrado de Watson Studio le ayudará a avanzar en el proceso.

En esta simulación, utilizará IBM Watson Studio para configurar su modelo de IA y luego ejecutará los algoritmos de forma competitiva para ver cuál predice con mayor precisión los préstamos de riesgo. Hay 31 pasos que seguir.

### Comprender la matriz de confusión

El servicio AutoAI ha generado un proyecto ganador. Cuando ha seleccionado este proyecto, ha podido obtener información sobre el mismo, incluida la matriz de confusión. La matriz de confusión le permite calcular el rendimiento de su modelo.

En el gráfico se puede ver que el modelo tiene una precisión de 0,784 en los datos de prueba. Eso significa:

- Alrededor del 78 % de las veces, el modelo predice con exactitud si alguien es un riesgo bueno para un préstamo.
- Aproximadamente el 22 % de las veces, el modelo hace una predicción inexacta.

#### Términos

Para comprender la matriz de confusión, deberá entender algunos términos.  

Gráfico de una matriz de confusión: una tabla de cuatro celdas con:

- Columna 1: Sin riesgo (previsto)  
- Columna 2: Riesgo (previsto)  
- Fila 1: Sin riesgo (observado)  
- Fila 2: Riesgo (observado)  

Celdas:

- R1C1: VERDADERO POSITIVO  
- R1C2: FALSO POSITIVO  
- R2C1: FALSO NEGATIVO  
- R2C2: VERDADERO NEGATIVO

### Resultados

Ahora que entiende la estructura de la matriz de confusión, vamos a examinar la generada por su experimento:

- Sin riesgo predicho / Sin riesgo observado: 296  
- Riesgo predicho / Sin riesgo observado: 37  
- Sin riesgo predicho / Riesgo observado: 71  
- Riesgo predicho / Riesgo observado: 96  

Porcentaje total correcto del modelo: 78.4 %  

- Columna Sin riesgo prevista: 80,6 %  
- Columna Riesgo pronosticado: 72,2 %  
- Fila Sin riesgo observado: 88,9 %  
- Fila Riesgo observado: 57,5 %

IBM Watson codifica por colores la matriz de confusión: verde indica más correcto y azul menos correcto.  

Conclusiones:

- Las celdas Verdadero positivo y Verdadero negativo muestran que el modelo es muy bueno para predecir Sin riesgo y Riesgo correctamente.  
- El modelo no es tan bueno con falsos positivos y falsos negativos.  
- Con más entrenamiento y un conjunto de datos mayor, el modelo debería mejorar su precisión.

### Resumen de pasos realizados

Ha creado un modelo de IA con cuatro algoritmos, los ha entrenado con el conjunto de datos del banco y ha ejecutado cada algoritmo para obtener predicciones sobre los solicitantes con mayor probabilidad de impago:

1. Ha utilizado AutoAI para construir sus modelos.  
2. Ha creado un experimento, adjuntando y configurando sus datos.  
3. Ha seleccionado una columna predecible en el conjunto de datos como eje X.  
4. Ha ejecutado el experimento.  
5. Ha mostrado la matriz de confusión, lista para aplicarla en la última parte de su proceso.
