---
title: Request Review
description: 'The repository requests a review from the review service. '
date: "2021-03-08"
scope: notify
sender: actor_1
pattern: offer/review_request
payload:
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


