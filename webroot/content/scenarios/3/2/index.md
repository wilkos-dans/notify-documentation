---
title: Announce Review
description: 'The overlay journal notifies the aggregator of a new review '
date: "2021-03-08"
scope: notify
sender: actor_1
pattern: announcement/announcement
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

