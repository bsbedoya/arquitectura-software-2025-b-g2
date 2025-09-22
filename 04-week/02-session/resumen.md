POO: Paradigma de programación que modela el entorno mediante objetos que combinan estado y comportamiento. Permite representar conceptos del dominio de forma reutilizable y modular. Sus pilares son:

- **Herencia**: Permite que una clase (subclase) herede atributos y métodos de otra clase (superclase), promoviendo la reutilización de código. Ejemplo: Una clase `Usuario` que hereda de una clase `Persona`, reutilizando atributos como `nombre` y `documento`.

- **Polimorfismo**: Habilidad de un objeto para tomar diferentes formas, permitiendo que un mismo método se comporte de manera distinta según el contexto. Ejemplo: Un método `autenticar()` que puede implementarse de manera diferente en clases `Usuario` y `Administrador`.

- **Encapsulamiento**: Restricción del acceso directo a los atributos de un objeto, proporcionando métodos para acceder y modificar su estado. Ejemplo: Usar métodos `getSaldo()` y `setSaldo()` en una clase `Cuenta` para proteger el atributo `saldo`.

- **Abstracción**: Proceso de ocultar los detalles de implementación y mostrar solo las funcionalidades esenciales. Ejemplo: Una clase abstracta `Figura` con un método `calcularArea()` que es implementado por clases concretas como `Circulo` y `Rectangulo`.

```plantuml
@startuml

@enduml
```


UML: Lenguaje de modelado visual estándar para especificar, visualizar, construir y documentar los artefactos de un sistema software. Se usa para representar diferentes vistas del sistema (estructura, comportamiento, interacciones).

Diagrama de clases: representa la estructura estática (clases, atributos, métodos y relaciones).

Diagrama de Casos de Uso: muestra las funcionalidades del sistema desde la perspectiva de los actores.

Diagrama de Dominio: describe los conceptos principales del dominio y sus relaciones sin detalles de diseño.

Diagrama de Paquetes: 

Diagrama de Secuencia: