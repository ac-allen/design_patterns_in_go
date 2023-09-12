#### Mediator: - A component that facilitates communication between other components without them necessarily being aware of each other or having direct (reference) access to each other.

Components may go in and out of a system at any time
- Chat room participants
- Players in an MMORPG

It makes no sense for them to have direct references (pointers) to one another
- Those references may go dead  

Solution: have then all refer to some central component that
facilitates communication