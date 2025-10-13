# Módulo 3

## El ecosistema del aprendizaje profundo

### Introducción

Las redes neuronales se inspiran en las células del cerebro humano. En este módulo, aprenderá cómo las redes neuronales pueden mejorar enormemente la capacidad de un sistema de IA para analizar datos complejos y mejorar sus resultados mediante el método de ensayo y error.

### Objetivos del curso

Después de completar este módulo, debería ser capaz de:

- Describir la forma en que las redes neuronales se inspiran en el cerebro humano.  
- Rastrear el flujo de información a través de los nodos de un perceptrón.  
- Describir el proceso de aprendizaje de ensayo y error del aprendizaje automático.  
- Definir y describir el aprendizaje profundo y su ecosistema.  

### Inspirado en el cerebro humano

El aprendizaje automático ha evolucionado hasta convertirse en una colección de potentes aplicaciones denominadas **ecosistema del aprendizaje profundo**. La base de muchas aplicaciones es la **red neuronal**, inspirada en cómo se comunican las neuronas en el cerebro humano.

- Las **neuronas humanas** tienen un cuerpo celular, un axón y terminales ramificadas que transmiten señales.  
- Un cerebro humano tiene aproximadamente 100.000 millones de neuronas, cada una conectada a unas 10.000 neuronas más.

### Perceptrón

Un **perceptrón** es el equivalente de una sola neurona en una red neuronal artificial:

- Capa de entrada: recibe señales.  
- Capas ocultas: ejecutan algoritmos y ajustan resultados.  
- Capa de salida: genera la respuesta final.  

Cada nodo de una capa utiliza una función de activación (como la **función sigmoidea**) para decidir si se “activa” o no, similar a superar un umbral en una neurona.

### Una ruta por una red neuronal

Ejemplo práctico: Marc decide si pedir pizza.

1. **Pedido con antelación:** nodo asigna valor 1, ponderación 3 → salida = 3  
2. **Pedir ensalada:** nodo asigna valor 0,5, ponderación 0,5 → salida = 0,25  
3. **Cupón gratis:** nodo asigna valor 1, ponderación 5 → salida = 5  

**Suma de salidas:** 3 + 0,25 + 5 = 8,25  
**Umbral de activación:** 6 → decisión: ¡hoy se cena pizza!

### Aprendizaje por ensayo y error

Las redes neuronales aprenden ajustando continuamente sus algoritmos:

- Comparan resultados con el corpus de datos almacenados.  
- Modifican cálculos si no coinciden con patrones previos.  
- Realizan miles de ajustes para mejorar precisión paso a paso.  
- Emiten un valor de **confianza** junto con la predicción, dejando la decisión final a un humano si es necesario.

### De los perceptrones al aprendizaje profundo

Los sistemas avanzados de IA utilizan múltiples capas ocultas, formando **redes neuronales profundas (DNN)**:

- Mayor capacidad de procesamiento y aprendizaje.  
- Las DNN pueden duplicarse en equipos competidores que aprenden de los errores de otros (aprendizaje de refuerzo).  

### Aplicaciones de las DNN

1. **Identificación de fotos históricas:** comparan imágenes con millones en el corpus y proporcionan nombres y ubicaciones.  
2. **Predicción de precios de vivienda:** ayudan a inmobiliarias a invertir eficientemente.  
3. **Vehículos autónomos:** modelan millones de situaciones de conducción.  
4. **Diagnóstico médico:** detectan variaciones en resonancias magnéticas y alertan sobre posibles cánceres tratables.  

El ecosistema de aprendizaje profundo influye ampliamente en la vida diaria, ofreciendo soluciones complejas a problemas que serían imposibles de manejar con perceptrones simples.
