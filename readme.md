# Dietary

## Server

* Clone this repo into your Go workspace (~/go/src/github.com/jesseobrien/dietary -- you can go get it)
* `cd server && go get`
* Run `go get github.com/codegangsta/gin` on your command line
* Run `gin -a=8080`
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

### Alpha Features

// Add as a cohesive base

### Beta Features

// Expand base with solid workable features

### Release Features

// Release functional site
