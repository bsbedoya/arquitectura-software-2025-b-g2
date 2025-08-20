# Diagrama de Objetos UML

## Concepto

El **diagrama de objetos** en UML es un tipo de **diagrama estructural** que muestra una instancia concreta de los elementos definidos en un **diagrama de clases**. Representa **objetos reales** (instancias de clases) y sus **relaciones en un momento específico del tiempo**, es decir, cómo se utilizan las clases en un caso particular del sistema.

Se centra en lo siguiente:

- Los **objetos** (instancias de clases) con sus atributos y valores actuales.
- Las **asociaciones** existentes entre esos objetos.
- El **estado del sistema** en un instante determinado.

## Elementos principales

1. **Objeto**  
   
   - Se representa con un rectángulo subrayado.  
   - Notación: `nombreObjeto:NombreClase`  
   - Ejemplo: `carrito1:CarritoCompra`

2. **Atributos del objeto**  
   
   - Se muestran dentro del rectángulo, con valores asignados.  
   - Ejemplo:  
     
     ```
     cliente1:Cliente
     nombre = "Juan Pérez"
     email = "juanp@gmail.com"
     ```

3. **Líneas de asociación**  
   
   - Conectan los objetos que interactúan entre sí.  
   - Representan enlaces entre instancias en lugar de asociaciones abstractas.

4. **Multiplicidad**  
   
   - Puede indicarse en los enlaces para mostrar cuántas instancias participan en la relación.

## Diferencias con el Diagrama de Clases

- **Diagrama de Clases**: muestra las estructuras **genéricas** del sistema (clases, atributos, métodos, relaciones).
- **Diagrama de Objetos**: muestra **ejemplos concretos** en un momento específico, útil para ilustrar cómo se comporta el sistema en la práctica.

## Usos principales

- Verificar la corrección del diagrama de clases.
- Mostrar ejemplos de cómo se instancian las clases.
- Explicar escenarios de uso con ejemplos concretos.
- Complementar pruebas y documentación del sistema.