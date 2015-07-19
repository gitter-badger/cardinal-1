var mongodb = require('mongodb');
var monk = require('monk');
var allSets = require("./AllSets.json");
var allCards = require("./AllCards-x.json")
var dbconfig = require("./dbconfig.json");
var db = monk(dbconfig.user + ':' + dbconfig.password + '@' + dbconfig.url + ':' + dbconfig.port + '/' + dbconfig.dbname);
var cardCollection = db.get('magic');
var totalCards = 0;
var cardsAdded = 0;
var timer = new Date();

function updateDb(err, doc){
    if(err) throw err;
    if(doc != null){
        console.log("Successfully added " + doc.name);
    } else {
        console.log(doc);
    }
    cardsAdded++;
    if(cardsAdded == totalCards){
        var timeElapsed = new Date() - timer;
        console.log(timeElapsed * 0.001);
        console.log(totalCards.toString() + " Added ending execution");
        process.exit(0);
    }
}

for(var card in allCards) {
    totalCards++;
    console.log(allCards[card].name);
    cardCollection.insert(allCards[card], function(err, doc){
        if(err) throw err;
        console.log("Successfully added " + doc.name);
        cardsAdded++;
        if(cardsAdded == totalCards){
            var timeElapsed = new Date() - timer;
            console.log(timeElapsed * 0.001);
            console.log(totalCards.toString() + " Added ending execution");
            process.exit(0);
        }
    });
}

for(var set in allSets) {
    for(var card in allSets[set].cards){
        totalCards++;
        cardCollection.findAndModify({
            query: {name: allSets[set].cards[card].name},
            update: { $push: { multiverseids: allSets[set].cards[card].multiverseid } }
        }, function(err, doc){
            updateDb(err, doc);
        });
    }
}
