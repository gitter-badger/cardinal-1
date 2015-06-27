#!/bin/bash

wget http://hearthstonejson.com/json/AllCards.json
node updateCards.js
