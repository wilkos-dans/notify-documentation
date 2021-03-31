---
title: "Request (by author) for review of repository resource (via overlay journal)"
date: "2021-03-08"
description: ""
status: [review,draft]
payload:
    contexts: ["sorg","ldp","ietf","nat"]
    id: "urn:uuid:0370c0fb-bb78-4a9b-87f5-bed307a509dd"
    type: ["Offer","nat:review-request"]
    origin:
        lookup: "overlay-journal"
    target:
        lookup: "repository"
    object:
        lookup: preprint
    actor:
        lookup: author
        obligation: MUST
---

