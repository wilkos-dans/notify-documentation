---
title: "Representation of repository resources"
date: "2021-03-08"
description: ""
headless: true
---

* The default `context` is Activity Streams 2.0
* Other contexts (namespaces) used:
    * [schema.org](https://schema.org)
    * [COAR Resource Types](http://purl.org/coar/resource_type)
    * [IETF Link Relations](http://www.iana.org/assignments/relation/)
* the repository resource is represented as an `object`
    * the `@id` property of the `object` contains the HTTP URI of the "landing page" for the resource
    * the `@type` property of the `object` contains the value _AboutPage_ from the [schema.org](https://schema.org/AboutPage) vocabulary
    * the `ietf:cite-as` property of the `object` contains the persistent HTTP URI (sometimes called the "PID") which is to be used to cite or link to the resource.
    * the `url` property contains the details of the actual content resource
        * the `@id` property of the `url` contains the HTTP URI to the content file for the resource
        * the `@type` property of the `url` describes the content file, including values from three vocabularies
        * the `mediaType` property of the `url` contains the MIME Type of the content file