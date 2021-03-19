---
headless: true
---

* The default `context` is Activity Streams 2.0
* Other contexts (namespaces) used:
    * [schema.org](https://schema.org)
    * [W3C Linked Data Platform (LDP) Vocabulary](https://www.w3.org/ns/ldp#)
    * [IETF Link Relations](http://www.iana.org/assignments/relation/)
    * [Notify Activity Types](http://purl.org/coar/notify_activity_type/)
    * [COAR Resource Types](http://purl.org/coar/resource_type)
* _Notify_ notifications always describe an Activity Streams 2.0 `activity` as the base entity
    * The `activity` has an `@id` which is an URN:UUID
    * The `activity` has a `@type` array with values from Activity Streams 2.0 and Notify Activity Types
    * The `activity` has an `origin` representing its 'originator'
    * The `activity` has a `target` representing its 'destination'
    * the resource being offered is represented as an `object`
        * the `@id` property of the `object` contains the HTTP URI to the "landing page" for the resource
        * the `@type` property of the `object` contains the value _AboutPage_ from the [schema.org](https://schema.org/AboutPage) vocabulary
        * the `ietf:cite-as` property of the `object` contains the persistent HTTP URI (sometimes called the "PID") which is to be used to cite or link to the resource.
        * the `url` property contains the details of the actual content resource
            * the `@id` property of the `url` contains the HTTP URI to the content file for the resource
            * the `@type` property of the `url` describes the content file, including values from three vocabularies
            * the `mediaType` property of the `url` contains the MIME Type of the content file