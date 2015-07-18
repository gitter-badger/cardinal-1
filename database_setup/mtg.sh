#!/bin/bash

if [ hash node 2>/dev/null ] 
    then
        wget http://mtgjson.com/json/AllSets.json
        node updateMagicCards.js
        rm AllSets.json
    else
        echo "Node is required for this operation please install node.js from http://nodejs.org/"
        exit 1
fi
