#!/bin/bash

wget http://mtgjson.com/json/AllCards.json
node updateCards.js
