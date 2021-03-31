---
title: "Announcement of review by overlay journal to aggregator"
date: "2021-03-08"
description: ""
layout: pattern_example
status: [review,draft]
payload:
    contexts: ["sorg","ldp","ietf","nat","nrr"]
    id: "urn:uuid:94ecae35-dcfd-4182-8550-22c7164fe23f"
    type: ["Announce","nat:review-success"]
    origin:
        lookup: "overlay-journal"
    target:
        lookup: "aggregation-service"
    object:
        lookup: "journal-review"
    actor:
        lookup: reviewer
---