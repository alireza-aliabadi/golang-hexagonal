# golang-hexagonal
hexagonal pattern golang implementation examples
includes:
* simple banking api
* simple library booking api
* simple invoice api
## hexagonal design patter

an **architectural pattern** from family of **layered artchitectures** sperate business logic from outside by ports (interfaces) and adapters (services implementations).
components:
* **Core (domain)**: Pure business rules, independent of frameworks, databases, or UI.
* **Ports**: Interfaces that define how the core communicates with the outside world.
* **Adapters**: Implementations of those ports (e.g., REST controllers, database repositories, message brokers).
