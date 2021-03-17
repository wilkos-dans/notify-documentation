---
title: Notification Patterns
description: 
date: 2018-12-20
---

*Notify* notifications are designed to be sent and received using
the [W3C Linked Data Notifications (LDN)](https://www.w3.org/TR/2017/REC-ldn-20170502/) standard. Payloads have a
predictable structure, based primarily on [Activity Streams 2.0](https://www.w3.org/TR/activitystreams-core/), with some
additional vocabularies included for particular properties.

{{% rfc_2119 %}}

## Core Payload

All *Notify* payloads define an *activity*. As such, they MUST use Activity Streams 2.0 for the default context. They
MAY add additional contexts. The following properties from Activity Streams 2.0 are used consistently in the various
notification patterns:

* **origin:** The originator of the activity, typically the service responsible for sending the notification.
* **target:** The intended destination of the activity, typically the service which *consumes* the notification.
* **activity:**
    * **@id:** This MUST be a URI. The use of URN:UUID is RECOMMENDED. An HTTP URI MAY be used, but in such cases the
      URI SHOULD resolve to a useful resource.
    * **@type:** This MUST be an array, which MUST include one of the Activity Stream 2.0 Activity Types, and which MUST
      also include a type from the [Notify Activity Types vocabulary](/vocabularies/activity_types/)

## Representation of repository resources

By default, repository resources SHOULD be represented as follows:

<div class="row">
  <div class="col">
  <h5>Properties</h5>  
    <ul>
      <li>the repository resource is represented as an <i>object</i></li>
      <li>the <i>@id</i> property of the <i>object</i> contains the HTTP URI to the "landing page" for the resource</li>
      <li>the <i>@type</i> property of the <i>object</i> contains the value <i>AboutPage</i> from the <a href="https://schema.org/AboutPage">schema.org</a> vocabulary</li>
      <li>the <i>ietf:cite-as</i> property of the <i>object</i> contains the persistent HTTP URI (sometimes called the "PID") which is to be used to cite or link to the resource.</li>
      <li>the <i>url</i> property contains the details of the actual content resource</li>
      <ul>
        <li>the <i>@id</i> property of the <i>url</i> contains the HTTP URI to the content file for the resource</li>
        <li>the <i>@type</i> property of the <i>url</i> describes the content file, including values from three vocabularies:</li>
        <ul>
          <li>Activity Streams 2.0</li>
          <li>schema.org</li>
          <li>COAR Resource Types</li>
        </ul>
        <li>the <i>mediaType</i> property of the <i>url</i> contains the MIME Type of the content file</li>
      </ul>
    </ul>
  </div>
  <div class="col">
    <h5>Example</h5>
    <pre><code class="language-json">"object": {
      "@id": "https://repository.org/preprint/201203/421/",
      "@type": "sorg:AboutPage",
      "ietf:cite-as": "https://doi.org/10.5555/12345680",
      "url": {
        "@id": "https://repository.org/preprint/201203/421/content.pdf",
        "@type": [
          "Article",
          "sorg:ScholarlyArticle",
          "rt:c_816b"
        ],
        "mediaType": "application/pdf"
      }
    }
    </code></pre>
  </div>
</div>
<br />




The following notification patterns are defined to be widely reusable. Their re-use is illustrated in
the [example scenarios](/scenarios/).