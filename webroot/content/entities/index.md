---
title: Entities used in examples
description: ""
date: 2018-12-20
---


The following fictitious entities (people, organisations, systems, resources) are used throughout the notification patterns and scenarios.

## `origin` &amp; `target`

#### Repository
* @id: https://repository.org
* ldp:inbox: https://repository.org/inbox/

#### Review Service
* @id: https://reviewservice.org
* ldp:inbox: https://reviewservice.org/requests/inbox/

#### Overlay Journal
* @id: https://overlay-journal.org
* ldp:inbox: https://overlay-journal.org/requests/inbox/

#### Aggregator Service
* @id: https://aggregator-service.org
* ldp:inbox: https://aggregator-service.org/requests/inbox/

## `actor`

#### Author
* @id: https://orcid.org/0000-0002-1825-0097
* name: Josiah Carberry
* ldp:inbox: https://josiahcarberry.com/ldn/inbox

#### Reviewer
* @id: https://isni.org/isni/0000000122832703
* name: H G Wells
* ldp:inbox: https://hgwells.com/ldn/inbox

## `object`

#### Repository Preprint
* @id: https://repository.org/preprint/201203/421/
* ietf:cite-as: https://doi.org/10.5555/12345680
* url @id: https://repository.org/preprint/201203/421/content.pdf

#### Repository Review Resource
* @id: https://repository.org/review/00001
* ietf:cite-as: https://doi.org/10.3214/000001

#### Review Service Review Resource
* @id: https://reviewservice.org/review/geo/202103/0021
* ietf:cite-as: https://doi.org/10.3214/987654

#### Overlay Journal Review Resource
* @id: https://overlay-journal.org/reviews/000001/00001
* ietf:cite-as: https://doi.org/10.3214/12345

#### Overlay Journal Published Article
* @id: https://overlay-journal.org/articles/00001/
* ietf:cite-as: https://overlay-journal.org/articles/00001/

#### Overlay Journal - Unpublished Submission
* @id: https://overlay-journal.org/submissions/00001/
* url @id: https://overlay-journal.org/submissions/00001/files/content.pdf

## `activity`

#### Offer
* @id: urn:uuid:0370c0fb-bb78-4a9b-87f5-bed307a509dd

#### Undo
* @id: urn:uuid:6eafed1d-935c-41b1-a5bb-645be4b7533f

#### Acknowledgement
* @id: urn:uuid:4fb3af44-d4f8-4226-9475-2d09c2d8d9e0

#### Announcement
* @id: urn:uuid:94ecae35-dcfd-4182-8550-22c7164fe23f

