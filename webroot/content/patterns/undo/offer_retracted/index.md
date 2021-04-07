---
title: "Retraction of offer"
date: "2021-03-08"
description: ""
status: [review,draft]
payload:
    contexts: ["sorg","ldp","nat"]
    id: "urn:uuid:6eafed1d-935c-41b1-a5bb-645be4b7533f"
    type: ["Undo"]
    origin:
        lookup: repository
    target:
        lookup: "review-service"
    object:
        lookup: "activity to be revoked"
        payload:
            "@id": "urn:uuid:0370c0fb-bb78-4a9b-87f5-bed307a509dd"
            "@type": ["Offer","nat:review-request"]

---

