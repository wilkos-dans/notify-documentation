---
title: "Request to deposit a resource in a repository"
date: "2021-03-08"
description: ""
layout: pattern_example
status: [review,draft]
payload:
    contexts: ["sorg","ldp","ietf","nat"]
    id: "urn:uuid:0370c0fb-bb78-4a9b-87f5-bed307a509dd"
    type: ["Offer","nat:ingest-request"]
    origin:
        lookup: "overlay-journal"
    target:
        lookup: "repository"
    object:
        lookup: submission
---
