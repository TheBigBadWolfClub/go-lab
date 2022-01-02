
# Clean Architecture | DDD | Screaming Architecture


<hr/>

# Hexagonal/Onion/Clean Architecture
[* source](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)


Any software architecture as the same objective, _the separation of responsibilities_, this separation is normally achieved by dividing the software into layers.

Over the years there is an architecture concept, that has been been constantly publish by several Authors.

Hexagonal, Onion or Clean are architectures very similar in design, where each author presents an evolution or  slight variations of the same concept. By doing so, it promotes this architecture concept as a great contender to be selected for a project.

Advantages:
- Independent of frameworks
    - easy replace DB and Frameworks.
    - software will be free of any technical constrain of those Frameworks
- Independent of system boundaries (UI, database, …)
    - any change in this or any external dependency will not impact use cases and business rules.
- Easy testable throw the use of mocks.
    - any layer in this architecture could easy be tested with mocks, and isolated from the other layers
- Flexible
    - System divided into distinct features with as little overlap in functionality as possible so that they can be combined freely.
- Evolvable
    - Is easy to adapt step by step to keep up with changes.
- Integrates with Agile principles
    -  allows quick changes through flexibility, evolvability and rapid deployment


### Concept

These architectures are defined around the concept of centric layers:
-  inner layers are responsible for business and application logic
-  outer layers are composed of frameworks and tools that can easy be swapped for better alternatives


![](https://huongdanjava.com/wp-content/uploads/2020/10/Clean-Architecture.png)

> ## The Dependency Rule
> this architecture only works because implements the dependency rule<br>
>
> "__Nothing in an inner circle can know anything at all about something in an outer circle__"
>

<br>

> ### business rules
> Business specific business rules
> - Fundamental business rules
>
> Application specific business rules
> -  example: validations, communication with user ...
>
> __It is very important to separate this rules in the application architecture__
> <br>

<br>


## Domain Driven Design


The microservice approach to division is different, splitting up into services organized around business capability. Such services take a broad-stack implementation of software for that business area, including user-interface, persistance storage, and any external collaborations. Consequently the teams are cross-functional, including the full range of skills required for the development: user-experience, database, and project management.

<hr/>

## Screaming Architecture
[* source](https://blog.cleancoder.com/uncle-bob/2011/09/30/Screaming-Architecture.html)


By the definition, a Screaming Architecture is a methodology of structuring a project. This type of structure should allow any developer to easy identify the propose of the project just by looking at is directory structure and file names.

Advantages:
- Easier to **Explain** the what the project does
- Easier to **Onboard** new developers (learning curve)
- Easier to **Identify** where a feature belongs
- Easier to **Identify and Fix** bugs, since there will be limited to a specific part of the system
- High cohesion (code that change together, should be kept together)


### Screaming example

```
internal
   auth (d)
   billing (d)
   customers (d)
   phones (d)
      controller.go
      model.go
      repository.go
```

By looking at the structure above for the first time, a new developer, could assume this is:
- some type of application to manage or use phones
- if ask a new developer to change the billing formula, he will easy identify where that formula is.
- if ask a new developer to identify a bug on reading "customer" from DB, he will easy go to the correct src file.

<br><br>

### No Screaming example
```
internal
   controllers (d)
   models (d)
   repositories (d)
      phones.go
      customers.go
      login.go
      auth.go
```

<br/><br/>
<hr/>



