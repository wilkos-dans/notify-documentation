---
title: "Notification Patterns"
description:
date: 2018-12-20
---

*Notify* notifications are designed to be sent and received using the [W3C Linked Data Notifications (LDN)](https://www.w3.org/TR/2017/REC-ldn-20170502/) standard. Payloads have a predictable structure, based primarily
on [Activity Streams 2.0](https://www.w3.org/TR/activitystreams-core/), with some additional vocabularies included for particular properties.

{{% rfc_2119 %}}

## Core Payload

All *Notify* payloads define an *activity*. As such, they MUST use Activity Streams 2.0 for the default context. They MAY add additional contexts. The following properties from Activity Streams 2.0 are used consistently in the various
notification patterns:

* **origin:** The originator of the activity, typically the service responsible for sending the notification.
* **target:** The intended destination of the activity, typically the service which *consumes* the notification.
* **activity:**
* **@id:** This MUST be a URI. The use of URN:UUID is RECOMMENDED. An HTTP URI MAY be used, but in such cases the URI SHOULD resolve to a useful resource.
* **@type:** This MUST be an array, which MUST include one of the Activity Stream 2.0 Activity Types, and which MUST also include a type from the [Notify Activity Types vocabulary](/vocabularies/activity_types/)

## Representation of repository resources

By default, repository resources SHOULD be represented as follows:

<div class="row">
    <div class="col">
        {{< load_local_html "repo_object_example.html" >}}
    </div>
    <div class="col">
        <h5>Example</h5>
        {{< load_local_json "repo_object_example.json" >}}
    </div>
</div>
<br/>


The following notification patterns are defined to be widely reusable. Their re-use is illustrated in the [example scenarios](/scenarios/).