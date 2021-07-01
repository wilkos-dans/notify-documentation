---
title: Request Review
description: 'The repository requests a review from the review service. '
date: "2021-03-08"
scope: notify
sender: actor_1
pattern: offer/offer
payload:
    contexts: ["sorg","ldp","ietf","nat"]
    id: "urn:uuid:0370c0fb-bb78-4a9b-87f5-bed307a509dd"
    type: ["Offer","coar-notify:ReviewRequest"]
    origin:
        lookup: repository
    target:
        lookup: "review-service"
    object:
        lookup: preprint
    actor:
        lookup: author
        obligation: MUST
---


