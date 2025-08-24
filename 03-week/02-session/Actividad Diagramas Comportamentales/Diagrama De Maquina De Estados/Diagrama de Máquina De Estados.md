# Documentación: Diagrama de Máquina de Estados (UML)

## Concepto General
El **diagrama de máquina de estados** en UML (Unified Modeling Language) es un diagrama **comportamental** que modela el **ciclo de vida de un objeto** en un sistema. Representa los diferentes **estados** en los que puede encontrarse un objeto, así como los **eventos, condiciones y transiciones** que provocan cambios entre estos estados.  
Es ampliamente utilizado para modelar **sistemas reactivos**, aquellos que responden a estímulos o eventos externos.

---

## Elementos Principales

1. **Estado (State)**  
   - Representa una condición o situación en la vida de un objeto.  
   - Se dibuja como un rectángulo con esquinas redondeadas.  
   - Ejemplo: `Activo`, `Inactivo`, `En espera`.

2. **Estado inicial (Initial State)**  
   - Punto de inicio del ciclo de vida de un objeto.  
   - Se representa como un **círculo sólido negro**.

3. **Estado final (Final State)**  
   - Representa la terminación del ciclo de vida de un objeto.  
   - Se dibuja como un **círculo negro con borde**.

4. **Transición (Transition)**  
   - Es el **cambio de un estado a otro** provocado por un evento.  
   - Se dibuja como una **flecha**.  
   - Puede estar acompañada de:  
     - **Evento**: desencadenante (`evento()`).
     - **Condición de guarda**: se cumple para que ocurra la transición (`[condición]`).
     - **Acción**: lo que sucede en el cambio (`/ acción`).

5. **Subestados y estados compuestos**  
   - Los **estados compuestos** son aquellos que contienen otros estados internos, usados para simplificar diagramas complejos.

6. **Historial (History State)**  
   - Representa la capacidad de un estado compuesto de recordar en qué subestado estaba antes de salir de él.  
   - Se dibuja como una **H** dentro de un círculo.

7. **Acciones en estados**  
   - **entry / acción** → lo que ocurre al entrar al estado.  
   - **exit / acción** → lo que ocurre al salir del estado.  
   - **do / acción** → acción que ocurre mientras se está en el estado.

---

## Ventajas del Diagrama de Máquina de Estados
- Modela con precisión el **ciclo de vida** de objetos y componentes.  
- Útil para sistemas donde las **respuestas dependen de eventos externos**.  
- Ayuda a **identificar casos límite y errores de transición**.  
- Es aplicable en **sistemas embebidos, controladores, interfaces y flujos de interacción**.

---

## Ejemplo Práctico: Cajero Automático (ATM)

```plantuml
@startuml
[*] --> Inactivo

Inactivo --> Autenticando : insertarTarjeta()
Autenticando --> SeleccionOperacion : validarPIN() / mostrarMenu()
Autenticando --> Inactivo : errorPIN() / expulsarTarjeta()

SeleccionOperacion --> Procesando : elegirOperacion()
Procesando --> ExpulsarTarjeta : operacionCompletada()
ExpulsarTarjeta --> Inactivo : tarjetaExpulsada()

ExpulsarTarjeta --> [*]
@enduml
