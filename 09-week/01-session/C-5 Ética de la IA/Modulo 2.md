# Módulo 5: Ética de IA

## ¿Qué es la imparcialidad?

### Introducción

En IA, **imparcialidad** es el tratamiento equitativo de individuos o grupos de individuos.  
La imparcialidad se consigue cuando se **mitigan los sesgos no deseados**.  

En IA, el **sesgo** es un error sistemático que se ha diseñado, intencionadamente o no, de forma que pueda generar decisiones injustas.  
El sesgo puede encontrarse en el sistema de IA, en los datos utilizados para entrenar y probar el sistema, o incluso en ambos.  

El sesgo puede surgir en un sistema de IA debido a **expectativas culturales, limitaciones técnicas o contextos de implantación imprevistos**.  

---

### Objetivos del curso

Después de completar este módulo, debería ser capaz de:  

- Describir la imparcialidad en IA.  
- Describir los atributos protegidos.  
- Identificar grupos privilegiados y grupos no privilegiados.  
- Explicar el sesgo de la IA.  

---

### Conocer al equipo

En la siguiente historia, **Jordan**, **Priscilla** y **Nischal**, empleados de una gran entidad bancaria nacional, le hablarán de imparcialidad y parcialidad al descubrir un problema en un sistema de IA.  

La entidad se está preparando para implantar un sistema de IA que les ayude a **identificar a los candidatos de mayor valor** en su bolsa de ascensos.  
Priscilla se da cuenta de que la mayoría de los candidatos pertenecen a una misma raza, y comienza la revisión del sistema para entender qué está provocando el resultado.  

Descubren que el sistema está **sesgado**.  
Esta historia aborda conceptos básicos sobre el sesgo en el contexto de la imparcialidad y cómo puede introducirse en un sistema.  

---

### El equipo

- **Jordan (ellos/ellas)** – Responsable de datos (CDO)  
- **Priscilla (ella)** – Directora de operaciones de personal (PeopleOps)  
- **Nischal (él)** – Director de ciencia de datos  

---

### Identificar el problema

Jordan acaba de asumir su nuevo rol como responsable de datos del banco.  
Antes de su llegada, la empresa había estado probando un nuevo sistema de IA para ayudar a identificar candidatos de alto valor en el grupo de ascensos.  

El sistema parecía funcionar bien, por lo que el equipo compartió la lista de ascensos con Priscilla.  
Al revisar la lista y los datos demográficos, **Priscilla nota que la mayoría de los candidatos promovidos son blancos**.  

Sorprendida, recuerda que había recomendado a un empleado no blanco altamente calificado para un ascenso, pero no aparece en la lista.  
Decide llamar a Jordan para comenzar una investigación.  

> 💬 **Priscilla:** “Jordan, ¿podrían tú y el equipo de ciencia de datos investigar esto?”  
> **Jordan:** “¡Sí, iré a reunir al equipo!”  

---

### Análisis de datos

El equipo de ciencia de datos revisa los **datos de los últimos 5 años** usados para entrenar el sistema.  
Una muestra de los datos incluye atributos como:  

| Atributo            | Ejemplo 1  | Ejemplo 2  |
| ------------------- | ---------- | ---------- |
| ID de empleado      | 1          | 2          |
| Departamento        | Datos e IA | Datos e IA |
| Cualificación       | Máster     | Doctorado  |
| Raza                | Blanca     | No blanca  |
| Años de antigüedad  | 10         | 12         |
| Valoración media    | ...        | ...        |
| Decisión de ascenso | Sí         | No         |

Un gráfico muestra que **555 candidatos blancos** y **85 candidatos no blancos** fueron promovidos.  

Jordan observa que históricamente ha habido **una desproporción en los ascensos** según la raza.  
Concluye que el sistema podría tener un **problema de imparcialidad**.  

---

### Reunión del equipo

Jordan lidera la investigación y pregunta al grupo si vale la pena mantener el sistema de IA.  

> **Jordan:** “¿Creéis que deberíamos seguir utilizando la IA para ayudar en el proceso de ascensos?”  
> **Priscilla:** “Si se hace bien y de forma que podamos confiar, nos ahorrará tiempo y podría ayudarnos a que el proceso sea más justo.”  

El equipo coincide en que la IA tiene valor si se aplica correctamente.  

---

### Explicar el problema

Jordan pide a Nischal que prepare material para explicar al equipo de PeopleOps lo que podría haber causado el problema.  

Al día siguiente, **Nischal** inicia la reunión hablando del **sesgo**.  

> **Jordan:** “¿Puedes explicar el significado de sesgo en el contexto de la imparcialidad?”  
> **Nischal:** “El sesgo, en general, es un error sistemático. En el contexto de la imparcialidad, el problema gira en torno al sesgo no deseado, que da a algunos grupos una ventaja sistemática y a otros una desventaja sistemática.”  

Nischal continúa explicando nuevos conceptos:  

- Dividir la población en grupos permite **mitigar la disparidad de resultados**.  
- El atributo que separa a los grupos se denomina **atributo protegido**.  
- Ejemplos de atributos protegidos: **raza, edad, sexo, identidad de género, origen étnico.**  

>  *Nota:* Por motivos legales o políticos, los atributos protegidos no siempre se incluyen en los datos. En esos casos, se pueden **imputar** o estimar mediante **aprendizaje por transferencia**.  

---

### Grupos privilegiados y no privilegiados

> “Tradicionalmente, si un grupo recibe resultados más favorables que otro, puede existir sesgo.  
> El grupo que recibe resultados más favorables es el **grupo privilegiado**.  
> El grupo que recibe menos o ningún resultado favorable es el **grupo no privilegiado**.”  

El objetivo de la imparcialidad en IA es **minimizar el impacto de los sesgos no deseados**.  

Nischal muestra un gráfico con los datos de ascensos por raza y señala:  

> “En nuestros datos, la raza es un posible atributo protegido.  
> Los candidatos blancos tienen un porcentaje de ascensos mucho mayor que los no blancos.”  

Jordan concluye:  

- Los empleados blancos pertenecen al **grupo privilegiado**.  
- Los empleados no blancos pertenecen al **grupo no privilegiado**.  
- La raza, cuando se utiliza para dividir la población, actúa como **atributo protegido**.  

---

### Abordar el problema

El equipo determina que el modelo de IA presenta un **problema de imparcialidad**.  
Para solucionarlo, deben **reducir los sesgos no deseados**.  

> En IA, el sesgo es un error sistemático que puede generar decisiones injustas.  
> Este sesgo puede encontrarse en los datos o en el propio sistema.  

Según los datos analizados, los sesgos introducidos afectaban directamente los resultados de ascensos.  

> **Conclusión:**  
> El equipo de ciencia de datos debe diseñar un método para comprobar los sesgos a lo largo del ciclo de vida del sistema de IA y aplicar medidas de mitigación.  
> **Abordar los problemas de imparcialidad es un deporte de equipo.**
