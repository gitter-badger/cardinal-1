#Cardinal [![Build Status](https://travis-ci.org/ChasingLogic/cardinal.svg)](https://travis-ci.org/ChasingLogic/cardinal)

[![Join the chat at https://gitter.im/ChasingLogic/cardinal](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/ChasingLogic/cardinal?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
This app has been a dream of mine for a long time. I'm a huge fan of card games of all types and always wanted an app that I could throw my card collections into and run reports on.

###How To Help
We're still in the early days of development so grab an issue, make a branch and start hacking!

Make sure that in your pull requests and commit messages you list any issue #'s that you are fixing/resolving.

###What you'll need

You'll need [node](http://nodejs.org/) to update your card database and install [bower](http://bower.io/), and [gulp](http://gulpjs.com/) which you need to work with the front-end. For front-end dependencies there is the bower.json and package.json make sure if you are adding a dependency with npm to use the npm install --save <your package here> option so the package.json will be updated and to use the equivalent bower install --save <your package here> option so the bower.json will be updated.

You'll need [Go](http://golang.org) installed to work on the back-end. You can install all Go dependencies using go get in the server folder.

If there are enough requests I will seperate the front-end into it's own repository
