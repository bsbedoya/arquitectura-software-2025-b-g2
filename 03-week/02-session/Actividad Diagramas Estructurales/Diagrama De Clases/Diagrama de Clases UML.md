# Diagrama de Clases UML

## Concepto

Un **diagrama de clases** es un diagrama estructural de UML que describe la **estructura estática** de un sistema, mostrando:

- Las **clases** que lo componen.
- Sus **atributos** y **métodos**.
- Las **relaciones** que existen entre ellas.

Es una de las herramientas principales para modelar sistemas orientados a objetos.

---

## Elementos de un Diagrama de Clases

### 1. Clases

Representan las entidades principales del sistema.

- Se muestran como un rectángulo dividido en tres secciones:
  - **Nombre de la clase**
  - **Atributos**
  - **Métodos**

Ejemplo (en notación textual simplificada):
+------------------+
|   Carrito        |
+------------------+
| -id: int         |
| -total: double   |
+------------------+
| +agregar()       |
| +eliminar()      |
+------------------+

### 2. Atributos

- Representan las propiedades o datos de la clase.  
- Notación: `visibilidad nombre: tipo`.

### 3. Métodos

- Son las operaciones o funciones que la clase puede realizar.  
- Notación: `visibilidad nombre(parámetros): tipo`.

### 4. Visibilidad

- `+` Público  
- `-` Privado  
- `#` Protegido  

---

## Relaciones entre Clases

### Asociación

Representa un vínculo simple entre clases.  
Notación: una línea sólida.
Usuario -------- Carrito

### Multiplicidad

Indica la cantidad de instancias relacionadas.

- `1` → uno
- `0..1` → cero o uno
- `*` → muchos
- `1..*` → uno o muchos

Ejemplo:
Usuario 1 -------- * Carrito

### Agregación

Representa una relación "todo-parte" débil (la parte puede existir sin el todo).  
Notación: rombo blanco.
Carrito o---- ItemCarrito

### Composición

Representa una relación "todo-parte" fuerte (la parte no puede existir sin el todo).  
Notación: rombo negro.
Carrito *---- ItemCarrito

### Herencia (Generalización)

Una clase hija hereda atributos y métodos de una clase padre.  
Notación: flecha con triángulo vacío.
Pago <|--- PagoTarjeta

### Realización

Se usa entre interfaces y clases que las implementan.  
Notación: línea discontinua con triángulo vacío.
InterfazPago <|.. PagoTarjeta

### Dependencia

Una clase usa temporalmente otra clase.  
Notación: línea discontinua con flecha.
Carrito ..> Producto

---

## Utilidad del Diagrama de Clases

- Modelar la estructura del sistema antes de programar.
- Mostrar claramente cómo se relacionan los objetos.
- Servir como base para el diseño orientado a objetos.
- Facilitar la comunicación entre analistas, diseñadores y programadores.
