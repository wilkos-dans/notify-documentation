# notify_diagramming

some vocabs we might want to use:

https://id.loc.gov/vocabulary/relators.html



Use the following example identifiers consistently, so that notification patterns can be combined into logically consistent scenarios:

### Repository
* @id: https://repository.org
* ldp:inbox: https://repository.org/inbox/

### Repository Object
* @id: https://repository.org/preprint/201203/421/
* ietf:cite-as: https://doi.org/10.5555/12345680
* url @id: https://repository.org/preprint/201203/421/content.pdf

### Review Service
* @id: https://reviewservice.org
* ldp:inbox: https://reviewservice.org/requests/inbox/

### Review Object
* @id: https://reviewservice.org/review/geo/202103/0021
* ietf:cite-as: https://doi.org/10.3214/987654

### Overlay Journal
* @id: https://journal.org
* ldp:inbox: https://journal.org/requests/inbox/

### Journal Object - review
* @id: https://journal.org/reviews/000001/00001
* ietf:cite-as: https://doi.org/10.3214/12345

### Journal Object - article
* @id: https://journal.org/articles/00001/
* ietf:cite-as: https://journal.org/articles/00001/

### Offer
* @id: urn:uuid:0370c0fb-bb78-4a9b-87f5-bed307a509dd

###Undo
* @id: urn:uuid:6eafed1d-935c-41b1-a5bb-645be4b7533f

### Acknowledgement
* @id: urn:uuid:4fb3af44-d4f8-4226-9475-2d09c2d8d9e0

### Announcement
* @id: urn:uuid:94ecae35-dcfd-4182-8550-22c7164fe23f

