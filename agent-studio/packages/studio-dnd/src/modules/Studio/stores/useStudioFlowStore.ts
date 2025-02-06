import { addEdge, applyEdgeChanges, applyNodeChanges, Edge, OnConnect, OnEdgesChange, OnNodesChange } from '@xyflow/react';
import { create } from 'zustand';

import { StudioNode } from '../types/graph';

const DEFAULT_VALUE = {
  reloadFlowCounter: 0,
  nodes: [],
  edges: [],
  linkedNodes: {},
};

type Store = {
  reloadFlowCounter: number;
  reloadFlow: () => void;

  nodes: StudioNode[];
  setNodes: (nodes: StudioNode[]) => void;
  addNode: (node: StudioNode) => void;
  addNodes: (nodes: StudioNode[]) => void;

  addLinkedNode: (nodeId: string, linkedNodeId: string) => void;
  linkedNodes: Record<string, string[]>;
  setLinkedNodes: (nodes: Record<string, string[]>) => void;

  updateNode: (node: StudioNode) => void;
  updateNodes: (nodes: StudioNode[]) => void;

  edges: Edge[];
  setEdges: (edges: Edge[]) => void;
  addEdge: (edge: Edge) => void;
  addEdges: (edges: Edge[]) => void;

  removeEdge: (edgeId: string) => void;
  removeNode: (id: string) => void;
  removeNodeAndAllBelong: (id: string) => void;

  removeLinkedNode: (nodeId: string, linkedNodeId: string) => void;

  onNodesChange: OnNodesChange<StudioNode>;
  onEdgesChange: OnEdgesChange;
  onConnect: OnConnect;

  clear: () => void;
};

const useStudioFlowStore = create<Store>((set, get) => ({
  ...DEFAULT_VALUE,

  reloadFlow: () => {
    set({ reloadFlowCounter: get().reloadFlowCounter + 1 });
  },

  setNodes: (nodes) => set({ nodes }),
  addNode: (node) => set({ nodes: [...get().nodes, node] }),
  addNodes: (nodes) => {
    set((state) => {
      // Prevent duplicate nodes
      const existingNodeIds = new Set(state.nodes.map((n) => n.id));
      const newNodes = nodes.filter((node) => !existingNodeIds.has(node.id));

      if (newNodes.length === 0) return state;

      return {
        nodes: [...state.nodes, ...newNodes],
      };
    });
  },

  addLinkedNode: (nodeId, linkedNodeId) => {
    set({
      linkedNodes: {
        ...get().linkedNodes,
        [nodeId]: [...(get().linkedNodes[nodeId] || []), linkedNodeId],
      },
    });
  },

  setLinkedNodes: (nodes) => set({ linkedNodes: nodes }),

  updateNode: (node) => {
    const updatedNodes = get().nodes.map((n) => (n.id === node.id ? node : n));
    set({ nodes: updatedNodes });
  },
  updateNodes: (nodes) => {
    const updatedNodes = get().nodes.map((n) => nodes.find((node) => node.id === n.id) || n);

    set({ nodes: updatedNodes });
  },

  setEdges: (edges) => set({ edges }),
  addEdge: (edge) => set({ edges: [...get().edges, edge] }),
  addEdges: (edges) => set({ edges: [...get().edges, ...edges] }),

  removeNode: (id: string) => {
    set((state) => {
      // Filter out removed nodes
      const updatedNodes = state.nodes.filter((node) => node.id !== id);

      // Filter out edges connected to removed nodes
      const updatedEdges = state.edges.filter((edge) => edge.source !== id && edge.target !== id);

      // Clean up linked nodes references
      const updatedLinkedNodes = Object.fromEntries(
        Object.entries(state.linkedNodes)
          .filter(([nodeId]) => nodeId !== id)
          .map(([nodeId, linkedIds]) => [nodeId, linkedIds.filter((linkedId) => linkedId !== id)])
          .filter(([, linkedIds]) => linkedIds.length > 0),
      );

      return {
        nodes: updatedNodes,
        edges: updatedEdges,
        linkedNodes: updatedLinkedNodes,
      };
    });
  },
  removeNodeAndAllBelong: (id: string) => {
    set((state) => {
      const nodesToRemove = new Set<string>();

      // Recursive function to collect all related nodes
      const collectNodesToRemove = (nodeId: string, visited = new Set<string>()) => {
        if (visited.has(nodeId)) return;

        visited.add(nodeId);
        nodesToRemove.add(nodeId);

        // Get direct linked nodes
        const directLinkedNodes = state.linkedNodes[nodeId] || [];

        // Get nodes that have edges to/from this node
        const connectedNodes = state.edges.filter((edge) => edge.source === nodeId).map((edge) => edge.target);

        // Process all related nodes
        [...directLinkedNodes, ...connectedNodes].forEach((relatedId) => {
          collectNodesToRemove(relatedId, visited);
        });
      };

      // Start recursive collection
      collectNodesToRemove(id);

      // Filter out removed nodes
      const updatedNodes = state.nodes.filter((node) => !nodesToRemove.has(node.id));

      // Filter out edges connected to removed nodes
      const updatedEdges = state.edges.filter((edge) => !nodesToRemove.has(edge.source) && !nodesToRemove.has(edge.target));

      // Clean up linked nodes references
      const updatedLinkedNodes = Object.fromEntries(
        Object.entries(state.linkedNodes)
          .filter(([nodeId]) => !nodesToRemove.has(nodeId))
          .map(([nodeId, linkedIds]) => [nodeId, linkedIds.filter((linkedId) => !nodesToRemove.has(linkedId))])
          .filter(([, linkedIds]) => linkedIds.length > 0),
      );

      return {
        nodes: updatedNodes,
        edges: updatedEdges,
        linkedNodes: updatedLinkedNodes,
      };
    });
  },

  removeEdge: (edgeId: string) => {
    set({ edges: get().edges.filter((edge) => edge.id !== edgeId) });
  },

  removeLinkedNode: (nodeId: string, linkedNodeId: string) => {
    set({
      linkedNodes: {
        ...get().linkedNodes,
        [nodeId]: get().linkedNodes[nodeId].filter((id) => id !== linkedNodeId),
      },
    });
  },

  onNodesChange: (changes) => {
    set({
      nodes: applyNodeChanges(changes, get().nodes),
    });
  },
  onEdgesChange: (changes) => {
    set({
      edges: applyEdgeChanges(changes, get().edges),
    });
  },
  onConnect: (connection) => {
    set((state) => {
      // Prevent self-connections
      if (connection.source === connection.target) {
        return state;
      }

      // Prevent duplicate connections
      const isDuplicate = state.edges.some((edge) => edge.source === connection.source && edge.target === connection.target);

      if (isDuplicate) {
        return state;
      }

      return {
        edges: addEdge(connection, state.edges),
      };
    });
  },

  clear: () => set(DEFAULT_VALUE),
}));

export default useStudioFlowStore;
