# Decoupling and Code
<h2 class="fragment">Or: Decoupling Through the Ages</h2>
<h2 class="fragment">Or: Decoupling. What Is It Good For?</h2>

Note:
Dear community, as we all know:

In the beginning there was chaos and then came ...

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

