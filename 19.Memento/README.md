#### Memento: A token representing the system state. Lets us roll back to the state when the token was generated. May or may not directly expose state information.

An object or system goes through changes
- E.g, abankaccount gets deposits and withdrawals

There are different ways of navigating those changes
- One way is to record every change (Command) and teach a command to ‘undo’ itself
- Also part of CQRS = Command Query Responsibility Segregation  

Another is to simply save snapshots of the system