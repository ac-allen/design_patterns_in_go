#### Visitor: A pattern where a component (visitor) is allowed to traverse the entire hierarchy of types. Implemented by propagating a single Accept() method  throughout the entire hierarchy.

Need to define a new operation on an entire type hierarchy
- E.g., given a document model (lists, paragraphs, etc.), we want to add  
printing functionality
- Do not want to keep modifying every type in the hierarchy
- Want to have the new functionality separate (SRP)  
This approach is often used for traversal
- Alternative to Iterator
- Hierarchy members help you traverse themselves