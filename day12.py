
# Python program for Dijkstra's single
# source shortest path algorithm. The program is
# for adjacency matrix representation of the graph
 
# Library for INT_MAX
import sys
 
 
class Graph():
 
    def __init__(self, vertices):
        self.V = vertices
        self.graph = [[0 for column in range(vertices)]
                      for row in range(vertices)]
 
    def printSolution(self, dist):
        print("Vertex \tDistance from Source")
        for node in range(self.V):
            print(node, "\t", dist[node])
 
    # A utility function to find the vertex with
    # minimum distance value, from the set of vertices
    # not yet included in shortest path tree
    def minDistance(self, dist, sptSet):
 
        # Initialize minimum distance for next node
        min = sys.maxsize
 
        # Search not nearest vertex not in the
        # shortest path tree
        for u in range(self.V):
            if dist[u] < min and sptSet[u] == False:
                min = dist[u]
                min_index = u
 
        return min_index
 
    # Function that implements Dijkstra's single source
    # shortest path algorithm for a graph represented
    # using adjacency matrix representation
    def dijkstra(self, src):
 
        dist = [sys.maxsize] * self.V
        dist[src] = 0
        sptSet = [False] * self.V
 
        for cout in range(self.V):
 
            # Pick the minimum distance vertex from
            # the set of vertices not yet processed.
            # x is always equal to src in first iteration
            x = self.minDistance(dist, sptSet)
 
            # Put the minimum distance vertex in the
            # shortest path tree
            sptSet[x] = True
 
            # Update dist value of the adjacent vertices
            # of the picked vertex only if the current
            # distance is greater than new distance and
            # the vertex in not in the shortest path tree
            for y in range(self.V):
                if self.graph[x][y] > 0 and sptSet[y] == False and \
                        dist[y] > dist[x] + self.graph[x][y]:
                    dist[y] = dist[x] + self.graph[x][y]
 
        self.printSolution(dist)
 
def assignValue(x):
    c = ["S", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "E"]
    for i in range(len(c)):
        if c[i] == x:
            return i
 
# Driver's code
if __name__ == "__main__":
    file = open('path.txt', 'r')
    lines = file.readlines()

    for index, line in enumerate(lines):
        x = line.split()
        for i in range(len(x)):
            print(x[i])
            for j in range(len(x[i])):
                print()
    
    file.close()
    g = Graph(0)

    # g.graph = [[0, 1, 2, 17, 16, 0, 0, 8, 0],
    #            [1, 0, 8, 0, 0, 0, 0, 11, 0],
    #            [1, 8, 0, 7, 0, 4, 0, 0, 2],
    #            [1, 0, 7, 0, 9, 14, 0, 0, 0],
    #            [1, 2, 4, 5, 6, 7, 8, 9],
    #            [0, 0, 4, 14, 10, 0, 2, 0, 0],
    #            [0, 0, 0, 0, 0, 2, 0, 1, 6],
    #            [8, 11, 0, 0, 0, 0, 1, 0, 7],
    #            [0, 0, 2, 0, 0, 0, 6, 7, 0]
    #            ]
 