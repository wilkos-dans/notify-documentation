---
title: Notification Patterns
description:
date: 2018-12-20
---

*Notify* notifications are designed to be sent and received using the [W3C Linked Data Notifications (LDN)](https://www.w3.org/TR/2017/REC-ldn-20170502/) standard. Payloads have a predictable structure, based primarily on [Activity Streams 2.0](https://www.w3.org/TR/activitystreams-core/), with some additional vocabularies included for particular properties.

{{% rfc_2119 %}}

## Core Payload
All *Notify* payloads define an *activity*. As such, they MUST use Activity Streams 2.0 for the default context. They MAY add additional contexts. The following properties from Activity Streams 2.0 are used consistently in the various notification patterns:

### Mandatory properties in all *Notify* payloads
* **Origin:** The originator of the activity, typically the service responsible for sending the notification.
* **Target:** The intended destination of the activity, typically the service which *consumes* the notification.









The following notification patterns are defined to be widely reusable. Their re-use is illustrated in the [example scenarios](/scenarios/).