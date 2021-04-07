---
title: "Rejection of offer"
date: "2021-03-08"
description: ""
layout: pattern_example
status: [review,draft]
payload:
    contexts: ["ldp","ietf","nat"]
    id: "urn:uuid:4fb3af44-d4f8-4226-9475-2d09c2d8d9e0"
    type: ["Reject"]
    origin:
        lookup: "review-service"
    target:
        lookup: "repository"
    in_reply_to:
        id: urn:uuid:0370c0fb-bb78-4a9b-87f5-bed307a509dd
        type: ["Offer","nat:review-request"]
        object:
            lookup: preprint
---

