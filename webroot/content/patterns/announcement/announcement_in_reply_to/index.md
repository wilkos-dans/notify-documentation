---
title: "Announcement (in response to offer)"
date: "2021-03-08"
description: "This pattern is used to announce the outcome of an activity, sometimes (but not always) linking an original resource to a new, related resource, in response to an offer made in a previous notification."
layout: pattern_example
status: [review,draft]
weight: 2
payload:
    contexts: ["ldp","ietf"]
    id: "urn:uuid:94ecae35-dcfd-4182-8550-22c7164fe23f"
    type: ["Announce"]
    origin:
        lookup: "generic-origin-system"
    target:
        lookup: "generic-target-system"
    object:
        lookup: generic-object
    in_reply_to:
        id: urn:uuid:0370c0fb-bb78-4a9b-87f5-bed307a509dd
        type: ["Offer"]
        object:
            lookup: generic-object-in-reply-to

---