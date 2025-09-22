# Quiz: Modelado de Entidades y Relaciones en Paquetes UML

## Instrucciones del Reto
El estudiante debe **asignar de manera adecuada cada entidad al paquete que corresponde (A, B, C o D)** según su naturaleza.  
Además, debe **definir las relaciones apropiadas** entre las entidades utilizando los conceptos de:  
- **Generalización**  
- **Agregación**  
- **Composición**  
- **Implementación**  

---

## Posibles Paquetes
- **Paquete A:** module_security  
- **Paquete B:** geolocation  
- **Paquete C:** financial_system  
- **Paquete D:** transportation  

---

## Entidades
- **continent** {id, name, code}  
- **user** {id, username, password_hash, email, is_active}  
- **department** {id, name, code, country_id}  
- **role** {id, name, description}  
- **city** {id, name, postal_code, department_id}  
- **person** {id, first_name, last_name, document_type, document_number, birth_date}  
- **country** {id, name, iso_code, continent_id}  
- **permission** {id, code, description}  
