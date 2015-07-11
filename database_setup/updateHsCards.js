var mongodb = require('mongodb');
var monk = require('monk');
var allCards = require("./AllCards.json");
var dbconfig = require("../dbconfig.json");
var db = monk(dbconfig.user + ':' + dbconfig.password + '@' + dbconfig.url + ':' + dbconfig.port + '/' + dbconfig.dbname);
var cardCollection = db.get('hearthstone');
var totalCards = 0;
var cardsAdded = 0;

function updateDb(err, doc){
    if(err) throw err;
    console.log("Successfully added " + doc.name);
    cardsAdded++;
    if(cardsAdded == totalCards){
        var timeElapsed = new Date() - timer;
        console.log(timeElapsed * 0.001);
        console.log(totalCards.toString() + " Added ending execution");
        process.exit(0);
    }
}

console.log(cardCollection);

for(var card in allCards) {
    totalCards++;
    console.log(allCards[card].name);
    cardCollection.insert(allCards[card], updateDb);
}
