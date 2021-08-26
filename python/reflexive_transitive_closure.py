from graph import Graph

#edges = {n:[] for n in input_set}
#for t in relation:
#    edges[t[0]].append(t[1])

#Dict and list comprehension straight out of hell
#edges = {n:[t[1] for t in relation if t[0] == n] for n in input_set}
#print("Edges: {}".format(edges))

def DFS(graph, v):
    discovered = []
    Stack = []
    Stack.append(v)
    while len(Stack) > 0:
        v = Stack.pop(-1)
        if v not in discovered:
            discovered.append(v)
            for w in graph.graph[v]:
                if w not in discovered:
                    Stack.append(w)
    return discovered

def transitiveReflexiveClosure_DFS(input_set, graph):
    for element in input_set:
        reachable_nodes = DFS(graph, element)
        print("Node: {} | Reachable Nodes: {}".format(element, reachable_nodes))
        for node in reachable_nodes:
            relation.append((element, node))

    closure = list(set(relation))
    closure.sort(key=lambda x: x[0])
    print("Reflexive-transitive Closure: {}".format(closure))
    return closure


if __name__ == '__main__':
    input_set = [0,1,2,3,4]
    relation = [(0,1),(1,2),(0,3),(3,4),(4,3)]

    g = Graph(len(input_set))
    g.constructGraph(relation)
    transitiveReflexiveClosure_DFS(input_set, g)