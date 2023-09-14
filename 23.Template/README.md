#### Template: A skeleton algorithm defined in a function. Function can either use an interface (like Strategy) or can take several functions as arguments.

Algorithms can be decomposed into common parts + specifics
Strategy pattern does this through composition

- High-level algorithm uses an interface

- Concrete implementations implement the interface

- We keep a pointer to the interface; provide concrete implementations  

Template Method performs a similar operation, but

- It’s typically just a function, not a struct with a reference to the

implementation
- Can still use interfaces (just like Strategy); or
- Can be functional (take several functions as parameters)