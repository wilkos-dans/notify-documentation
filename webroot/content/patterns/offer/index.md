---
title: Offer
date: "2021-03-08"
weight: 1
description: "This pattern describes an *offer* activity, where one system offers one of its resources for some activity to be conducted by a second system."
---

## Generic Offer

<div class="row">
    <div class="col">
        <h5>Properties</h5>
        <ul>
            <li>The default <code>context</code> is Activity Streams 2.0</li>
            <li>Other contexts (namespaces) used:</li>
            <ul>
                <li><a href="https://schema.org">schema.org</a></li>
                <li><a href="https://www.w3.org/ns/ldp#">W3C Linked Data Platform (LDP) Vocabulary</a></li>
                <li><a href="http://www.iana.org/assignments/relation/">IETF Link Relations</a></li>
                <li><a href="http://purl.org/coar/notify_activity_type/">Notify Activity Types</a></li>
                <li><a href="http://purl.org/coar/resource_type">COAR Resource Types</a></li>
            </ul>
            <li>Notify notifications always describe an Activity Streams 2.0 <code>activity</code> as the base entity</li>
            <li>The <code>activity</code> has an <code>@id</code> which is an URN:UUID</li>
            <li>The <code>activity</code> has a <code>@type</code> array with values from Activity Streams 2.0 and Notify Activity Types</li>
            <li>The <code>activity</code> has an <code>origin</code> representing its 'originator'</li>
            <li>The <code>activity</code> has a <code>target</code> representing its 'destination'</li>
            <li>the resource being offered is represented as an <code>object</code></li>
            <ul>
                <li>the <code>@id</code> property of the <code>object</code> contains the HTTP URI to the "landing page"
                    for
                    the
                    resource
                </li>
                <li>the <code>@type</code> property of the <code>object</code> contains the value <i>AboutPage</i> from
                    the
                    <a href="https://schema.org/AboutPage">schema.org</a> vocabulary
                </li>
                <li>the <code>ietf:cite-as</code> property of the <code>object</code> contains the persistent HTTP URI
                    (sometimes called
                    the "PID") which is to be used to cite or link to the resource.
                </li>
                <li>the <code>url</code> property contains the details of the actual content resource</li>
                <ul>
                    <li>the <code>@id</code> property of the <code>url</code> contains the HTTP URI to the content file
                        for
                        the resource
                    </li>
                    <li>the <code>@type</code> property of the <code>url</code> describes the content file, including
                        values
                        from three
                        vocabularies
                    </li>
                    <li>the <code>mediaType</code> property of the <code>url</code> contains the MIME Type of the
                        content
                        file
                    </li>
                </ul>
            </ul>
        </ul>
    </div>
    <div class="col">
        <h5>Example</h5>
        {{< load_json "/patterns/offer/1.json" >}}
    </div>
</div>