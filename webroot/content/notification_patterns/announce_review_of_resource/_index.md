---
title: Announcing a review of a resource
description: Announcing a review of a resource. In payload example 1, the actor is an identified reviewer. In payload example 2, the reviewer is not identified.
payload_example_1: |4
    {
      "@context": "https://www.w3.org/ns/activitystreams",
      "@id": "",
      "@type": "Announce",
      "actor": "https://orcid.org/0000-0002-1825-0097",
      "object": "http://example.net/review",
      "target": "http://example.org/repository_article",
      "updated": "2016-06-28T19:56:20.114Z"
    }
payload_example_2: |-
    {
      "@context": "https://www.w3.org/ns/activitystreams",
      "@id": "",
      "@type": "Announce",
      "actor": "",
      "object": "http://example.net/review",
      "target": "http://example.org/repository_article",
      "updated": "2016-06-28T19:56:20.114Z"
    }
payload_example_3: ""
payload_example_4: ""
date: "2021-02-24"
layout: notification_pattern
---

