# Sentry — Development Progress

**Project:** Sentry — Engineering Intelligence & Decision Orchestration System  
**Language:** Go  
**LLM:** Qwen 2.5 7B via Ollama  
**Architecture:** Knowledge Graph + AST + Context Compression + Socratic Reasoning

## Phase 0 — Research & Architecture

- [x] Finalize system architecture.
- [ ] Define supported programming languages for initial AST analysis.
- [ ] Research AST parsing libraries and language-specific parsers.
- [ ] Define engineering knowledge graph schema.
- [ ] Define graph node types and relationship types.
- [ ] Define observed, proposed, accepted and rejected states.
- [ ] Define graph versioning and proposed-change handling.
- [ ] Establish initial performance and accuracy benchmarks.

## Phase 1 — Project Foundation

- [ ] Initialize Go project and module structure.
- [ ] Establish modular application architecture.
- [ ] Implement configuration management.
- [ ] Implement Ollama client integration.
- [ ] Verify Qwen 2.5 7B inference.
- [ ] Implement structured logging and error handling.
- [ ] Define internal interfaces for graph, parser and reasoning components.

## Phase 2 — AST Analysis Engine

- [ ] Integrate Tree-sitter for multi-language parsing.
- [ ] Implement Go-native AST analysis using `go/ast`.
- [ ] Extract packages and modules.
- [ ] Extract functions and methods.
- [ ] Extract structs, classes and interfaces.
- [ ] Extract imports and dependencies.
- [ ] Identify function calls and references.
- [ ] Identify API handlers and routes.
- [ ] Identify database models and operations.
- [ ] Identify event publishers and consumers.
- [ ] Implement incremental codebase analysis.
- [ ] Implement source-location tracking for extracted entities.

## Phase 3 — Engineering Knowledge Graph

- [ ] Design graph node and edge schemas.
- [ ] Implement graph node management.
- [ ] Implement relationship management.
- [ ] Define service, function, API, database, event and business-rule nodes.
- [ ] Define dependency, invocation, read, write, publish and consume relationships.
- [ ] Implement graph traversal algorithms.
- [ ] Implement dependency discovery.
- [ ] Implement graph persistence.
- [ ] Track provenance, confidence and timestamps.
- [ ] Support observed, proposed, accepted and rejected entities.
- [ ] Implement graph updates when the codebase changes.
- [ ] Support both greenfield and existing projects.

## Phase 4 — Feature Impact & Blast-Radius Analysis

- [ ] Implement feature proposal ingestion.
- [ ] Map proposed features to existing graph entities.
- [ ] Identify directly affected components.
- [ ] Identify indirectly affected components.
- [ ] Implement multi-hop dependency traversal.
- [ ] Identify affected APIs and contracts.
- [ ] Identify affected databases and data models.
- [ ] Identify affected events and consumers.
- [ ] Identify potentially violated business invariants.
- [ ] Detect dependency conflicts.
- [ ] Implement proposed-change graph generation.
- [ ] Compare current and proposed architecture graphs.
- [ ] Generate structural blast-radius reports.
- [ ] Generate semantic impact reports.

## Phase 5 — Context Retrieval & Compression

- [ ] Implement query-based graph retrieval.
- [ ] Identify relevant graph neighbourhoods.
- [ ] Implement dependency-aware context selection.
- [ ] Implement compact graph serialization.
- [ ] Implement redundant information elimination.
- [ ] Preserve critical constraints and relationships.
- [ ] Preserve decision provenance and unresolved assumptions.
- [ ] Experiment with structural compression.
- [ ] Experiment with semantic compression.
- [ ] Benchmark compression against full conversation context.
- [ ] Measure token reduction using the Qwen tokenizer.
- [ ] Evaluate information retention after compression.

## Phase 6 — Socratic Reasoning Engine

- [ ] Integrate Qwen 2.5 with the reasoning pipeline.
- [ ] Design the engineering thought-partner system prompt.
- [ ] Implement feature and design proposal interpretation.
- [ ] Implement requirement and constraint identification.
- [ ] Implement assumption identification.
- [ ] Implement architecture trade-off evaluation.
- [ ] Implement risk identification.
- [ ] Implement contradiction detection against graph facts.
- [ ] Implement Socratic question generation.
- [ ] Implement structured model responses.
- [ ] Implement response validation.
- [ ] Implement YES / NO / CONDITIONAL / INSUFFICIENT CONTEXT outcomes.
- [ ] Ensure recommendations are grounded in available evidence.

## Phase 7 — Proposed Architecture & Decision Management

- [ ] Maintain separate actual and proposed architecture graphs.
- [ ] Track proposed changes without modifying the actual graph.
- [ ] Implement developer approval and rejection workflows.
- [ ] Track architectural decisions and alternatives.
- [ ] Update graph after accepted changes.
- [ ] Reconcile proposed architecture with actual code.
- [ ] Detect architectural drift.
- [ ] Track dependencies affected by accepted decisions.

## Phase 8 — Evaluation & Benchmarking

- [ ] Create representative backend engineering test repositories.
- [ ] Test AST extraction accuracy.
- [ ] Test graph relationship accuracy.
- [ ] Test direct dependency identification.
- [ ] Test indirect dependency identification.
- [ ] Evaluate blast-radius accuracy.
- [ ] Evaluate semantic impact analysis.
- [ ] Compare full-context and compressed-context reasoning.
- [ ] Measure token reduction.
- [ ] Measure inference latency.
- [ ] Evaluate hallucination and unsupported conclusions.
- [ ] Evaluate Socratic response quality.
- [ ] Create regression tests for reasoning behaviour.

## Phase 9 — Developer Interface

- [ ] Implement project initialization.
- [ ] Implement repository ingestion.
- [ ] Implement feature proposal interface.
- [ ] Implement conversational interaction.
- [ ] Display YES / NO / CONDITIONAL outcomes.
- [ ] Display affected components and dependency paths.
- [ ] Display identified risks and unresolved assumptions.
- [ ] Display architecture decision summaries.
- [ ] Provide graph visualization.
- [ ] Support incremental project analysis.

## Initial MVP Definition

The first MVP should demonstrate the following workflow:

1. Ingest an existing Go backend repository.
2. Parse the codebase and extract structural information using AST analysis.
3. Construct an engineering knowledge graph.
4. Accept a proposed feature or architectural modification.
5. Identify directly and indirectly affected components.
6. Retrieve and compress the relevant graph context.
7. Send the context to Qwen 2.5 through Ollama.
8. Generate a grounded YES / NO / CONDITIONAL / INSUFFICIENT CONTEXT response.
9. Ask a meaningful Socratic question that helps the developer reason about the identified impact.

**Core engineering principle:** Sentry should help developers understand the consequences of their decisions rather than make those decisions for them.
