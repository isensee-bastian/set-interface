# Generics Part 2: Set Implementation with empty interface
* Why generalize our set implementation
    * Allows code reuse with other element types
* Using the empty interface type
    * No restrictions, satisfied by every specifc type
    * Allows code reuse
    * Common solution before Go 1.18
    * Element types can be mixed (often not wanted)
    * No guarantees by the compiler, requires type assertions at runtime
