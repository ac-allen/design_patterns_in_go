For some components it only makes sense to have one in the system  
:thought_balloon: Database repository  
:thought_balloon: Object factory  

E.g., the construction call is expensive  
:thought_balloon: We only do it once  
:thought_balloon: We give everyone the same instance  

Want to prevent anyone creating additional copies  

Need to take care of lazy instantiation