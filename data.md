## Data

Answer some questions about what kinds of data we'll be collecting from users, generating server-side, etc. The reason for answering this is to try to nail down how/why the data is related, and which data stores make the most sense for the types of data. For instance, redis is great for tracking metrics and statistical type relationships and things that need quick retreival or possibly have intersection operations done in-memory. Mongo does a decent job of storing embeddable data in nice, neat collections for easy JSON retrieval later on. SQL does what it does best, relational data at the expense of speed on indexing and other such things.

### User Data

*Profile*
* Email:string
* Name:string
* Gender:string (optional)
* Weight:unsigned int (optional)
* weight:units (lbs/kg/stone)
* age:unsigned int
* allergies:set
* intolerances:set
* preferred contact method:string
* SMS:string
* twitter/facebook/google connect?
* personal rss feed potentially?
  * Idea here would be give them their own rss for their updates for people who still love RSS
  * dietary.io/{random_hash}/feed or something
* 

### Scrape Data

* Nutrition facts from somewhere

### Generated Data

* Recommended nutrition intake based on age, weight, sex, etc

