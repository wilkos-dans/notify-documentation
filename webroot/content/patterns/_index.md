---
title: "Notification Patterns"
description:
date: 2018-12-20
layout: overview
---

*Notify* notifications are designed to be sent and received using the [W3C Linked Data Notifications (LDN)](https://www.w3.org/TR/2017/REC-ldn-20170502/) standard. Payloads have a predictable structure, based primarily
on [Activity Streams 2.0](https://www.w3.org/TR/activitystreams-core/), with some additional vocabularies included for particular properties.

{{% rfc_2119 %}}

## Core Payload
All *Notify* payloads define an `activity`. As such, they MUST use Activity Streams 2.0 for the default context. They MAY add additional contexts. The following properties from Activity Streams 2.0 are used consistently in all the
notification patterns:

* **The `activity`** property MUST contain the following properties:
    * **`id`:** This MUST be a URI. The use of URN:UUID is RECOMMENDED. An HTTP URI MAY be used, but in such cases the URI SHOULD resolve to a useful resource.
    * **`type`:** This MUST include one of the [Activity Stream 2.0 Activity Types](https://www.w3.org/TR/activitystreams-vocabulary/) and MAY (depending on the activity) also include a type from the [Notify Activity Types vocabulary](/vocabularies/activity_types/)
    * **`origin`:** The originator of the activity, typically the service responsible for *sending* the notification.
    * **`target`:** The intended destination of the activity, typically the service which *consumes* the notification.
* **The `activity`** property MAY (and often will) contain the following properties:
    * **`object`:** This MUST be a Web *resource*, generally the focus of the activity. Other object properties MAY appear in notifications, as properties of other properties. The activity itself MUST have zero or one object properties.
    * **`actor`:** Sometimes it is useful to identify the party or process that initiated the activity, in which case this initiator MUST be expressed as an `actor`.

The following notification patterns are defined to be widely reusable. Their re-use is illustrated in the [example scenarios](/scenarios/).