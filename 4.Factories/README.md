# Factories design pattern

#### The factory pattern is a commonly used pattern in object oriented programming. In Go, there are many different ways in which you can use the factory pattern to make your code cleaner and more concise.

### Simple factory
The simplest, and most commonly used factory is a function that takes in some arguments, and returns an instance of Person. We can also return pointers to the Person instance instead.

Factory functions are a better alternative to initializing instances using something like p := &Person{} because, the function signature ensures that everyone will supply the required attributes.

### Interface factories
Factory functions can return interfaces instead of structs. Interfaces allow you to define behavior without exposing internal implementation. This means we can make structs private, while only exposing the interface outside our package.

### Factory generators
Factory generators are factories of factories… factoryception! While this may seem a bit excessive, factory generators are incredibly useful, when you want to construct instances of different structs or interfaces that are not mutually exclusive, or if you want multiple factories with different defaults.