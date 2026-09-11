# ADR-0055: Single-document CSV output

## Status
Accepted

## Context
`--csv` emitted bare line-delimited rows with no header, awkward for agents.

## Decision
Emit one valid CSV document with a header row (`kind,key,value`) so agents can
`csv.reader` / `pandas.read_csv` the whole stdout. Errors captured inline.

## Consequences
Deterministic tabular output for agents and data pipelines.
