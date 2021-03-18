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
            <li>The default <code>context</code> is Activity Streams 2.0</li>
            <li>Other contexts (namespaces) used:</li>
            <ul>
                <li><a href="https://schema.org">schema.org</a></li>
                <li><a href="http://purl.org/coar/resource_type">COAR Resource Types</a></li>
            </ul>
            <li>the repository resource is represented as an <code>object</code></li>
            <ul>
                <li>the <code>@id</code> property of the <code>object</code> contains the HTTP URI to the "landing page"
                    for the
                    resource
                </li>
                <li>the <code>@type</code> property of the <code>object</code> contains the value <i>AboutPage</i> from
                    the <a
                            href="https://schema.org/AboutPage">schema.org</a> vocabulary
                </li>
                <li>the <code>ietf:cite-as</code> property of the <code>object</code> contains the persistent HTTP URI
                    (sometimes called
                    the "PID") which is to be used to cite or link to the resource.
                </li>
                <li>the <code>url</code> property contains the details of the actual content resource</li>
                <ul>
                    <li>the <code>@id</code> property of the <code>url</code> contains the HTTP URI to the content file
                        for the resource
                    </li>
                    <li>the <code>@type</code> property of the <code>url</code> describes the content file, including
                        values from three
                        vocabularies
                    </li>
                    <li>the <code>mediaType</code> property of the <code>url</code> contains the MIME Type of the
                        content file
                    </li>
                </ul>
            </ul>
        </ul>
    </div>
    <div class="col">
        <h5>Example</h5>
        {{< load_json "/patterns/repo_object_example.json" >}}
    </div>
</div>
<br/>


The following notification patterns are defined to be widely reusable. Their re-use is illustrated in
the [example scenarios](/scenarios/).