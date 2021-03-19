---
headless: true
---

* The default `context` is Activity Streams 2.0
* Other contexts (namespaces) used:
    * [W3C Linked Data Platform (LDP) Vocabulary](https://www.w3.org/ns/ldp#)
    * [IETF Link Relations](http://www.iana.org/assignments/relation/)
    * [Notify Activity Types](http://purl.org/coar/notify_activity_type/)
* _Notify_ notifications always describe an Activity Streams 2.0 `activity` as the base entity
    * The `activity` has an `@id` which is an URN:UUID
    * The `activity` has a `@type` array with the value *Accept* from Activity Streams 2.0, together with a more specific type from *Notify Activity Types*
    * The `activity` has an `origin` representing its 'originator'
    * The `activity` has a `target` representing its 'destination'
    * `inReplyTo` is specified because this is an acknowledgement of a previous notification
        * `inReplyTo` has an `@id` which is the URI (URN:UUID or HTTP URI) which identifies the `activity` for which this is a response
      * `inReplyTo` has a `@type` array which mirrors the `@type` of the `activity` for which this is a response
      * `inReplyTo` also contains an `object` which represents the original `object` of the `activity` for which this is a response
        * the `@id` property of the `object` contains the HTTP URI for the resource
        * the `ietf:cite-as` property of the `object` contains the persistent HTTP URI (sometimes called the "PID") which is to be used to cite or link to the resource.