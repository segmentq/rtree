# rtree
_A single dimension version of [tidwall/rtree](https://github.com/tidwall/rtree)_

[![Build](https://github.com/segmentq/rtree/actions/workflows/build.yml/badge.svg)](https://github.com/segmentq/rtree/actions/workflows/build.yml)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=segmentq_rtree&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=segmentq_rtree)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=segmentq_rtree&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=segmentq_rtree)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=segmentq_rtree&metric=coverage)](https://sonarcloud.io/summary/new_code?id=segmentq_rtree)
---

## Usage
```go
import "github.com/segmentq/rtree"

// Init the tree
rt := rtree.NewOneD[int, string]()

// Insert a range
rt.Insert(1, 10, "wow cool rtree!")

// Search for it...
rt.Search(5, 5, func(min, max int, data string) bool {
    println(data)
    return true
})

// Delete it
rt.Delete(1, 10, "wow cool rtree!")

// Cleanup
rt.Clear()
```
