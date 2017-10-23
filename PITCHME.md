# Boxes Versus Lego Bricks in Go Programming

![Box](assets/box.png)
<div>Icon made by <a href="http://www.freepik.com" title="Freepik">Freepik</a> from <a href="https://www.flaticon.com/" title="Flaticon">www.flaticon.com</a> is licensed by <a href="http://creativecommons.org/licenses/by/3.0/" title="Creative Commons BY 3.0" target="_blank">CC 3.0 BY</a></div>

versus

![Lego brick](assets/lego.png)
<div>Icon made by <a href="https://www.flaticon.com/authors/coucou" title="Coucou">Coucou</a> from <a href="https://www.flaticon.com/" title="Flaticon">www.flaticon.com</a> is licensed by <a href="http://creativecommons.org/licenses/by/3.0/" title="Creative Commons BY 3.0" target="_blank">CC 3.0 BY</a></div>

---

## Boxes are great

![Box](assets/box.png)
<div>Icon made by <a href="http://www.freepik.com" title="Freepik">Freepik</a> from <a href="https://www.flaticon.com/" title="Flaticon">www.flaticon.com</a> is licensed by <a href="http://creativecommons.org/licenses/by/3.0/" title="Creative Commons BY 3.0" target="_blank">CC 3.0 BY</a></div>

- They have a simple interface |
- They hide details |
- They can be used everywhere |

---

## Boxes are awful


---

## Procedural Programming
![Logo](assets/call-graph.png)

Note:
A call decouples from
- registers and
- variables and state in the calling procedure.

A procedure can be called from and reused anywhere.

This is still used a lot esp. for utility functions.

---

## Object Orientation
![Logo](assets/class-diagram.jpg)

Note:
- A class encapsulates functionality and state.
- Thus it decouples from other classes.
- Inheritence creates a lot of coupling.
- OOP works best for purely technical domains.
- It grew strong together with GUIs.

---

## Clean Code
![Logo](assets/clean-code.jpg)

Note:
- Favour Composition over Inheritance
- Separation of Concerns
- Information Hiding Principle
- Tell, don´t ask! (High cohesion)
- Dependency injection
- SOLID principles!!!
- Single Responsibility Principle
- Open Closed Principle: software entities should be open for extension, but closed for modification
- Liskov Substitution Principle
- Interface Segregation Principle: many client-specific interfaces are better than one general-purpose interface
- Dependency Inversion Principle: one should depend upon abstractions, not concretions

---

## Messaging
![Logo](assets/messaging.jpg)

Note:
- Why don't we use HTTP internaly?
- Asynchronism prevents time wise coupling (no timeout handling).
- Persistence prevents stability wise coupling (no exponential backoff that has to be randomised, no circuit breaker).
- Sender and receiver are decoupled with messaging.

---

## Functional Programming

---

## The Goal

- Decoupling as a natural thing to happen |
- Good match for business logic           |
- Easy to understand                      |
- Natural graphical representation        |

