---
title: Announce Ingest
description: "The repository notifies the overlay journal that it has ingested the resource"
date: "2021-03-08"
scope: notify
sender: "actor_2"
pattern: "announcement/announcement_in_reply_to"
payload:
    contexts: ["sorg","ldp","ietf","nat","nrr"]
    id: "urn:uuid:94ecae35-dcfd-4182-8550-22c7164fe23f"
    type: ["Announce","nat:ingest-success"]
    origin:
        lookup: "repository"
    target:
        lookup: "overlay-journal"
    object:
        lookup: "preprint"
    in_reply_to:
        id: urn:uuid:0370c0fb-bb78-4a9b-87f5-bed307a509dd
        type: ["Offer","nat:ingest-request"]
        object: 
---

