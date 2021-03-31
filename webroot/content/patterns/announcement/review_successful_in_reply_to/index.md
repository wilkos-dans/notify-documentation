---
title: "Announcement of review, in response to offer"
date: "2021-03-08"
description: ""
layout: pattern_example
status: [review,draft]
payload:
    contexts: ["sorg","ldp","ietf","nat","nrr"]
    id: "urn:uuid:94ecae35-dcfd-4182-8550-22c7164fe23f"
    type: ["Announce","nat:review-success"]
    origin:
        lookup: "review-service"
    target:
        lookup: repository
    object:
        lookup: review
    in_reply_to:
        id: urn:uuid:0370c0fb-bb78-4a9b-87f5-bed307a509dd
        type: ["Offer","nat:review-request"]
        object:
            lookup: preprint
    actor:
        lookup: reviewer
---