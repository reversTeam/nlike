## Scope:

 - Create an endpoint for create a topic (raw topic, don't alter the data)
 - Create an endpoint to delete a topic (raw topic, don't alter the data)
 - Create endpoint from proto files (GRPC)

## Constraints:

 - Code coverage
 - Unit test (with and without --race)
 - Fonctionnal test (with and without --race)
 - Compilation with (--race, only for dev)

## Dynamic load

Réflection of the dynamic load :
 - Reflection code
 - Embeding GRPC deploy and communicate with an other service