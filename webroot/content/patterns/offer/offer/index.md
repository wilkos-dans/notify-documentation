---
title: "Offer of a resource to a service"
date: "2021-03-08"
description: "This pattern is used to make an offer, where one system offers one of its resources for some activity to be conducted by a second system."
layout: pattern_example
status: [review,draft]
payload:
    id: "urn:uuid:0370c0fb-bb78-4a9b-87f5-bed307a509dd"
    type: ["Offer"]
    origin:
        lookup: "generic-origin-system"
    target:
        lookup: "generic-target-system"
    object:
        lookup: "generic-object"
    actor:
        lookup: "generic-actor"
        obligation: may
---
