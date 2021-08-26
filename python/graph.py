from collections import defaultdict

class Graph:
  
    def __init__(self, n_vertices):
        self.n_vertices = n_vertices
        #self.graph = defaultdict(list)
        self.graph = {n: [] for n in range(self.n_vertices)}

    def __repr__(self):
        return str(self.graph)

    def addEdge(self, u, v):
        self.graph[u].append(v)

    def constructGraph(self, relation):
        for (fromNode, toNode) in relation:
            self.addEdge(fromNode, toNode)

    def printEdges(self):
        list_edges = [[(i, n) for n in self.graph[i]] for i in self.graph]
        edges = [item for sublist in list_edges for item in sublist]
        print(edges)
    
    def graphFromMatrix(self, matrix):
        n_nodes = len(matrix)
        #adj_list = {n: [] for n in range(n_nodes)}
        #adj_list = defaultdict(list)
        for i in range(n_nodes):
            for j in range(n_nodes):
                if matrix[i][j] == 1:
                    self.graph[i].append(j)

    def graphToMatrix(self):
        matrix = []
        for i in range(len(self.graph)):
            matrix.append([0]*len(self.graph))
            for j in self.graph[i]:
                matrix[i][j] = 1
        return matrix