# Transitive Closure
In this repository I am collecting some implementations for computation of the reflexive and non-reflexive transitive closure. I see this problem as a good practice which is why I decided to implement different algorithms in Python and in Go.

Let $M$ be set and $R$ a relation on $M$. Then the relation $R^n$ is recursively defined for all $n \in \mathbb{N}_0$ with $R^0 := Id_M$ and $R^n := R^{n-1}R$ for all $n \in \mathbb{N}$.

$$R^n = \{ (x,y) \in M \times M | \exists z_0,\dots,z_n \in M: z_0 = x \land z_n = y \land (z_0,z_1) \in R \land (z_1,z_2) \in R \land \dots \land (z_{n-1},z_n) \in R \}$$

If you translate the relation into a graph, the tuple $(x,y)$ is in $R^n$ when there is a finite path of $n$ Arrows in $R$ from $x$ to $y$.

## Transitive closure of a relation
The transitive closure of the relation $R$ is the relation $R^{+}$ defined by: $R^{+} := \bigcup_{n \in \mathbb{N}} R^n$
It can be constructed by adding an edge between one node and all nodes reachable from said node.

## Reflexive transitive closure of a relation
The reflexive transitive closure of the relation $R$ is the relation $R^{\*}$ defined by: $R^{\*} := R^{+} \cup Id_M = \bigcup_{n \in \mathbb{N}_0} R^n$
It can be constructed the same way the transitive closure gets constructed with the addition of loops for every node in the graph.

## Implementation
The algorithms and their implementation are inspired by and taken from the work of **Esko Nuutila**.
Nuutila, E., **Efficient Transitive Closure Computation in Large Digraphs**. Acta Polytechnica Scandinavica, Mathematics and Computing in Engineering Series No. 74, Helsinki 1995, 124 pages. Published by the Finnish Academy of Technology. ISBN 951-666-451-2. ISSN 1237-2404. UDC 681.3.
http://www.cs.hut.fi/~enu/thesis.html

## TODO
- Implement Simple_TC in Python and Go
- Write scripts to compare the runtime
- Write tests
- Add graphics for illustration
- Fix curly brackets in readme
