# Vertical Order Traversal

> Striver SDE Sheet 113 · Day 17 Binary Tree — tag every node with (column, row), then one three-key sort does all the ordering rules at once. Go, stdlib only, tested.

## Problem

List the tree's values column by column (left to right), top to bottom within a column, and — the fiddly LeetCode-987 rule — sorted by value when two nodes occupy the same row *and* column.

## Approach

Resist maintaining ordered structures during traversal. Walk once, flattening every node to a (col, row, val) record; then a single sort with the three-field comparator *is* the specification, line for line. Grouping the sorted records into per-column slices is a linear sweep comparing neighbors. When a problem's difficulty is in its ordering rules rather than its computation, turning the rules into a comparator makes the correspondence auditable — you can read the comparator against the problem statement clause by clause.

**Complexity:** O(n log n) time, O(n) space.

## Visual walkthrough

Read the picture top to bottom — each numbered frame is one moment of the algorithm
running on this repo's test example.

![Step-by-step visual walkthrough](viz.svg)

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
