---
title: "Announcement of new resource, relating to a different resource, **not** in response to a previously made offer"
date: "2021-03-08"
description: ""
headless: true
---

* The default `context` is Activity Streams 2.0
* Other contexts (namespaces) used:
    * [schema.org](https://schema.org)
    * [W3C Linked Data Platform (LDP) Vocabulary](https://www.w3.org/ns/ldp#)
    * [IETF Link Relations](http://www.iana.org/assignments/relation/)
    * [Notify Activity Types](http://purl.org/coar/notify_activity_type/)
    * [COAR Resource Types](http://purl.org/coar/resource_type)
    * [CITO](http://purl.org/spar/cito/)
* _Notify_ notifications always describe an Activity Streams 2.0 `activity` as the base entity
    * The `activity` has an `@id` which is an URN:UUID
    * The `activity` has a `@type` array with the value *Announce* from Activity Streams 2.0, together with a more specific type from *Notify Activity Types*
    * The `activity` has an `origin` representing its 'originator'
    * The `activity` has a `target` representing its 'destination'
    * The `activity` contains an `object` which represents a resource created by this `activity`
        * the `@id` property of the `object` contains the HTTP URI for the resource
        * the `@type` property of the `object` contains an array of values describing the resource with terms from several vocabularies
        * the `ietf:cite-as` property of the `object` contains the persistent HTTP URI (sometimes called the "PID") which is to be used to cite or link to the resource.
        * the `cito:` property links the resource represented by the `object` belonging to this `activity`, to a related resource
            * the `cito:` property contains an `object` representing the related resource.
                * the `@id` property of the related `object` contains the HTTP URI for that resource
                * the `ietf:cite-as` property of the `object` contains the persistent HTTP URI (sometimes called the "PID") which is to be used to cite or link to the resource.