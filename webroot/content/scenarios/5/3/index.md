---
title: Acknowledge acceptance
description: The review service acknowledges the request, indicating that it accepts the request
date: "2021-03-08"
scope: notify
sender: actor_2
pattern: acknowledgement/offer_accepted
Payload:
    origin:
        lookup: "review-service"
    target:
        lookup: "repository"
    object:
        payload:
            id: urn:uuid:0370c0fb-bb78-4a9b-87f5-bed307a509dd
            type: ['Offer','coar-notify:ReviewRequest']
            object: "https://repository.org/resource/0021"
    in_reply_to:
        id: urn:uuid:0370c0fb-bb78-4a9b-87f5-bed307a509dd
---
