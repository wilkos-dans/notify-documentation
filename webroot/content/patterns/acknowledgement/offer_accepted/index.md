---
title: "Acceptance of an offer"
date: "2021-03-08"
description: "This acknowledgement pattern is used to accept an offer made in a previous notification."
layout: pattern_example
status: [review,draft]
payload:
    id: "urn:uuid:4fb3af44-d4f8-4226-9475-2d09c2d8d9e0"
    type: ["Accept"]
    origin:
        lookup: "generic-origin-system"
    target:
        lookup: "generic-target-system"
    object:
        payload:
            id: urn:uuid:0370c0fb-bb78-4a9b-87f5-bed307a509dd
            type: ['Offer']
    context:
        lookup: "generic-object-repository"
    in_reply_to:
        id: urn:uuid:0370c0fb-bb78-4a9b-87f5-bed307a509dd
---

