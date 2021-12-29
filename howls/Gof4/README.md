# Gof4 - Gang of Four Design Patterns


> The elements of this language are entities called patterns. Each pattern describes a problem that occurs over and over again in our environment, and then describes the core of the solution to that problem, in such a way that you can use this solution a million times over, without ever doing it the same way twice.
> <p>— Christopher Alexander, A Pattern Language



module to test the implementation of some Gof4 design patterns in GO


<hr>

### Creational Design Patterns

- __Abstract Factory:__ Allows the creation of objects without specifying their concrete type.
- __Builder:__ Uses to create complex objects.
- __Factory Method:__ Creates objects without specifying the exact class to create.
- __Prototype:__ Creates a new object from an existing object.
- __Singleton:__ Ensures only one instance of an object is created.

<hr>

### Structural Design Patterns

- __Adapter__ Allows for two incompatible classes to work together by wrapping an interface around one of the existing classes.
- __Bridge__ Decouples an abstraction so two classes can vary independently.
- __Composite__ Takes a group of objects into a single object.
- __Decorator__ Allows for an object’s behavior to be extended dynamically at run time.
- __Facade__ Provides a simple interface to a more complex underlying object.
- __Flyweight__ Reduces the cost of complex object models.
- __Proxy__ Provides a placeholder interface to an underlying object to control access, reduce cost, or reduce complexity.

<hr>

### Behavior Design Patterns

- __Chain of Responsibility:__ Delegates commands to a chain of processing objects.
- __Command:__ Creates objects which encapsulate actions and parameters.
- __Interpreter:__ Implements a specialized language.
- __Iterator:__ Accesses the elements of an object sequentially without exposing its underlying representation.
- __Mediator:__ Allows loose coupling between classes by being the only class that has detailed knowledge of their methods.
- __Memento:__ Provides the ability to restore an object to its previous state.
- __Observer:__ Is a publish/subscribe pattern which allows a number of observer objects to see an event.
- __State:__ Allows an object to alter its behavior when its internal state changes.
- __Strategy:__ Allows one of a family of algorithms to be selected on-the-fly at run-time.
- __Template Method:__ Defines the skeleton of an algorithm as an abstract class, allowing its sub-classes to provide concrete behavior.
- __Visitor:__ Separates an algorithm from an object structure by moving the hierarchy of methods into one object.