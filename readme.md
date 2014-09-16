# Dietary

## Server

* Clone this repo into your Go workspace (~/go/src/github.com/jesseobrien/dietary -- you can go get it)
* `cd server && go get`
* Run `go get github.com/pilu/fresh` on your command line
* Run `fresh -c fresh.conf`
* Run `sass --watch sass:public/css`
* Have a drink

### Premise

Dietary recommendations as a service. Sign up, punch in your targets for nutrient values, any allergies or intolerances, pick a way to receive the recommendation (sms, tweet, email, post) and off you go!

### Ideas

* Interface with workout and fitness tracking apps.
   * Option to adjust calories/macronutrients based on workouts logged to fitness apps
* Offer discounts with food ordering services like naturebox or something.
* Include appliance-based weighting for picking out recipes/meals. Favor the items the user uses most.
* Kitchen inventory/grocery list and recipe bias (mobile device barcode scanner integration).
* Dietary recommendations based on medical conditions (with option to follow or ignore and put in custom). ie - low-carb/low-GI for Diabetes, Gluten/Casein free for Celiac, etc (putting in this data is optional, but sets up presets for groups of options)
* Automatic categorizing of recipes based on ingredients used (ie - "dairy free" if no dairy is used, "paleo" if no dairy, refined sugar, or grain is used, etc)
* Fill/augment data from recipe APIs, such as http://api.bigoven.com/
* Allow for community-submitted recipes, voting systems allows for the most popular recipes to be added to the main database
* Rating main recipe store (we need to delve into this further to figure out how we want to do ratings in a way that relevant)
* "Cheat" option -- recommends "cheat" food/dessert based on community rating (still prioritizes set limitations?)
* Interval (weekly/daily/bi-monthly) grocery list generation
* 

### Tech/Libs

* Validation
  * https://github.com/mholt/binding

### Alpha Features

* User profiles
* Basic recipes/meals and associated framework/structure
* Basic recipe/meal/food tagging

### Beta Features

// Expand base with solid workable features

### Release Features

// Release functional site
