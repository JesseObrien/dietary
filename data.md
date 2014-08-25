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

### Scrape Data

* Nutrition facts from somewhere
  * http://nutritiondata.self.com/ has some of the most thorough data, particularly for base foods
* Recipes from places
  * There appears to be some kind of "recipe display" standard...thing...that we ought to be able to read from many/most sites (example - http://wellnessmama.com/8959/homemade-fruit-snacks/)
  * Combine recipe data with base ingredient nutrition information to obtain numbers for a given recipe/meal

### Generated Data

* Recommended nutrition intake based on age, weight, sex, etc
* Include body fat % estimates/calculations based on information input (select formula based on what data is input)

### Thoughts On Nutritional Data

Nutritional data can get crazy, especially when we start talking about recipes/compound foods and whatnot. We'll want to think over how we go about storing this data to keep thrings from getting too chaotic (see also: the MyFitnessPal database).

Will we want to store only one set of measurements (tablespoons/cups vs grams/mL) and perhaps find a way to convert to the other, or should we store both in two separate fields? We'll probably want to look at what the scraped data gives us for the most part (with enough data, we can probably fill in the blanks in scrapes that only get some of the data).
