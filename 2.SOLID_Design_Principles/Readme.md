# The SOLID ideas are:

## The Single-responsibility principle: 
#### "There should never be more than one reason for a class to change."In other words, every class should have only one responsibility.

This principle states that a class should have only one reason to change. If we violate this principle, the class will have multiple responsibilities, making it harder to maintain, test and extend. This can lead to code that is tightly coupled, difficult to reuse, and prone to errors.

## The Open–closed principle: 
#### "Software entities ... should be open for extension, but closed for modification."

This principle states that classes should be open for extension but closed for modification. If we violate this principle, we may have to modify existing code to add new functionality, which can introduce bugs and make it difficult to maintain the code. This can also result in code that is difficult to test and reuse.

## The Liskov substitution principle: 
#### "Functions that use pointers or references to base classes must be able to use objects of derived classes without knowing it."

This principle states that subtypes should be substitutable for their base types. If we violate this principle, we may introduce behavior that is unexpected and inconsistent, which can lead to errors that are difficult to track down. This can also make it difficult to write code that works with a variety of different types.

## The Interface segregation principle: 
#### "Clients should not be forced to depend upon interfaces that they do not use."

This principle states that clients should not be forced to depend on interfaces they do not use. If we violate this principle, we may have interfaces that are too large and contain methods that are not relevant to some clients, which can lead to code that is difficult to understand and maintain. This can also result in code that is not reusable, and that can cause unnecessary coupling between modules.

## The Dependency inversion principle: 
#### "Depend upon abstractions, not concretions."

This principle states that high-level modules should not depend on low-level modules. Instead, both should depend on abstractions. If we violate this principle, we may have code that is difficult to test and reuse, and that is tightly coupled. This can also result in code that is difficult to maintain and extend.